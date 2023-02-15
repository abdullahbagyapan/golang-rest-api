package dbConfig

import (
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"petshop/entity"
)

var Database *gorm.DB
var DATABASE_URI string = "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

func Connect() error {
	var err error

	Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entity.Dog{})

	return nil
}
