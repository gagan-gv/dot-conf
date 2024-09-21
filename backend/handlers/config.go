package handlers

import (
	"dot_conf/services"
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

}

func (c ConfigHandler) Delete(w http.ResponseWriter, r *http.Request) {

}

func (c ConfigHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (c ConfigHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (c ConfigHandler) GetAll(w http.ResponseWriter, r *http.Request) {

}
