package routers

import (
	"net/http"

	"encoding/json"

	"fmt"

	"github.com/fasthall/cappa/docker"
	"github.com/fasthall/cappa/redis"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
)

func TasksGET(c *gin.Context) {
	task := c.Param("task")
	jsonValue, err := redis.Get("tasks", task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when reading from Redis",
		})
		return
	}
	var value map[string]string
	err = json.Unmarshal([]byte(jsonValue), &value)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when parsing value",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"uuid":   value["uuid"],
		"image":  value["image"],
		"status": value["status"],
	})
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
	value := map[string]string{"image": image, "status": "pulling", "uuid": key}
	jsonValue, _ := json.Marshal(value)
	redis.Set("tasks", key, string(jsonValue))
	go docker.Pull(key, image)

	c.JSON(http.StatusOK, gin.H{
		"uuid":   key,
		"image":  image,
		"status": "pulling",
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
