package controllers

import (
	"net/http"

	"github.com/fear-the-dice/api/models"
)

// ControllerConfig allows for storage of the host, port, and template location
type ControllerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Controller struct {
	config  *ControllerConfig
	handler *http.Handler
}

// Attach accepts a http.ServeMux object and then creates all routes
func (this *Controller) Attach(mux *http.ServeMux) {
}

// Return a new Controller object with its own http.Handler
func NewController(config *ControllerConfig) *Controller {
	return &Controller{
		config:  config,
		handler: new(http.Handler),
	}
}
