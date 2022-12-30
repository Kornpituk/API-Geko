package config

import (
	"myapp-project-gecko/internal/config/gecko"
	"os"
)

type Config struct {
	Env         string
	Port        string
	ContextPath string
	Api gecko.Gecko
}

func New() (*Config, error) {

	var config Config

	config.Env = os.Getenv("GO_ENV")
	config.Port = os.Getenv("PORT")
	config.ContextPath = os.Getenv("CONTEXT_PATH")
	config.Api.Coin = os.Getenv("API_GECKO_COIN")

	return &config, nil

}
