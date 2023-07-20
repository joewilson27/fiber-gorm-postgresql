package main

import (
	"log"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm-postgres/database"
	"gorm-postgres/routes"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/allbooks", routes.AllBooks)
	app.Post("/addbook", routes.AddBook)
	app.Post("/book", routes.Book)
	app.Put("/update", routes.Update)
	app.Delete("/delete", routes.Delete)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // 404 Not Found
	})

	var port int = 3000
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("started at localhost:%d", port)
	log.Fatal(app.Listen(addr))
}