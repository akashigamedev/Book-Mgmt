package routes

import (
	"github.com/akashigamedev/book-mgmt/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

var RegistrBookStoreRoutes = func(app *fiber.App) {
	app.Post("/book", controllers.CreateBook)
	app.Get("/book", controllers.GetBook)
	app.Get("/book/:bookId", controllers.GetBookById)

	app.Put("/book/:bookId", controllers.UpdateBook)
	app.Delete("/book/:bookId", controllers.DeleteBook)
}
