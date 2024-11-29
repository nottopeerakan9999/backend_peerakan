package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main02() {
  router := gin.Default()
  router.GET("/employee", getEmployee)
    
  router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getEmployee(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
  "message": "Employee GET Method!",
  })
}

func PUTEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee PUT Method!",
    })
  }

  func POSTEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee POST Method!",
    })
  }

  func DELETEEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee DELETE Method!",
    })
  }