package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akashigamedev/book-mgmt/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	// CSP middleware to prevent Content-Security-Policy issue
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Security-Policy", "default-src *; connect-src *;")
		return c.Next()
	})

	routes.RegistrBookStoreRoutes(app)

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
	fmt.Println("Server started at port 8080")
}
