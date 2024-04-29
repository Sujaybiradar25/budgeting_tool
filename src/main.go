package main

import (
	"github.com/gofiber/fiber/v2"
	"kniru/src/controller"
	"kniru/src/controller/kniru"
	"log"
)

func main() {
	app := fiber.New()
	api := app.Group("/kniru")
	api.Get("/hello", kniru.SimpleGet)
	api.Post("/createBudget", controller.CreateBudget)
	api.Get("/budget/:id", controller.GetBudget)
	api.Post("/addTransaction", controller.AddTransaction)
	api.Post("/newUser", controller.CreateUser)
	log.Fatal(app.Listen(":8000"))
}
