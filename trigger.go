package main

import (
	"fmt"
	"github.com/fasthall/cappa/docker"
	"github.com/fasthall/cappa/redis"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
	"io"
	"net/http"
	"os"
)

func triggerPOST(c *gin.Context) {
	// Find the image
	task := c.Query("task")
	image := redis.Get("tasks", task)
	if image == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"task":   task,
			"image":  "not found",
			"result": "",
		})
	}

	// Mount a file if specified
	file, header, err := c.Request.FormFile("upload")
	if header != nil {
		filename := header.Filename
		fmt.Println(header.Filename)
		out, err := os.Create("./tmp/" + filename)
		if err != nil {
			panic(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
	}

	// Create and start the container
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
