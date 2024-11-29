package employee

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

//GET
func GetEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET ALL Method!",
    })
}

//GET  BY ID
func GetEmployeeBYID(c *gin.Context) {
	id :=c.Param ("id")
    c.JSON(http.StatusOK, gin.H{
    "message": id,
    })
}

//POST
func POSTEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET Method!",
    })
}

//PUT
func PUTEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET Method!",
    })
}

//DELETE
func DELETEEmployee(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET Method!",
    })
}