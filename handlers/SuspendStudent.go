package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuspendStudentParam struct {
	StudentEmail string `json:"student"`
}

func (h handler) SuspendStudentHandler(c *gin.Context) {
	param, issues := parseJsonBodyForSuspendStudent(c)
	if len(issues) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, ParamError{issues})
		return
	}

	student := h.FindStudentByEmail(param.StudentEmail)

	if student.Email == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "student not found"})
		return
	} else if student.IsSuspended {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "student is already suspended"})
		return
	}

	student.IsSuspended = true
	h.DB.Save(student)

	c.Status(http.StatusNoContent)
}

func parseJsonBodyForSuspendStudent(c *gin.Context) (SuspendStudentParam, map[string]interface{}) {
	var param SuspendStudentParam

	issues := make(map[string]interface{})

	if err := c.BindJSON(&param); err != nil {
		issues["general"] = ValidationIssue{"", err.Error()}

		return param, issues
	}

	validateParamsReceivedForSuspendStudent(param, issues)

	return param, issues
}

func validateParamsReceivedForSuspendStudent(param SuspendStudentParam, issues map[string]interface{}) {
	if param.StudentEmail == "" {
		issues["student"] = ValidationIssue{"required", ""}
	} else if !IsValidEmailFormat(param.StudentEmail) {
		issues["student"] = ValidationIssue{"invalid email format", ""}
	}
}
