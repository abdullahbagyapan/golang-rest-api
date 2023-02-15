package controller

import (
	"fmt"
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

	return c.Status(http.StatusCreated).JSON(animal)
}

func UpdateAnimal(c *fiber.Ctx) error {

	requestBody := c.Request().Body()

	animal, err := service.UpdateAnimal(requestBody)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Animal did not update")
	}

	return c.Status(http.StatusOK).JSON(animal)

}

func DeleteAnimal(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fmt.Sprintf("There is no id with = %d", id))
	}

	err = service.DeleteAnimal(id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Animal did not delete")
	}

	return c.SendStatus(http.StatusOK)
}
