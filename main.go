package main

import (
	"go-crud/config"
	"go-crud/routers"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()

}

func main() {
	app := fiber.New()
	routers.SetUpRoutes(app)

	port := os.Getenv("PORT")
	app.Listen(port)
}
