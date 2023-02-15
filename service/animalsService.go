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

func UpdateAnimal(requestBody []byte) (entity.Animal, error) {
	animalDTO := new(entity.Animal)
	err := json.Unmarshal([]byte(requestBody), &animalDTO)

	if err != nil {
		return *animalDTO, err
	}

	var animal entity.Animal
	result := dbConfig.Database.First(&animal, "id = ?", animalDTO.ID) // id query

	if result.Error != nil {
		return *animalDTO, result.Error
	}

	result = dbConfig.Database.Model(&entity.Animal{}).Where("id = ?", animalDTO.ID).Updates(animalDTO) // update entity

	if result.Error != nil {
		return *animalDTO, result.Error
	}
	return *animalDTO, nil
}

func DeleteAnimal(id int) error {

	result := dbConfig.Database.Model(&entity.Animal{}).Where("id = ?", id).Delete(&entity.Animal{}) // update entity

	if result.Error != nil {
		return result.Error
	}

	return nil
}
