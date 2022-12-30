package ms

import (
	"myapp-project-gecko/internal/config"
	httptransport "myapp-project-gecko/internal/controller"

	"github.com/labstack/echo/v4"
)

type Service struct {
	Echo   *echo.Echo
	HTTP   *httptransport.Service
	Config *config.Config
}
