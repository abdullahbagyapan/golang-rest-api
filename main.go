package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"petshop/controller"
	"petshop/dbConfig"
)

func main() {

	app := fiber.New()

	dbConfig.Connect()

	app.Get("/animals", controller.GetAnimals)
	app.Post("/animals", controller.AddAnimal)
	app.Put("/animals", controller.UpdateAnimal)
	app.Delete("/animals/:id", controller.DeleteAnimal)git 

	log.Fatal(app.Listen(":8080"))

}
