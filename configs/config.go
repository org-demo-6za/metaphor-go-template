package configs

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type Config struct {
	ChartVersion string `env:"CHART_VERSION"`
	DockerTag    string `env:"DOCKER_TAG"`
	SecretOne    string `env:"SECRET_ONE"`
	SecretTwo    string `env:"SECRET_TWO"`
	ConfigOne    string `env:"CONFIG_ONE"`
	ConfigTwo    string `env:"CONFIG_TWO"`
}

func ReadConfig() *Config {
	config := Config{}

	if err := env.Parse(&config); err != nil {
		log.Err(err).Msg("something went wrong loading the environment variables")
		log.Panic().Send()
	}

	return &config
}
