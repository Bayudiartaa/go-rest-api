package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Bayudiartaa/go-rest-api/models"
	"github.com/Bayudiartaa/go-rest-api/controllers"
)

func main(){
    models.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", bookController.Index)
	book.Get("/:id", bookController.Show)
	book.Post("/", bookController.Create)
	book.Put("/:id", bookController.Update)	
	book.Delete("/:id", bookController.Delete)
	
	app.Listen(":8080")
} 