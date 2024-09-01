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

type ICompanyHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Fetch(w http.ResponseWriter, r *http.Request)
	FetchAll(w http.ResponseWriter, r *http.Request)
}

type CompanyHandler struct {
	CompanyService services.ICompanyService
}

var (
	companyHandlerInstance ICompanyHandler
	companyHandlerOnce     sync.Once
)

func NewCompanyHandler() ICompanyHandler {
	companyHandlerOnce.Do(func() {
		companyHandlerInstance = &CompanyHandler{
			CompanyService: services.NewCompanyService(),
		}
	})

	return companyHandlerInstance
}

func (h CompanyHandler) Register(w http.ResponseWriter, r *http.Request) {
	log.Info("Initialized company registration")
	w.Header().Set("Content-Type", "application/json")
	response := h.CompanyService.Register(r)
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (h CompanyHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Info("Initialized company update")
	var rc *dto.RegisterCompany

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&rc)
	if err != nil {
		log.Error("Unable to decode the request body due to: ", err.Error())
		response := utils.NewErrorResponse(http.StatusInternalServerError, constants.GeneralError, err.Error())
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.DecodingErrorResponse)
		w.Write([]byte(responseJson))
		return
	}

	vars := mux.Vars(r)
	companyId, ok := vars["companyId"]

	if !ok {
		log.Error("Unable to find the company id")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.NoPathVariableFound, constants.NoPathVariableFound)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.NoCompanyIdFoundResponse)
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

	response := h.CompanyService.Update(rc, int64(companyIdInt))
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (h CompanyHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	log.Info("Fetching company with specific id")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	companyId, ok := vars["companyId"]

	if !ok {
		log.Error("Unable to find the company id")
		response := utils.NewErrorResponse(http.StatusBadRequest, constants.NoPathVariableFound, constants.NoPathVariableFound)
		w.WriteHeader(response.StatusCode)
		responseJson := utils.ConvertJsonToString(response, constants.NoCompanyIdFoundResponse)
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

	response := h.CompanyService.Fetch(int64(companyIdInt))
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}

func (h CompanyHandler) FetchAll(w http.ResponseWriter, r *http.Request) {
	log.Info("Fetching all companies")
	w.Header().Set("Content-Type", "application/json")

	response := h.CompanyService.FetchAll()
	responseJson := utils.ConvertJsonToString(response, constants.Empty)
	w.Write([]byte(responseJson))
}
