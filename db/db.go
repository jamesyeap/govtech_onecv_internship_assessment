package db

import (
	"log"

	"teacher_api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//const DB_URL = "postgres://pg:pass@localhost:8081/crud"

const DB_URL = "postgres://pg:pass@database:5432/crud"

func Init() *gorm.DB {

	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Teacher{}, &models.Student{})

	return db
}
