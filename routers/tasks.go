package routers

import (
	"net/http"

	"github.com/fasthall/cappa/docker"
	"github.com/fasthall/cappa/redis"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
)

// TasksGET replies the task info in the redis to the client
func TasksGET(c *gin.Context) {
	task := c.Param("task")
	value, err := redis.Hgetall("task", task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when reading from Redis",
		})
		return
	}
	if value != nil {
		c.JSON(http.StatusOK, value)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't find the task",
		})
	}
}

// TasksPOST inserts a new hash into Redis and ask Docker daemon to create and start a container
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
	_, err = redis.Hmset("task", key, key, image, "pulling")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when writing to Redis",
		})
		return
	}
	go docker.Pull(key, image)

	c.JSON(http.StatusOK, gin.H{
		"uuid":   key,
		"image":  image,
		"status": "pulling",
	})
}

// TasksDELETE deletes a posted task
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
