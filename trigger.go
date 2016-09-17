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

	event, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "couldn't generate new event ID",
		})
	}
	eid := event.String()

	// Mount a file if specified
	pwd, err := os.Getwd()
	file, header, err := c.Request.FormFile("upload")
	env := []string{}
	if header != nil {
		filename := header.Filename
		fmt.Println(header.Filename)
		os.Mkdir(pwd+"/tmp", 0755)
		os.Mkdir(pwd+"/tmp/"+eid, 0755)
		out, err := os.Create(pwd + "/tmp/" + eid + "/" + filename)
		if err != nil {
			panic(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		env = append(env, "PAYLOAD=/payload/"+filename)
	}

	// Create and start the container
	cid := docker.Create(image, []string{pwd + "/tmp/" + eid + ":/payload"}, env)
	docker.Start(cid)
	logs := docker.Logs(cid)
	redis.Set("logs", eid, logs)
	c.JSON(http.StatusOK, gin.H{
		"task":  task,
		"image": image,
		"event": eid,
	})
}
