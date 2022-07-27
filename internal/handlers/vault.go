package handlers

import (
	"github.com/kubefirst/metaphor-go/configs"
	"github.com/kubefirst/metaphor-go/pkg"
	"net/http"
)

type VaultHandler struct {
	config     *configs.Config
	httpClient pkg.HTTPDoer
}

func NewVault(config *configs.Config, httpClient pkg.HTTPDoer) VaultHandler {
	return VaultHandler{
		config:     config,
		httpClient: httpClient,
	}
}

func (handler VaultHandler) Vault(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
}
