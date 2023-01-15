package main

import (
	"easyPrice/app/middleware"
	"easyPrice/config"
	"easyPrice/config/router"
	_ "easyPrice/docs"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/swaggo/http-swagger"
	"log"
	"os"
)

// @title                      Easy Price - get price auto is easy!
// @version                    1.0.0
// @in                         header
// @name                       Authorization
// @BasePath                   /
// @securityDefinitions.basic ApiKeyAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("Not found .env file, err: ", err)
	}

	app := fiber.New(config.FiberConfig())
	middleware.AppMiddleware(app)
	router.AppRouter(app)

	config.AppConfig()

	err := app.Listen(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))

	if err != nil {
		log.Panicf("Server is not running: %v", err.Error())
	}

}

// command for init
// $GOPATH/bin/swag init
