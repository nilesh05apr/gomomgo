package routes

import (
	"github.com/gofiber/fiber/v2"
	"gomomgo/controllers"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userID", controllers.GetAUser)
	app.Put("/user/:userID", controllers.EditAUser)
	app.Delete("/user/:userID", controllers.DeleteAUser)
	app.Get("/user", controllers.GetAUser)

}