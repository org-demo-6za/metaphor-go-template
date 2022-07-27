package handlers

import (
	"github.com/kubefirst/metaphor-go/configs"
	"github.com/kubefirst/metaphor-go/pkg"
	"net/http"
)

type KubernetesHandler struct {
	config     *configs.Config
	httpClient pkg.HTTPDoer
}

func NewKubernetesHandler(config *configs.Config, httpClient pkg.HTTPDoer) KubernetesHandler {
	return KubernetesHandler{
		config:     config,
		httpClient: httpClient,
	}
}

func (handler KubernetesHandler) KubernetesConfigMapData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAlreadyReported)
}
