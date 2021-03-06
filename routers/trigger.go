package routers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/fasthall/cappa/docker"
	"github.com/fasthall/cappa/redis"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
)

// TriggerPOST triggers a posted task
func TriggerPOST(c *gin.Context) {
	// Find the image
	task := c.Query("task")
	value, err := redis.Hgetall("task", task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when reading from Redis",
		})
		return
	}
	if value == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't find the task",
		})
		return
	}
	if value["status"] != "downloaded" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The container image of the task has not been downloaded yet",
		})
		return
	}

	event, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "couldn't generate new event ID",
		})
		return
	}
	eid := event.String()

	// read the payload
	file, header, err := c.Request.FormFile("upload")
	go createAndStart(file, header, eid, value["image"])

	c.JSON(http.StatusOK, gin.H{
		"task":  task,
		"image": value["image"],
		"event": eid,
	})
}

func createAndStart(file multipart.File, header *multipart.FileHeader, eid string, image string) {
	fmt.Println(header, eid, image)
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var env []string
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
}
