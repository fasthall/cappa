package main

import (
	"github.com/fasthall/cappa/docker"
	"github.com/fasthall/cappa/redis"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func triggerPOST(c *gin.Context) {
	task := c.Query("task")
	image := redis.Get("tasks", task)
	if image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"task":   task,
			"image":  "not found",
			"result": "",
		})
	}
	cid := docker.Create(image)
	docker.Start(cid)
	logs := docker.Logs(cid)
	logid, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot generate new key",
		})
		return
	}
	redis.Set("logs", logid.String(), logs)
	c.JSON(http.StatusOK, gin.H{
		"task":  task,
		"image": image,
		"event": logid.String(),
	})
}
