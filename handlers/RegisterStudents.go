package handlers

import (
	"net/http"
	"teacher_api/models"

	"github.com/gin-gonic/gin"
)

type RegisterStudentsParam struct {
	TeacherEmail  string   `json:"teacher"`
	StudentEmails []string `json:"students"`
}

func (h handler) RegisterStudentsHandler(c *gin.Context) {
	param, issues := parseJsonBodyForRegisterStudents(c)
	if len(issues) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, ParamError{issues})
		return
	}

	// fetch the teacher record using the email
	teacher := h.FindTeacherByEmail(param.TeacherEmail)

	for _, studentEmail := range param.StudentEmails {
		student := h.FindStudentByEmail(studentEmail)

		if student == nil {
			// for now, just add the student to the DB if doesn't exist
			h.DB.Create(models.Student{Email: studentEmail})
			student = h.FindStudentByEmail(studentEmail)
		}

		// add associations
		h.DB.Model(&student).Association("Teachers").Append(teacher)
		h.DB.Model(&teacher).Association("Students").Append(student)
	}

	c.Status(http.StatusNoContent)
}

func parseJsonBodyForRegisterStudents(c *gin.Context) (RegisterStudentsParam, map[string]interface{}) {
	var param RegisterStudentsParam

	issues := make(map[string]interface{})

	if err := c.BindJSON(&param); err != nil {
		issues["general"] = ValidationIssue{"", err.Error()}

		return param, issues
	}

	validateParamsReceivedForRegisterStudents(param, issues)

	return param, issues
}

func validateParamsReceivedForRegisterStudents(param RegisterStudentsParam, issues map[string]interface{}) {
	if param.TeacherEmail == "" {
		issues["teacher"] = ValidationIssue{"required", ""}
	} else if !IsValidEmailFormat(param.TeacherEmail) {
		issues["teacher"] = ValidationIssue{"invalid email format", ""}
	}

	if param.StudentEmails == nil || len(param.StudentEmails) == 0 {
		issues["students"] = ValidationIssue{"required", ""}
	} else if !allStudentEmailsAreValid(param.StudentEmails) {
		issues["students"] = ValidationIssue{"invalid email format found in email list", ""}
	}
}

func allStudentEmailsAreValid(studentEmails []string) bool {
	var allEmailsValid = true

	for _, email := range studentEmails {
		if !IsValidEmailFormat(email) {
			allEmailsValid = false
			break
		}
	}

	return allEmailsValid
}
