package routes

import (
	"github.com/drsimplegraffiti/go_fibre_mvc/controllers" // replace
	"github.com/gofiber/fiber/v2"
)

func CatchphrasesRoute(route fiber.Router) {
	route.Get("/", controllers.GetAllCatchphrases)
	route.Get("/:id", controllers.GetCatchphrase)
	route.Post("/", controllers.AddCatchphrase)
	route.Put("/:id", controllers.UpdateCatchphrase)
	route.Delete("/:id", controllers.DeleteCatchphrase)
}
