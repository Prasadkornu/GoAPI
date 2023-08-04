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




////
package main

import (
	// "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Address struct {
	City    string `json:"city" gorm:"not null"`
	State   string `json:"state" gorm:"not null"`
	Zip     int    `json:"zip" gorm:"not null"`
	Country string `json:"country" gorm:"not null"`
}

var add []Address

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:2033@tcp(127.0.0.1:3306)/customer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db
}
func MigrateDatabase() {
	DB.AutoMigrate(&Address{})
}
func Get(c *gin.Context) {
	var address []Address
	if err := DB.Find(&address).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}
	c.JSON(http.StatusOK, address)
}
func GetAddress(c *gin.Context) {
	var address Address
	addressID, _ := strconv.Atoi(c.Param("id"))

	if err := DB.First(&address, addressID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}

	c.JSON(http.StatusOK, address)
}

func CreateAddress(c *gin.Context) {
	if len(add) > 0 {
		c.String(200, "Already Found")
		return
	}
	var newAddress Address
	if err := c.ShouldBindJSON(&newAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := DB.Create(&newAddress).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create address"})
		return
	}

}

func main() {
	r := gin.Default()
	ConnectDatabase()
	MigrateDatabase()

	r.GET("/get", Get)
	r.GET("/addresses/:id", GetAddress)
	r.POST("/post", CreateAddress)
	// r.PUT("/addresses/:id", UpdateAddress)
	// r.DELETE("/addresses/:id", DeleteAddress)
	r.Run(":8080")
}
