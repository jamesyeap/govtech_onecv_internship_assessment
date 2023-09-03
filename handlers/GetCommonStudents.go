package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetCommonStudentsResponse struct {
	Students []string `json:"students"`
}

func (h handler) GetCommonStudents(c *gin.Context) {

	queryParams := c.Request.URL.Query()
	teacherEmailList := queryParams["teacher"]

	teacherEmailSet := make(map[string]bool)
	for _, email := range teacherEmailList {
		teacherEmailSet[email] = true
	}

	issues := make(map[string]interface{})
	validateParamsReceivedForGetCommonStudents(teacherEmailSet, issues)
	if (len(issues) > 0) {
		c.AbortWithStatusJSON(http.StatusBadRequest, ParamError{ issues })
		return
	}

	var numberOfTeachers = len(teacherEmailSet)

	studentToTeacherMap := make(map[string][]string);

	for teacherEmail := range teacherEmailSet {
		registeredStudents := h.FindStudentsRegisteredToTeacherByEmail(teacherEmail)

		for i := 0; i < len(*registeredStudents); i++ {
			student := (*registeredStudents)[i]

			studentToTeacherMap[student.Email] = append(studentToTeacherMap[student.Email], teacherEmail)
		}
	}

	commonStudentEmailList := make([]string, 0)
	for studentEmail, registeredTeacherEmailList := range studentToTeacherMap {
		log.Println(len(registeredTeacherEmailList), numberOfTeachers)

		if (len(registeredTeacherEmailList) == numberOfTeachers) {
			commonStudentEmailList = append(commonStudentEmailList, studentEmail)
		}
	}

	c.JSON(
		http.StatusOK,
		GetCommonStudentsResponse{
			Students: commonStudentEmailList,
	})
}

func validateParamsReceivedForGetCommonStudents(teacherEmailSet map[string]bool, issues map[string]interface{}) {
	if (len(teacherEmailSet) == 0) {
		issues["teacher"] = ValidationIssue{"required", ""}
		return
	}

	for email := range teacherEmailSet {
		if (!IsValidEmailFormat(email)) {
			issues["teacher"] = ValidationIssue{"invalid email format", ""}
			return
		}
	}
}