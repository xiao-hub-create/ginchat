package service

import (
	"github.com/gin-gonic/gin"
)

func GETIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
