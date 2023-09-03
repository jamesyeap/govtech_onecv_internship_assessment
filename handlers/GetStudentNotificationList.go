package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetStudentNotificationListParam struct {
	TeacherEmail string `json:"teacher"`
	NotificationMessage string `json:"notification"`
}

type GetStudentNotificationListResponse struct {
	Students []string `json:"recipients"`
}

func (h handler) GetStudentNotificationList(c *gin.Context) {
	param, issues := parseJsonBodyForGetStudentNotificationList(c)
	if (len(issues) > 0) {
		c.AbortWithStatusJSON(http.StatusBadRequest, ParamError{ issues })
		return
	}

	unsuspendedRegisteredStudents := h.FindUnsuspendedStudentsRegisteredToTeacherByEmail(param.TeacherEmail)

	studentEmailsMentioned := findAllMentionedStudentsInMessage(param.NotificationMessage)

	emailSet := make(map[string]bool)
	for _, student := range *unsuspendedRegisteredStudents {
		emailSet[student.Email] = true
	}
	for _, studentEmail := range studentEmailsMentioned {
		emailSet[studentEmail] = true
	}

	notificationList := make([]string, 0)
	for studentEmail := range emailSet {
		notificationList = append(notificationList, studentEmail)
	}

	c.JSON(http.StatusOK, GetStudentNotificationListResponse{
		Students: notificationList,
	})
}

func parseJsonBodyForGetStudentNotificationList(c *gin.Context) (GetStudentNotificationListParam, map[string]interface{}) {
	var param GetStudentNotificationListParam;

	issues := make(map[string]interface{})

	if err := c.BindJSON(&param); err != nil {
		issues["general"] = ValidationIssue{"", err.Error()}

		return param, issues
	}

	validateParamsReceivedForGetStudentNotificationList(param, issues)

	return param, issues;
}

func validateParamsReceivedForGetStudentNotificationList(param GetStudentNotificationListParam, issues map[string]interface{}) {
	if (param.TeacherEmail == "") {
		issues["teacher"] = ValidationIssue{"required",""}
	} else if (!IsValidEmailFormat(param.TeacherEmail)) {
		issues["teacher"] = ValidationIssue{"invalid email format", ""}
	}
}

func findAllMentionedStudentsInMessage(message string) []string {

	mentionedStudentEmails := make([]string, 0)

	tokens := strings.Fields(message)

	for _, token := range filterStringList(tokens) {
		mentionedStudentEmails = append(mentionedStudentEmails, trimLeftChar(token))
	}

	return mentionedStudentEmails
}

func filterStringList(stringList []string) (filteredStringList []string) {
	for _, s := range stringList {
		if (IsValidMentionedFormat(s)) {
			filteredStringList = append(filteredStringList, s)
		}
	}

	return;
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}