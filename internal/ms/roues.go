package ms

import (
	"myapp-project-gecko/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func (s *Service) InitRoutes() {

	e := s.Echo
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	contextPath := s.Config.ContextPath

	e.GET(contextPath+"/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, domain.Health{
			Status: "OK",
		})
	})

	e.GET("/gecko", s.HTTP.Gecko.Gecko)
}
