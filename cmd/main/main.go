package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/udodinho/bookstore/pkg/models"
	"github.com/udodinho/bookstore/pkg/routes"
)

func main() {
	
	port := ":3000"
	
	r := routes.Repository{
		DB: models.DB,
	}
	
	app := fiber.New()
	r.SetupRoutes(app)
	fmt.Println("Server started listening on port", port)
	app.Listen(port)
}
