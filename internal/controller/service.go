package controller

import (
	"myapp-project-gecko/internal/usecase"
	"myapp-project-gecko/internal/controller/gecko"

)

type Service struct {
	Gecko *gecko.Service
}

func NewService(usecases *usecase.Service) *Service{
	return &Service{
		Gecko: gecko.NewService(usecases.Gecko),
	}
}