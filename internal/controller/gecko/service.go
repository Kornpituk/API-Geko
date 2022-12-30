package gecko

import "myapp-project-gecko/internal/usecase/gecko"

type Service struct {
	gecko *gecko.Service
}

func NewService(gecko *gecko.Service) *Service {
	return &Service{
		gecko: gecko,
	}
}