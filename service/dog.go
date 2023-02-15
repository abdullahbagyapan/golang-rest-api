package service

import (
	"encoding/json"
	"petshop/dbConfig"
	"petshop/entity"
)

func GetDogs() []entity.Dog {

	var dogs []entity.Dog
	dbConfig.Database.Find(&dogs)

	return dogs
}

func AddDog(requestBody []byte) (entity.Dog, error) {
	dog := new(entity.Dog)

	err := json.Unmarshal([]byte(requestBody), &dog) // setter to dog object's values

	if err != nil {
		return *dog, err
	}

	dbConfig.Database.Create(&dog)

	return *dog, nil
}
