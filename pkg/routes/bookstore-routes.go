package routes

import (
	// "fmt"
	// "log"

	"github.com/gofiber/fiber/v2"
	// "github.com/joho/godotenv"
	"github.com/udodinho/bookstore/pkg/controllers"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func(r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api/v1/books")
	api.Get("/", controllers.GetAllBooks)
	api.Post("/", controllers.CreateBook)
	api.Get("/:id", controllers.GetBook)
	api.Put ("/:id", controllers.UpdateBook)
	api.Delete("/:id", controllers.DeleteBook)

}
