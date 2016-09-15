package main

import (
	"github.com/fasthall/cappa/docker"
	"github.com/fasthall/cappa/redis"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func tasksGET(c *gin.Context) {
	task := c.Param("task")
	image := redis.Get("tasks", task)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"task":   task,
		"image":  image,
	})
}

func tasksPOST(c *gin.Context) {
	image := c.PostForm("image")
	uuid, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot generate new key",
		})
		return
	}
	key := uuid.String()
	redis.Set("tasks", key, image)
	docker.Pull(image)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"uuid":   key,
		"image":  image,
	})
}

func tasksDELETE(c *gin.Context) {
	task := c.Param("task")
	redis.Del("tasks", task)
}
