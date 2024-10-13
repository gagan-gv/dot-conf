package services

import (
	"context"
	"dot_conf/constants"
	"dot_conf/database"
	"dot_conf/dto"
	"dot_conf/models"
	"dot_conf/proto"
	"dot_conf/utils"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type IConfigService interface {
	Add(details dto.ConfigDetails, appId, createdBy string) dto.Response
	Delete(configId string) dto.Response
	DeleteAll(appId string) error
	Update(details dto.ConfigDetails, configId, updatedBy string) dto.Response
	Get(configId string) dto.Response
	GetAll(appId string) dto.Response
	Fetch(context.Context, *proto.ConfigRequest) (*proto.ConfigResponse, error)
}

type ConfigService struct {
	db   *gorm.DB
	mail IMailingService
	proto.ConfigServiceServer
}

var (
	configServiceInstance IConfigService
	configServiceOnce     sync.Once
)

func NewConfigService() IConfigService {
	configServiceOnce.Do(func() {
		configServiceInstance = &ConfigService{
			db:   database.GetDB(),
			mail: GetMailingService(),
		}
	})
	return configServiceInstance
}

func (c ConfigService) Add(details dto.ConfigDetails, appId, createdBy string) dto.Response {
	if database.ConfigAlreadyExists(details.Name, appId, &models.Config{}) {
		log.Info("Config already exists with app name ", details.Name, " for appId ", appId)
		return utils.NewErrorResponse(http.StatusBadRequest, constants.ConfigAlreadyExists, constants.AlreadyExists)
	}

	config := models.NewConfigBuilder().
		SetName(details.Name).
		SetDescription(details.Description).
		SetType(details.Type).
		SetValue(details.Value).
		SetAppID(appId).
		SetCreatedBy(createdBy).
		Build()

	err := c.db.Create(config).Error
	if err != nil {
		log.Error("Failed to save config: ", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, "config", config)
	return utils.NewSuccessResponse(http.StatusCreated, constants.Created, constants.Created, config)
}

func (c ConfigService) Delete(configId string) dto.Response {
	var config *models.Config
	err := database.FindConfigById(configId, &config).Error

	if err != nil {
		log.Error("Failed to fetch config: ", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	err = c.db.Delete(&config).Error
	if err != nil {
		log.Error("Failed to delete config: ", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	return utils.NewSuccessResponse(http.StatusNoContent, constants.DeletedSuccessfully, constants.Deleted, nil)
}

func (c ConfigService) DeleteAll(appId string) error {
	return c.db.Unscoped().Where("app_id = ?", appId).Delete(&models.Config{}).Error
}

func (c ConfigService) Update(details dto.ConfigDetails, configId, updatedBy string) dto.Response {
	var config *models.Config
	err := database.FindConfigById(configId, &config).Error

	if err != nil {
		log.Error("Failed to fetch config: ", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	if details.Value == nil || details.Value.(string) == "" {
		log.Error("Cannot update empty value")
		return utils.NewErrorResponse(http.StatusBadRequest, constants.InvalidUpdateRequest, constants.InvalidRequest)
	}

	config.Value = details.Value
	config.ModifiedBy = updatedBy
	year, month, date := time.Now().Date()
	config.ModifiedOn = fmt.Sprintf("%02d-%02d-%d", date, month, year)

	err = database.Update(&config)
	if err != nil {
		log.Error("Failed to update config: ", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, "config", config)

	return utils.NewSuccessResponse(http.StatusCreated, constants.Updated, constants.Updated, data)
}

func (c ConfigService) Get(configId string) dto.Response {
	var config *models.Config
	err := database.FindConfigById(configId, &config).Error

	if err != nil {
		log.Error("Failed to fetch config: ", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, "config", config)

	return utils.NewSuccessResponse(http.StatusOK, constants.Fetched, constants.Fetched, data)
}

func (c ConfigService) GetAll(appId string) dto.Response {
	var configs []models.Config
	err := c.db.Model(&models.Config{}).
		Select("name", "description", "type").
		Where("app_id = ?", appId).Find(&configs).
		Error

	if err != nil {
		log.Error("Failed to fetch configs for the app: ", appId, " ", err)
		return utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
	}

	data := utils.BuildData()
	utils.AddToData(data, "configs", configs)
	return utils.NewSuccessResponse(http.StatusOK, constants.Fetched, constants.Fetched, data)
}

func (c ConfigService) Fetch(ctx context.Context, request *proto.ConfigRequest) (*proto.ConfigResponse, error) {
	var app *models.App
	err := database.FindAppByKey(request.GetAppKey(), &app).Error

	if err != nil {
		log.Error("Error fetching app: ", err)
		return nil, errors.New("error fetching the config")
	}

	var config *models.Config
	err = c.db.Model(&models.Config{}).
		Where("app_id = ? AND name = ?", app.ID, request.GetConfigName()).
		Find(&config).
		Error

	if err != nil {
		log.Error(ctx, "Error fetching config: ", err.Error())
		return nil, errors.New("error fetching the config")
	}

	value, err := json.Marshal(config.Value)
	unquotedValue, err := strconv.Unquote(string(value))

	if err != nil {
		log.Error(ctx, "Error serializing config: ", err.Error())
		return nil, errors.New("error serializing config")
	}

	response := &proto.ConfigResponse{
		Value: unquotedValue,
	}

	return response, nil
}
