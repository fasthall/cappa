package main

import (
	"fmt"

	"github.com/fasthall/cappa/mq"
	"github.com/gin-gonic/gin"
)

func main() {
	// listen to message queue
	mqc, _ := mq.NewMQ()
	go mqc.Listen()

	router := gin.Default()

	// tasks
	router.GET("/tasks/:task", tasksGET)
	router.POST("/tasks", tasksPOST)
	router.DELETE("/tasks/:task", tasksDELETE)

	// rules
	router.POST("/rules", rulesPOST)

	// trigger
	router.POST("/trigger", triggerPOST)

	// logs
	router.GET("/logs/:log", logsGET)

	router.Run(":8080")
	fmt.Println("Cappa server is running on port 8080.")
	fmt.Println("^C to terminate.")
}
