package main

import (
	"github.com/gorilla/mux"
	"github.com/kubefirst/metaphor-go/configs"
	"github.com/kubefirst/metaphor-go/internal/handlers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func main() {
	// setup logging with color and code line on logs
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Caller().Logger()

	// mux router
	r := mux.NewRouter().StrictSlash(true)

	// server swagger-ui
	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("../../third_party/swagger-ui/")))
	r.PathPrefix("/swagger-ui/").Handler(sh)

	config := configs.ReadConfig()
	httpClient := http.DefaultClient

	appHandler := handlers.NewApp(config, httpClient)
	vaultHandler := handlers.NewVault(config, httpClient)
	kubernetesHandler := handlers.NewKubernetesHandler(config, httpClient)

	appRouter := r.PathPrefix("/app").Subrouter()
	appRouter.HandleFunc("/", appHandler.Info).Methods(http.MethodGet)
	appRouter.HandleFunc("/performance", appHandler.Performance).Methods(http.MethodGet)
	appRouter.HandleFunc("/healthz", appHandler.Healthz).Methods(http.MethodGet)
	appRouter.HandleFunc("/kill", appHandler.Kill).Methods(http.MethodGet)

	r.HandleFunc("/vault", vaultHandler.Vault).Methods(http.MethodGet)
	r.HandleFunc("/kubernetes", kubernetesHandler.KubernetesConfigMapData).Methods(http.MethodGet)

	port := ":3000"
	log.Info().Msgf("API listening at %q port", port[1:])
	if err := http.ListenAndServe(port, r); err != nil {
		log.Panic().Err(err).Msg("API is down")
	}
}
