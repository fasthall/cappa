package routers

import (
	"net/http"

	"github.com/fasthall/cappa/docker"
	"github.com/fasthall/cappa/redis"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
)

func TasksGET(c *gin.Context) {
	task := c.Param("task")
	image, err := redis.Get("tasks", task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when reading from Redis",
		})
		return
	}
	if image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"task":  task,
			"image": image,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"task":  task,
			"image": image,
		})
	}
}

func TasksPOST(c *gin.Context) {
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
		"uuid":  key,
		"image": image,
	})
}

func TasksDELETE(c *gin.Context) {
	task := c.Param("task")
	err := redis.Del("tasks", task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Task is deleted",
		})
	}
}
