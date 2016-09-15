package main

import (
	"github.com/fasthall/cappa/docker"
	"github.com/fasthall/cappa/redis"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func tasksGET(c *gin.Context) {
	image := c.Param("image")
	tail := c.Param("tail")
	if tail != "" {
		image = image + tail
	}
	exist := docker.Exist(image)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"image":  image,
		"exist":  exist,
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
	image := c.Param("image")
	tail := c.Param("tail")
	if tail != "" {
		image = image + tail
	}
	docker.Remove(image)
}
