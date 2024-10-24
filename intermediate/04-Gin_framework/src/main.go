package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name"`
}

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func queryHandler(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "World"
	}
	c.String(http.StatusOK, fmt.Sprintf("Hello, %s!", name))
}

func jsonHandler(c *gin.Context) {
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello, %s!", person.Name),
	})
}

func main() {
	router := gin.Default()
	router.GET("/hello", helloHandler)
	router.GET("/query", queryHandler)
	router.POST("/json", jsonHandler)

	fmt.Println("Server starting on port 8080...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
