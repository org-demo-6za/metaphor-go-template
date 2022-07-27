package handlers

import (
	"github.com/kubefirst/metaphor-go/configs"
	"github.com/kubefirst/metaphor-go/pkg"
	"net/http"
)

type AppHandler struct {
	config     *configs.Config
	httpClient pkg.HTTPDoer
}

func NewApp(config *configs.Config, httpClient pkg.HTTPDoer) *AppHandler {
	return &AppHandler{
		config:     config,
		httpClient: httpClient,
	}
}

func (handler AppHandler) Info(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
}

func (handler AppHandler) Performance(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadGateway)
}

func (handler AppHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (handler AppHandler) Kill(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
