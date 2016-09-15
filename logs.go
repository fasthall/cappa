package main

import (
	"github.com/fasthall/cappa/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func logsGET(c *gin.Context) {
	log := c.Param("log")
	content := redis.Get("logs", log)
	if content == "" {
		c.String(http.StatusBadRequest, "Log not found")
	} else {
		c.String(http.StatusOK, content)
	}
}
