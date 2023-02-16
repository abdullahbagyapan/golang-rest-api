package controller

import (
	fiber "github.com/gofiber/fiber/v2"
	"net/http"
	"petshop/service"
)

func GetAnimals(c *fiber.Ctx) error {
	param, _ := c.ParamsInt("id", -1) // default param value ->> -1

	body, err := service.GetAnimals(param)

	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

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

	id, err := c.ParamsInt("id", -1)
	requestBody := c.Request().Body()

	animal, err := service.UpdateAnimal(id, requestBody)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Animal did not update")
	}

	return c.Status(http.StatusOK).JSON(animal)

}

func DeleteAnimal(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", -1)

	err = service.DeleteAnimal(id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Animal did not delete")
	}

	return c.SendStatus(http.StatusOK)
}
