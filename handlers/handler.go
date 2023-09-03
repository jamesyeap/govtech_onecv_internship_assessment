package handlers

import (
	"log"
	"teacher_api/models"

	"gorm.io/gorm"
)

type handler struct {
	DB  *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

// TODO: shift these helper functions into their own respective files in their own directory /////
func (h handler) FindStudentByEmail(email string) *models.Student {
	var student models.Student;

	if result := h.DB.Where("email = ?", email).First(&student); result.Error != nil {
		// TODO: log and return an appropriate error
		log.Println(result.Error)
	}

	return &student
}
func (h handler) FindTeacherByEmail(email string) *models.Teacher {
	var teacher models.Teacher;

	if result := h.DB.Where("email = ?", email).First(&teacher); result.Error != nil {
		// TODO: log and return an appropriate error
		log.Println(result.Error)
	}

	return &teacher
}

func (h handler) FindStudentsRegisteredToTeacherByEmail(teacherEmail string) *[]models.Student {
	var students []models.Student

	teacher := h.FindTeacherByEmail(teacherEmail)

	// if teacher does not exist, return an empty array
	if (teacher == nil) {
		return &students
	}
	
	h.DB.Model(&teacher).Association("Students").Find(&students)

	return &students
}

func (h handler) FindUnsuspendedStudentsRegisteredToTeacherByEmail(teacherEmail string) *[]models.Student {
	registeredStudents := h.FindStudentsRegisteredToTeacherByEmail(teacherEmail)

	var students []models.Student
	for _, student := range *registeredStudents {
		if student.IsSuspended == false {
			students = append(students, student)
		}
	}

	return &students
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////