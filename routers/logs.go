package routers

import (
	"net/http"

	"github.com/fasthall/cappa/redis"
	"github.com/gin-gonic/gin"
)

func LogsGET(c *gin.Context) {
	log := c.Param("log")
	content, err := redis.Get("logs", log)
	if err != nil {
		panic(err)
	}
	if content == "" {
		c.String(http.StatusBadRequest, "Log not found")
	} else {
		c.String(http.StatusOK, content)
	}
}
