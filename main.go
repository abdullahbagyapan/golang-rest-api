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

	app.Get("/animals/:id<int>?", controller.GetAnimals) // paramName , paramType , optionalParam
	app.Post("/animals", controller.AddAnimal)
	app.Put("/animals", controller.UpdateAnimal)
	app.Delete("/animals/:id", controller.DeleteAnimal)

	log.Fatal(app.Listen(":8080"))

}
