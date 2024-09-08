package handlers

import (
	"dot_conf/constants"
	"dot_conf/dto"
	"dot_conf/jwt"
	"dot_conf/services"
	"dot_conf/utils"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
)

type IAppHandler interface {
	Add(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	FetchAll(w http.ResponseWriter, r *http.Request)
}

type AppHandler struct {
	AppService services.IAppService
}

var (
	appHandlerInstance IAppHandler
	appHandlerOnce     sync.Once
)

func NewAppHandler() IAppHandler {
	appHandlerOnce.Do(func() {
		appHandlerInstance = &AppHandler{
			AppService: services.NewAppService(),
		}
	})

	return appHandlerInstance
}

func (h *AppHandler) Add(w http.ResponseWriter, r *http.Request) {
	log.Info("Initialized new app addition")
	var ad dto.AppRegistrationDetails

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&ad)

	if err != nil {
		log.Error("Unable to decode the request body due to: ", err.Error())
		response := utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.DecodingErrorResponse)
		w.Write([]byte(responseJson))
		return
	}

	headers := r.Header
	companyId := headers.Get("company-id")

	if companyId == "" {
		log.Error("company-id header is missing")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.HeaderMissing, constants.HeaderMissing)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.HeaderIsMissing)
		w.Write([]byte(responseJson))
		return
	}

	companyIdInt, err := strconv.Atoi(companyId)
	if err != nil {
		log.Error("Unable to convert the company id to int due to: ", err.Error())
		response := utils.NewErrorResponse(http.StatusInternalServerError, constants.TextToIntConversionError, constants.TextToIntConversionError)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.TextToIntErrorResponse)
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

	response := h.AppService.Add(ad, int64(companyIdInt), createdBy)
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (h *AppHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Info("Initialized new app deletion")

	w.Header().Set("Content-Type", "application/json")
	headers := r.Header
	appKey := headers.Get("app-key")

	if appKey == "" {
		log.Error("app-key header is missing")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.HeaderMissing, constants.HeaderMissing)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.HeaderIsMissing)
		w.Write([]byte(responseJson))
		return
	}

	username, err := jwt.GetUsername(r)
	if err != nil {
		log.Error("Error fetching username due to: ", err)
		response := utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.Empty)
		w.Write([]byte(responseJson))
		return
	}

	response := h.AppService.Delete(appKey, username)
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (h *AppHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("Initialized new app update")
	var ad dto.AppRegistrationDetails

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&ad)

	if err != nil {
		log.Error("Unable to decode the request body due to: ", err.Error())
		response := utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.DecodingErrorResponse)
		w.Write([]byte(responseJson))
		return
	}

	headers := r.Header
	appKey := headers.Get("app-key")

	if appKey == "" {
		log.Error("app-key header is missing")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.HeaderMissing, constants.HeaderMissing)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.HeaderIsMissing)
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

	response := h.AppService.Update(appKey, updatedBy, ad)
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (h *AppHandler) FetchAll(w http.ResponseWriter, r *http.Request) {
	log.Info("Initialized new app fetchAll")
	w.Header().Set("Content-Type", "application/json")

	headers := r.Header
	companyId := headers.Get("company-id")

	if companyId == "" {
		log.Error("company-id header is missing")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.HeaderMissing, constants.HeaderMissing)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.HeaderIsMissing)
		w.Write([]byte(responseJson))
		return
	}

	companyIdInt, err := strconv.Atoi(companyId)
	if err != nil {
		log.Error("Unable to convert the company id to int due to: ", err.Error())
		response := utils.NewErrorResponse(http.StatusInternalServerError, constants.TextToIntConversionError, constants.TextToIntConversionError)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.TextToIntErrorResponse)
		w.Write([]byte(responseJson))
		return
	}

	response := h.AppService.FetchAll(int64(companyIdInt))
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}
