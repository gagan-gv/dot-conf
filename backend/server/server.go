package server

import (
	"dot_conf/constants"
	"dot_conf/database"
	"dot_conf/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Initialize() {
	// Database
	err := database.Initialize()
	if err != nil {
		log.Error("Error initializing database, ", err)
		return
	}
	log.Info("Database setup successful")

	// Handlers
	companyHandler := handlers.NewCompanyHandler()

	// Endpoint Setup
	server := mux.NewRouter()

	// Company Routes
	server.HandleFunc(constants.ApiV1+constants.CompanyPath, companyHandler.Register).Methods(http.MethodPost)
	server.HandleFunc(constants.ApiV1+constants.CompanyPath+constants.CompanyId, companyHandler.Update).Methods(http.MethodPatch)
	server.HandleFunc(constants.ApiV1+constants.CompanyPath+constants.CompanyId, companyHandler.Fetch).Methods(http.MethodGet)
	server.HandleFunc(constants.ApiV1+constants.CompanyPath, companyHandler.FetchAll).Methods(http.MethodGet)

	// User Routes

	// Init Listen
	err = http.ListenAndServe(":9898", server)
	if err != nil {
		log.Error("Failed to listen at port 9898: ", err)
		return
	}
}
