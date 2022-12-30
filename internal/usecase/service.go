package usecase

import (
	"myapp-project-gecko/internal/config"
	"myapp-project-gecko/internal/usecase/gecko"
)

type Service struct {
	Gecko *gecko.Service
}

func NewService(config *config.Config) *Service {
	return &Service{
		Gecko: gecko.NewService(config),
	}
}
