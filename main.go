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

	app.Get("/dogs", controller.GetDogs)
	app.Post("/dogs", controller.AddDog)

	log.Fatal(app.Listen(":8080"))

}
