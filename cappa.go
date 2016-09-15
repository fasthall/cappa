package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// tasks
	router.GET("/tasks/:image", tasksGET)
	router.GET("/tasks/:image/*tail", tasksGET)
	router.POST("/tasks", tasksPOST)
	router.DELETE("/tasks/:image", tasksDELETE)
	router.DELETE("/tasks/:image/*tail", tasksDELETE)

	// trigger
	router.POST("/trigger", triggerPOST)

	// logs
	router.GET("/logs/:log", logsGET)

	router.Run(":8080")
	fmt.Println("Cappa server is running on port 8080.")
	fmt.Println("^C to terminate.")
}
