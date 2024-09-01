package services

import (
	"dot_conf/constants"
	"dot_conf/database"
	"dot_conf/dto"
	"dot_conf/jwt"
	"dot_conf/models"
	"dot_conf/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type IUserService interface {
	Register(details dto.UserDetails, companyId int64) dto.Response
	Login(details dto.UserDetails) dto.Response
	Deactivate(email string, companyId int64) dto.Response
}

type UserService struct {
	mail IMailingService
	db   *gorm.DB
}

func NewUserService() IUserService {
	return &UserService{
		db:   database.GetDB(),
		mail: GetMailingService(),
	}
}

func (u *UserService) Register(details dto.UserDetails, companyId int64) dto.Response {
	if (database.EmailAlreadyExists(details.Email, &models.User{})) {
		log.Info("Email already exists")
		return utils.NewErrorResponse(http.StatusBadRequest, constants.UserAlreadyExists, constants.UserAlreadyExists)
	}

	user := models.NewUserBuilder().
		SetName(details.Name).
		SetEmail(details.Email).
		SetCompanyId(companyId).
		SetPassword(details.Password).
		Build()

	err := u.db.Save(user).Error

	if err != nil {
		log.Error("Failed while saving the data", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.UnableToSaveData, err.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, constants.User, user)

	return utils.NewSuccessResponse(http.StatusCreated, constants.UserCreated, constants.Created, data)
}

func (u *UserService) Login(details dto.UserDetails) dto.Response {
	if details.Email == "" || details.Password == "" {
		log.Info("Email and/or Password is empty")
		return utils.NewErrorResponse(http.StatusUnprocessableEntity, constants.CredentialsMissing, constants.CredentialsMissing)
	}

	var user models.User
	err := database.FindByEmail(details.Email, &user).Error

	if err != nil {
		log.Error("User with this email doesn't exists", err)
		return utils.NewErrorResponse(http.StatusBadRequest, constants.UserNotFound, err.Error())
	}

	if details.Password != user.Password {
		log.Info("Invalid password", details.Email)
		return utils.NewErrorResponse(http.StatusBadRequest, constants.InvalidCredentials, constants.InvalidCredentials)
	}

	authToken, err := jwt.Generate(details.Email, user.Role.String())

	if err != nil {
		log.Error("Error generating the token", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, constants.User, user)
	utils.AddToData(data, constants.Token, authToken)

	return utils.NewSuccessResponse(http.StatusOK, constants.LoggedInSuccess, constants.LoggedIn, data)
}

func (u *UserService) Deactivate(email string, companyId int64) dto.Response {
	var user models.User
	err := database.FindByEmail(email, &user).Error

	if err != nil {
		log.Error("User with this email doesn't exists", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.UserNotFound, err.Error())
	}

	if user.CompanyID != companyId {
		log.Info("Tried to deactivate unknown user")
		return utils.NewErrorResponse(http.StatusBadRequest, constants.UserNotFound, constants.UserNotFound)
	}

	user.Status = models.INACTIVE

	err = u.db.Save(user).Error
	if err != nil {
		log.Error("Failed while saving the data", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.UnableToSaveData, err.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, constants.User, user)

	return utils.NewSuccessResponse(http.StatusOK, constants.DeactivatedSuccessfully, constants.Deactivated, data)
}
