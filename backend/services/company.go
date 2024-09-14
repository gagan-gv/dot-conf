package services

import (
	"dot_conf/constants"
	"dot_conf/database"
	"dot_conf/dto"
	"dot_conf/models"
	"dot_conf/utils"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type ICompanyService interface {
	Register(r *http.Request) dto.Response
	FetchAll() dto.Response
	Fetch(companyId int64) dto.Response
	Update(rc *dto.RegisterCompany, companyId int64) dto.Response
}

type CompanyService struct {
	mail IMailingService
	db   *gorm.DB
}

func NewCompanyService() ICompanyService {
	return &CompanyService{
		mail: GetMailingService(),
		db:   database.GetDB(),
	}
}

func (s *CompanyService) Register(r *http.Request) dto.Response {
	var rc *dto.RegisterCompany

	reqBody := r.FormValue("metadata")
	err := json.Unmarshal([]byte(reqBody), &rc)
	if err != nil {
		log.Error("Unable to decode the request body due to: ", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	if database.EmailAlreadyExists(rc.CompanyEmail, &models.Company{}) {
		log.Info("Company has already been registered with email id: ", rc.CompanyEmail)
		return utils.NewErrorResponse(http.StatusBadRequest, constants.CompanyAlreadyRegistered, constants.CompanyAlreadyRegistered)
	}

	if database.EmailAlreadyExists(rc.AdminEmail, &models.User{}) {
		log.Info("User has already been registered with email id: ", rc.AdminEmail)
		return utils.NewErrorResponse(http.StatusBadRequest, constants.UserAlreadyExists, constants.UserAlreadyExists)
	}

	filePath, err := utils.Upload(r)
	if err != nil {
		log.Error("Unable to upload the file due to: ", err.Error())
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	admin := models.NewUserBuilder().
		SetName(rc.AdminName).
		SetEmail(rc.AdminEmail).
		SetRole(models.ADMIN).
		SetPassword(rc.Password).
		SetStatus(models.ACTIVE).
		Build()
	company := models.NewCompanyBuilder().
		SetName(rc.CompanyName).
		SetAdminId(admin.ID).
		SetEmail(rc.CompanyEmail).
		SetDocumentPath(filePath).
		Build()

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if er := tx.Create(&company).Error; err != nil {
			log.Error("Error saving company: ", err.Error())
			return er
		}

		admin.CompanyID = company.ID
		if er := tx.Create(&admin).Error; err != nil {
			log.Error("Error saving admin: ", err.Error())
			return er
		}

		return nil
	})

	if err != nil {
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.UnableToSaveData, err.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, constants.Company, company)
	utils.AddToData(data, constants.Admin, admin)

	return utils.NewSuccessResponse(http.StatusCreated, constants.CompanyCreated, constants.Created, data)
}

func (s *CompanyService) FetchAll() dto.Response {
	var companies []models.Company
	result := s.db.Find(&companies)

	if result.Error != nil {
		log.Error("Could not fetch all companies due to: ", result.Error)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.CouldNotFetchFromDatabase, result.Error.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, constants.Companies, companies)
	return utils.NewSuccessResponse(http.StatusOK, constants.ValueFetched, constants.Fetched, data)
}

func (s *CompanyService) Fetch(companyId int64) dto.Response {
	var company models.Company
	result := s.db.First(&company, companyId)

	if result.Error != nil {
		log.Error("Could not company details for companyId: ", companyId, " due to: ", result.Error)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.CouldNotFetchFromDatabase, result.Error.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, constants.Company, company)
	return utils.NewSuccessResponse(http.StatusOK, constants.ValueFetched, constants.Fetched, data)
}

func (s *CompanyService) Update(rc *dto.RegisterCompany, companyId int64) dto.Response {
	var company models.Company
	var user models.User
	result := s.db.First(&company, companyId)

	if result.Error != nil {
		log.Error("Could not company details for companyId: ", companyId, " due to: ", result.Error)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.CouldNotFetchFromDatabase, result.Error.Error())
	}

	if rc.AdminEmail != "" {
		result = s.db.Model(&user).Where("email = ?", rc.AdminEmail).First(&user)
	}

	if result.Error != nil {
		log.Error("Could not new admin details for companyId: ", companyId, " due to: ", result.Error)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.CouldNotFetchFromDatabase, result.Error.Error())
	}

	if rc.CompanyName != "" && rc.CompanyEmail != "" {
		company.Name = rc.CompanyName
		company.Email = rc.CompanyEmail
	}

	if rc.AdminEmail != "" {
		company.AdminId = user.ID
	}

	err := database.Update(company)

	if err != nil {
		log.Error("Could not update company details for companyId: ", companyId, " due to: ", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.FailedUpdatingData, err.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, constants.Company, company)
	return utils.NewSuccessResponse(http.StatusOK, constants.Updated, constants.Updated, data)
}
