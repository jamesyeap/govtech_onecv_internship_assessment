package db

import (
	"log"

	"teacher_api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const DB_URL = "postgres://pg:pass@localhost:8081/crud"
const DB_URL = "postgres://pg:pass@database:5432/crud"

func Init() *gorm.DB {

    db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
    if err != nil {
        log.Fatalln(err)
    }

    // TODO: delete this after testing ///////
    db.Migrator().DropTable(&models.Teacher{})
    db.Migrator().DropTable(&models.Student{})
    //////////////////////////////////////////

    db.AutoMigrate(&models.Teacher{}, &models.Student{})

    // TODO: delete this after testing ///
    addDefaultTeachersToDB(db)
    addDefaultStudentsToDB(db)
    //////////////////////////////////////

    return db
}

var DEFAULT_TEACHERS = []models.Teacher {
    { Email: "teacherken@gmail.com" },
    { Email: "teacherjoe@gmail.com" },
}

func addDefaultTeachersToDB(db *gorm.DB) {
    db.Create(DEFAULT_TEACHERS)
}

var DEFAULT_STUDENTS = []models.Student {
    { Email: "studentjon@gmail.com" },
    { Email: "studenthon@gmail.com" },
    { Email: "commonstudent1@gmail.com" },
    { Email: "commonstudent2@gmail.com" },
    { Email: "studentmary@gmail.com" },
}

func addDefaultStudentsToDB(db *gorm.DB) {
    db.Create(DEFAULT_STUDENTS)
}
