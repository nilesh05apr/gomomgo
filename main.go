package main

import (
	"github.com/gofiber/fiber/v2"
    "gomomgo/configs"
    "gomomgo/routes"
)

func main() {
    app := fiber.New()

    configs.ConnectDB()

    routes.UserRoute(app)


    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{"data": "Running"})
    })
  
    app.Listen(":3000")
}