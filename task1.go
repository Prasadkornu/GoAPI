package main

import (
	"github.com/gin-gonic/gin"
)

func first(c *gin.Context) {
	c.JSON(200, gin.H{
		"messge": "Aptroid pvt ltd",
	})
}
func second(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hyderabad",
	})
}
func main() {
	router := gin.Default() // Creates a gin router with default middleware
	router.GET("/first", first)
	router.GET("/second", second)
	router.Run(":8080")
}
