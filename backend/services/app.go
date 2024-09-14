package services

import (
	"dot_conf/constants"
	"dot_conf/database"
	"dot_conf/dto"
	"dot_conf/models"
	"dot_conf/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"slices"
	"time"
)

type IAppService interface {
	Add(details dto.AppRegistrationDetails, companyId int64, createdBy string) dto.Response
	Delete(appKey, email string) dto.Response
	Update(appKey, updatedBy string, details dto.AppRegistrationDetails) dto.Response
	FetchAll(companyId int64) dto.Response
}

type AppService struct {
	mail IMailingService
	db   *gorm.DB
}

func NewAppService() IAppService {
	return &AppService{
		db:   database.GetDB(),
		mail: GetMailingService(),
	}
}

func (a AppService) Add(details dto.AppRegistrationDetails, companyId int64, createdBy string) dto.Response {
	if database.AppAlreadyExists(companyId, details.Name, &models.App{}) {
		log.Info("App already exists with the given appName", details.Name, " for the company id", companyId)
		return utils.NewErrorResponse(http.StatusBadRequest, constants.AppAlreadyExists, constants.AlreadyExists)
	}

	app := models.NewAppBuilder().
		SetName(details.Name).
		SetCompanyId(companyId).
		SetCreatedBy(createdBy).
		Build()

	err := a.db.Create(app).Error
	if err != nil {
		log.Error("Failed to create app", details.Name, " due to:", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, constants.UnableToSaveData)
	}

	data := utils.BuildData()
	utils.AddToData(data, "app", app)
	return utils.NewSuccessResponse(http.StatusCreated, constants.Created, constants.Created, app)
}

func (a AppService) Delete(appKey, email string) dto.Response {
	var app *models.App
	err := database.FindAppByKey(appKey, &app).Error

	if err != nil {
		log.Error("App with the given key is not found", err)
		return utils.NewErrorResponse(http.StatusBadRequest, constants.AppNotFound, err.Error())
	}

	if !slices.Contains(app.Owners, email) {
		log.Info("Forbidden to delete the app")
		return utils.NewErrorResponse(http.StatusForbidden, constants.Forbidden, constants.Forbidden)
	}

	err = a.db.Delete(app).Error

	if err != nil {
		log.Error("Error while deleting app", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	return utils.NewSuccessResponse(http.StatusNoContent, constants.DeletedSuccessfully, constants.Deleted, nil)
}

func (a AppService) Update(appKey, updatedBy string, details dto.AppRegistrationDetails) dto.Response {
	var app *models.App
	err := database.FindAppByKey(appKey, &app).Error
	if err != nil {
		log.Error("App with the given key is not found", err)
		return utils.NewErrorResponse(http.StatusBadRequest, constants.AppNotFound, err.Error())
	}

	if !slices.Contains(app.Owners, updatedBy) {
		log.Info("Forbidden to update the app")
		return utils.NewErrorResponse(http.StatusForbidden, constants.Forbidden, constants.Forbidden)
	}

	if details.Name != "" {
		app.Name = details.Name
	}

	if details.OwnerEmails != nil {
		app.Owners = details.OwnerEmails
	}

	year, month, date := time.Now().Date()
	app.ModifiedBy = updatedBy
	app.ModifiedOn = fmt.Sprintf("%02d-%02d-%d", date, month, year)

	err = a.db.Save(app).Error
	if err != nil {
		log.Error("Unable to save the app", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, "app", app)
	return utils.NewSuccessResponse(http.StatusOK, constants.Updated, constants.Updated, app)
}

func (a AppService) FetchAll(companyId int64) dto.Response {
	var apps []models.App
	err := a.db.Where("company_id = ?", companyId).Find(&apps).Error
	if err != nil {
		log.Error("Could not fetch all apps for the company with ID: ", companyId, " with err: ", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.CouldNotFetchFromDatabase, err.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, "apps", apps)
	return utils.NewSuccessResponse(http.StatusOK, constants.ValueFetched, constants.Fetched, data)
}
