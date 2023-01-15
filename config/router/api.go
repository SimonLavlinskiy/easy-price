package router

import (
	"easyPrice/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func apiRouter(f fiber.Router) {
	//appController := initializationController()

}

func initializationController() controllers.Controller {
	return controllers.Controller{}
}
