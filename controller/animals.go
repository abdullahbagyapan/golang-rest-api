package controller

import (
	fiber "github.com/gofiber/fiber/v2"
	"net/http"
	"petshop/service"
)

func GetAnimals(c *fiber.Ctx) error {
	body := service.GetAnimals()

	return c.Status(http.StatusOK).JSON(body)
}

func AddAnimal(c *fiber.Ctx) error {
	requestBody := c.Request().Body()

	animal, err := service.AddAnimal(requestBody)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Animal did not add")
	}

	return c.Status(201).JSON(animal)
}