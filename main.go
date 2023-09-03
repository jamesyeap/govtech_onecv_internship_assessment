package main

import (
	"teacher_api/db"
	"teacher_api/handlers"

	"github.com/gin-gonic/gin"
)

const HOST_ADDRESS = ":8080"

func main() {
	DB := db.Init()
	h := handlers.New(DB)

	router := gin.Default()

	router.GET("/api/commonstudents", h.GetCommonStudentsHandler)
	router.POST("/api/register", h.RegisterStudentsHandler)
	router.POST("/api/suspend", h.SuspendStudentHandler)
	router.POST("/api/retrievefornotifications", h.GetStudentNotificationListHandler)

	router.Run(HOST_ADDRESS)
}
