package main

import (
	"github.com/gin-gonic/gin"
)

var a = map[string]string{
	"name": "prash",
}

func task(c *gin.Context) {
	res := c.Query("type")
	response, ok := a[res]
	if !ok {
		c.JSON(404, gin.H{
			"error": "Key not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": response,
	})
}

func main() {
	router := gin.Default()
	router.GET("/task2", task)
	router.POST("/api/data", func(c *gin.Context) {
		// Check if the request method is POST
		if c.Request.Method == "POST" {
			c.JSON(200, gin.H{
				"message": "Data received successfully via POST request!",
			})
		} else {
			c.JSON(400, gin.H{
				"error": "Invalid request method",
			})
		}
	})
	router.Run()
}
