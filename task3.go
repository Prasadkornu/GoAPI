package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// items slice to store data about items.
var items = []Item{
	{ID: 1, Name: "Item 1", Price: 100},
	{ID: 2, Name: "Item 2", Price: 150},
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.GET("/items/:id", getItemByID)
	router.POST("/items", createItem)
	router.PUT("/items/:id", updateItem)
	router.DELETE("/items/:id", deleteItem)
	router.Run(":8080")
}

func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func getItemByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func createItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	newItem.ID = len(items) + 1
	items = append(items, newItem)

	c.JSON(http.StatusCreated, newItem)
}

func updateItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var updatedItem Item
	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	for i, item := range items {
		if item.ID == id {
			updatedItem.ID = id
			items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func deleteItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
		
		}
	}
}
