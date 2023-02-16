package service

import (
	"encoding/json"
	"errors"
	"petshop/dbConfig"
	"petshop/entity"
)

func GetAnimals(id int) (value interface{}, er error) {

	if id != -1 { // default param value = -1
		var animal entity.Animal
		result := dbConfig.Database.First(&animal, "id = ?", id)

		if result.Error != nil {
			return nil, result.Error
		}

		return animal, nil

	}
	var animals []entity.Animal
	dbConfig.Database.Find(&animals)

	return animals, nil

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

func UpdateAnimal(id int, requestBody []byte) (interface{}, error) {

	animalDTO := new(entity.Animal)
	err := json.Unmarshal([]byte(requestBody), &animalDTO) // assign

	if err != nil {
		return nil, err
	}

	var animal entity.Animal
	result := dbConfig.Database.First(&animal, "id = ?", id) // id query

	if result.Error != nil {
		return nil, result.Error
	}

	result = dbConfig.Database.Model(&entity.Animal{}).Where("id = ?", id).Updates(animalDTO) // update entity

	if result.Error != nil {
		return nil, result.Error
	}

	return *animalDTO, nil
}

func DeleteAnimal(id int) error {

	if id != -1 {

		result := dbConfig.Database.Model(&entity.Animal{}).Where("id = ?", id).Delete(&entity.Animal{}) // delete entity

		if result.Error != nil {
			return result.Error
		}
		return nil
	}
	return errors.New("invalid param")

}
