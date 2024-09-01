package server

import (
	"dot_conf/constants"
	"dot_conf/database"
	"dot_conf/handlers"
	"dot_conf/jwt"
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
	userHandler := handlers.NewUserHandler()

	// Endpoint Setup
	router := mux.NewRouter()
	adminRouter := router.PathPrefix(constants.ApiV1).Subrouter()
	adminRouter.Use(jwt.Verify("ADMIN"))
	superAdminRouter := router.PathPrefix(constants.ApiV1).Subrouter()
	superAdminRouter.Use(jwt.Verify("SUPER_ADMIN"))

	// Company Routes
	router.HandleFunc(constants.ApiV1+constants.CompanyPath+constants.Empty, companyHandler.Register).Methods(http.MethodPost)
	adminRouter.HandleFunc(constants.CompanyPath+constants.CompanyId, companyHandler.Update).Methods(http.MethodPatch)
	superAdminRouter.HandleFunc(constants.CompanyPath+constants.CompanyId, companyHandler.Fetch).Methods(http.MethodGet)
	superAdminRouter.HandleFunc(constants.CompanyPath+constants.Empty, companyHandler.FetchAll).Methods(http.MethodGet)

	// User Routes
	router.HandleFunc(constants.ApiV1+constants.UserPath+constants.Auth, userHandler.Login).Methods(http.MethodPost)
	adminRouter.HandleFunc(constants.UserPath, userHandler.Register).Methods(http.MethodPost)
	adminRouter.HandleFunc(constants.UserPath+constants.EmailId, userHandler.Deactivate).Methods(http.MethodPatch)

	// Init Listen
	err = http.ListenAndServe(":9898", router)
	if err != nil {
		log.Error("Failed to listen at port 9898: ", err)
		return
	}
}
