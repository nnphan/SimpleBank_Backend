package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

func (pc *PingController) Ping(c *gin.Context) {
	name := c.DefaultQuery("name", "phan")
	uuid := c.Query("uuid")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping ....." + name,
		"uuid":    uuid,
		"users":   []string{"cr7", "m10"},
	})
}