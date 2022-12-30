package gecko

import "myapp-project-gecko/internal/config"

type Service struct {
	config *config.Config
}

func NewService(config *config.Config) *Service {
	return &Service{
		config: config,
	}
}
