package main

import (

	"log"
	"myapp-project-gecko/internal/config"
	"myapp-project-gecko/internal/ms"
	"myapp-project-gecko/internal/controller"
	"myapp-project-gecko/internal/usecase"


	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load("./etc/local.env")
	if err != nil {
		log.Printf("please consider environment variables: %s", err)
	}


	msService := newMSService(echo.New())
	msService.InitRoutes()
	msService.Server()

}

func newMSService(e *echo.Echo) *ms.Service {

	config, err := config.New()
	if err != nil {
		panic(err)
	}

	usercases := usecase.NewService(config)

	return &ms.Service{
		Echo: e,
		HTTP: controller.NewService(usercases),
		Config: config,
	}
}
