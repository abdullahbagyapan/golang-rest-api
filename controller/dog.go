package controller

import (
	fiber "github.com/gofiber/fiber/v2"
	"net/http"
	"petshop/service"
)

func GetDogs(c *fiber.Ctx) error {
	body := service.GetDogs()

	return c.Status(http.StatusOK).JSON(body)
}

func AddDog(c *fiber.Ctx) error {
	requestBody := c.Request().Body()

	dog, err := service.AddDog(requestBody)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Dog did not add")
	}

	return c.Status(201).JSON(dog)
}
