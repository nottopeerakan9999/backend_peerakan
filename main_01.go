package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main01() {
	r := gin.Default()

	//------------------------------EMPLOYEE METHOD---------------------------------
	// GET Method for Employee
	r.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET EMPLOYEE METHOD",
		})
	})

	// POST Method for Employee
	r.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST EMPLOYEE METHOD",
		})
	})

	// PUT Method for Employee
	r.PUT("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT EMPLOYEE METHOD",
		})
	})

	// DELETE Method for Employee
	r.DELETE("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE EMPLOYEE METHOD",
		})
	})

	//------------------------CUSTOMER METHOD----------------------------------------------
	// GET Method for Customer
	r.GET("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET CUSTOMER METHOD",
		})
	})

	// POST Method for Customer
	r.POST("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST CUSTOMER METHOD",
		})
	})

	// PUT Method for Customer
	r.PUT("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT CUSTOMER METHOD",
		})
	})

	// DELETE Method for Customer
	r.DELETE("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE CUSTOMER METHOD",
		})
	})

	// Start the server
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
