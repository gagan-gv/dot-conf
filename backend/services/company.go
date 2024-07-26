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
	Verify(companyId int64) dto.Response
	Update(rc *dto.RegisterCompany, companyId int64) dto.Response
}

type CompanyService struct {
	mail *MailingService
	db   *gorm.DB
}

func NewCompanyService() *CompanyService {
	return &CompanyService{
		mail: GetMailingService(),
		db:   database.GetDB(),
	}
}

func (s *CompanyService) Register(r *http.Request) dto.Response {
	var rc dto.RegisterCompany

	err := json.NewDecoder(r.Body).Decode(&rc)
	if err != nil {
		log.Error("Unable to decode the request body due to: ", err.Error())
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	filePath, err := utils.Upload(r)
	if err != nil {
		log.Error("Unable to upload the file due to: ", err.Error())
		return utils.NewErrorResponse(http.StatusNotImplemented, constants.GeneralError, err.Error())
	}

	admin := models.NewUserBuilder().
		SetName(rc.AdminName).
		SetEmail(rc.AdminEmail).
		SetRole(models.ADMIN).
		SetPassword(rc.Password).
		SetStatus(models.INACTIVE).
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

	data := make(map[string]interface{})
	data["company"] = company
	data["admin"] = admin

	return utils.NewSuccessResponse(http.StatusCreated, constants.CompanyCreated, constants.Created, data)
}
