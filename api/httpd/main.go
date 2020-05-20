package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
) 

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fmt.Println("Serving at port: 5000")
	router.Run("0.0.0.0:5000")
}