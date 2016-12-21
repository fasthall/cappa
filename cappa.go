package main

import (
	"fmt"

	"github.com/fasthall/cappa/mq"
	"github.com/fasthall/cappa/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// listen to message queue
	mqc, err := mq.NewMQ()
	if err != nil {
		panic(err)
	}
	go mqc.Listen()

	router := gin.Default()

	// tasks
	router.GET("/tasks/:task", routers.TasksGET)
	router.POST("/tasks", routers.TasksPOST)
	router.DELETE("/tasks/:task", routers.TasksDELETE)

	// rules
	router.POST("/rules", routers.RulesPOST)

	// trigger
	router.POST("/trigger", routers.TriggerPOST)

	// logs
	router.GET("/logs/:log", routers.LogsGET)

	fmt.Println("Cappa server is running on port 8080.")
	fmt.Println("^C to terminate.")
	router.Run(":8080")

}
