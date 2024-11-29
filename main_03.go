package main

import (
	EmployeeController "backend/api/controller/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Employee API Methods
	router.GET("/employee", EmployeeController.GetEmployee)           // GET
	router.GET("/employee/:id", EmployeeController.GetEmployeeBYID)   // GET by ID
	router.POST("/employee", EmployeeController.POSTEmployee)         // POST
	router.PUT("/employee/:id", EmployeeController.PUTEmployee)       // PUT
	router.DELETE("/employee/:id", EmployeeController.DELETEEmployee) // DELETE

	router.Run() // Listen and serve on 0.0.0.0:8080 (or "localhost:8080" for Windows)
}
