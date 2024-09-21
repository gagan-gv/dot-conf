package server

import (
	"dot_conf/constants"
	"dot_conf/database"
	"dot_conf/handlers"
	"dot_conf/jwt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func Initialize() {
	// Database
	err := database.Initialize()
	if err != nil {
		log.Error("Error initializing database: ", err)
		return
	}
	log.Info("Database setup successful")

	httpServer()
}

func httpServer() {
	allowedHost := viper.GetString(constants.AllowedHost)

	// Handlers
	companyHandler := handlers.NewCompanyHandler()
	userHandler := handlers.NewUserHandler()
	appHandler := handlers.NewAppHandler()
	configHandler := handlers.NewConfigHandler()

	// Endpoint Setup
	router := mux.NewRouter()
	router.Use(hostRestrictionMiddleware(allowedHost))
	adminRouter := router.PathPrefix(constants.ApiV1).Subrouter()
	adminRouter.Use(jwt.Verify("ADMIN"))
	superAdminRouter := router.PathPrefix(constants.ApiV1).Subrouter()
	superAdminRouter.Use(jwt.Verify("SUPER_ADMIN"))
	userRouter := router.PathPrefix(constants.ApiV1).Subrouter()
	userRouter.Use(jwt.Verify("USER"))

	// Company Routes
	router.HandleFunc(constants.ApiV1+constants.CompanyPath, companyHandler.Register).Methods(http.MethodPost)
	adminRouter.HandleFunc(constants.CompanyPath+constants.CompanyId, companyHandler.Update).Methods(http.MethodPatch)
	superAdminRouter.HandleFunc(constants.CompanyPath+constants.CompanyId, companyHandler.Fetch).Methods(http.MethodGet)
	superAdminRouter.HandleFunc(constants.CompanyPath, companyHandler.FetchAll).Methods(http.MethodGet)

	// User Routes
	router.HandleFunc(constants.ApiV1+constants.UserPath+constants.Auth, userHandler.Login).Methods(http.MethodPost)
	adminRouter.HandleFunc(constants.UserPath, userHandler.Register).Methods(http.MethodPost)
	adminRouter.HandleFunc(constants.UserPath+constants.EmailId, userHandler.Deactivate).Methods(http.MethodPatch)

	// App Routes
	userRouter.HandleFunc(constants.AppPath, appHandler.Add).Methods(http.MethodPost)
	userRouter.HandleFunc(constants.AppPath, appHandler.Delete).Methods(http.MethodDelete)
	userRouter.HandleFunc(constants.AppPath, appHandler.Update).Methods(http.MethodPatch)
	userRouter.HandleFunc(constants.AppPath, appHandler.FetchAll).Methods(http.MethodGet)

	// Config Routes
	userRouter.HandleFunc(constants.ConfigPath, configHandler.Add).Methods(http.MethodPost)
	userRouter.HandleFunc(constants.ConfigPath, configHandler.Update).Methods(http.MethodPatch)
	userRouter.HandleFunc(constants.ConfigPath, configHandler.Delete).Methods(http.MethodDelete)
	userRouter.HandleFunc(constants.ConfigPath+constants.AppId, configHandler.GetAll).Methods(http.MethodGet)
	userRouter.HandleFunc(constants.ConfigPath+constants.ConfigId, configHandler.Get).Methods(http.MethodGet)

	// Init Listen
	err := http.ListenAndServe(":9898", router)
	if err != nil {
		log.Error("Failed to listen at port 9898: ", err)
		return
	}
}

func hostRestrictionMiddleware(allowedHost string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			host := r.Host

			if strings.EqualFold(host, allowedHost) {
				next.ServeHTTP(w, r)
				return
			}

			// If host is not allowed, return 403 Forbidden
			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}
