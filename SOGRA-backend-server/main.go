package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	fmt.Println("Hello, World!")
	r := gin.Default()
	r.GET("/ping", pingHandler)
	r.Run()
}
