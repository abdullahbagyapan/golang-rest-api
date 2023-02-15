package service

import (
	"encoding/json"
	"petshop/dbConfig"
	"petshop/entity"
)

func GetAnimals() []entity.Animal {

	var animals []entity.Animal
	dbConfig.Database.Find(&animals)

	return animals
}

func AddAnimal(requestBody []byte) (entity.Animal, error) {
	animal := new(entity.Animal)

	err := json.Unmarshal([]byte(requestBody), &animal) // setter to dog object's values

	if err != nil {
		return *animal, err
	}

	dbConfig.Database.Create(&animal)

	return *animal, nil
}
