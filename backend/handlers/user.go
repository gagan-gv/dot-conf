package handlers

import (
	"dot_conf/constants"
	"dot_conf/dto"
	"dot_conf/services"
	"dot_conf/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
)

type IUserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Deactivate(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	UserService services.IUserService
}

var (
	userHandlerInstance IUserHandler
	userHandlerOnce     sync.Once
)

func NewUserHandler() IUserHandler {
	userHandlerOnce.Do(func() {
		userHandlerInstance = &UserHandler{
			UserService: services.NewUserService(),
		}
	})

	return userHandlerInstance
}

func (u UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	log.Info("Initialized user registration")
	var ud dto.UserDetails

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&ud)

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

	response := u.UserService.Register(ud, int64(companyIdInt))
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (u UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	log.Info("Initialized user login")
	var ud dto.UserDetails

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&ud)

	if err != nil {
		log.Error("Unable to decode the request body due to: ", err.Error())
		response := utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.DecodingErrorResponse)
		w.Write([]byte(responseJson))
		return
	}

	response := u.UserService.Login(ud)
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (u UserHandler) Deactivate(w http.ResponseWriter, r *http.Request) {
	log.Info("Initialized user deactivation")
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

	vars := mux.Vars(r)
	userEmail, ok := vars["email"]

	if !ok {
		log.Error("Unable to find the user email for deactivation")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.NoPathVariableFound, constants.NoPathVariableFound)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.NoCompanyIdFoundResponse)
		w.Write([]byte(responseJson))
		return
	}

	response := u.UserService.Deactivate(userEmail, int64(companyIdInt))
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}
