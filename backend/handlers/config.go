package handlers

import (
	"dot_conf/constants"
	"dot_conf/dto"
	"dot_conf/jwt"
	"dot_conf/services"
	"dot_conf/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

type IConfigHandler interface {
	Add(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

type ConfigHandler struct {
	ConfigService services.IConfigService
}

var (
	configHandlerInstance IConfigHandler
	configHandlerOnce     sync.Once
)

func NewConfigHandler() IConfigHandler {
	configHandlerOnce.Do(func() {
		configHandlerInstance = &ConfigHandler{
			ConfigService: services.NewConfigService(),
		}
	})

	return configHandlerInstance
}

func (c ConfigHandler) Add(w http.ResponseWriter, r *http.Request) {
	log.Info("Initializing config addition")
	var cd dto.ConfigDetails

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&cd)

	if err != nil {
		log.Error("Unable to decode the request body due to: ", err.Error())
		response := utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.DecodingErrorResponse)
		w.Write([]byte(responseJson))
		return
	}

	headers := r.Header
	appId := headers.Get("app-id")

	if appId == "" {
		log.Error("app-id header is missing")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.HeaderMissing, constants.HeaderMissing)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.HeaderIsMissing)
		w.Write([]byte(responseJson))
		return
	}

	createdBy, err := jwt.GetUsername(r)
	if err != nil {
		log.Error("Error fetching username due to: ", err)
		response := utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.Empty)
		w.Write([]byte(responseJson))
		return
	}

	response := c.ConfigService.Add(cd, appId, createdBy)
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (c ConfigHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("Initializing config deletion")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	configId, ok := vars["configId"]
	if !ok {
		log.Error("Unable to find the configId in the path")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.NoPathVariableFound, constants.NoPathVariableFound)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.NoConfigIdFoundResponse)
		w.Write([]byte(responseJson))
		return
	}

	response := c.ConfigService.Delete(configId)
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (c ConfigHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("Initializing config update")
	var cd dto.ConfigDetails
	vars := mux.Vars(r)

	configId, ok := vars["configId"]
	if !ok {
		log.Error("Unable to find the configId in the path")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.NoPathVariableFound, constants.NoPathVariableFound)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.NoConfigIdFoundResponse)
		w.Write([]byte(responseJson))
		return
	}

	updatedBy, err := jwt.GetUsername(r)
	if err != nil {
		log.Error("Error fetching username due to: ", err)
		response := utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.Empty)
		w.Write([]byte(responseJson))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewDecoder(r.Body).Decode(&cd)

	if err != nil {
		log.Error("Unable to decode the request body due to: ", err.Error())
		response := utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.DecodingErrorResponse)
		w.Write([]byte(responseJson))
		return
	}

	response := c.ConfigService.Update(cd, configId, updatedBy)
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (c ConfigHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("Initializing config get")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	configId, ok := vars["configId"]
	if !ok {
		log.Error("Unable to find the configId in the path")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.NoPathVariableFound, constants.NoPathVariableFound)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.NoConfigIdFoundResponse)
		w.Write([]byte(responseJson))
		return
	}

	response := c.ConfigService.Get(configId)
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (c ConfigHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("Initializing config getAll")
	w.Header().Set("Content-Type", "application/json")
	headers := r.Header
	appId := headers.Get("app-id")

	if appId == "" {
		log.Error("app-id header is missing")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.HeaderMissing, constants.HeaderMissing)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.HeaderIsMissing)
		w.Write([]byte(responseJson))
		return
	}

	response := c.ConfigService.GetAll(appId)
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}
