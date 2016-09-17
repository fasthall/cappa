package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	router := gin.Default()

	// tasks
	router.GET("/tasks/:task", tasksGET)
	router.POST("/tasks", tasksPOST)
	router.DELETE("/tasks/:task", tasksDELETE)

	// trigger
	router.POST("/trigger", triggerPOST)

	// logs
	router.GET("/logs/:log", logsGET)

	router.Run(":8080")
	fmt.Println("Cappa server is running on port 8080.")
	fmt.Println("^C to terminate.")
}
