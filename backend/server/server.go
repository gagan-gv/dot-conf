package server

import (
	"dot_conf/constants"
	"dot_conf/database"
	"dot_conf/handlers"
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

	// Handlers
	companyHandler := handlers.NewCompanyHandler()

	// Endpoint Setup
	server := http.NewServeMux()
	server.HandleFunc(constants.ApiV1+constants.CompanyPath+constants.Register, companyHandler.Register)
	server.HandleFunc(constants.ApiV1+constants.CompanyPath+constants.Update+constants.CompanyId, companyHandler.Update)
	server.HandleFunc(constants.ApiV1+constants.CompanyPath+constants.Fetch+constants.CompanyId, companyHandler.Fetch)
	server.HandleFunc(constants.ApiV1+constants.CompanyPath+constants.Fetch, companyHandler.FetchAll)
	err = http.ListenAndServe(":9898", server)
	if err != nil {
		log.Error("Failed to listen at port 9898: ", err)
		return
	}
}
