package routes

import (
	"github.com/gin-gonic/gin"
	"GoLang_Project/controllers"
)

// DefineEmployeeRoutes defines routes for EmployeeController methods
func DefineEmployeeRoutes(r *gin.Engine, ec *controllers.EmployeeController) {
	r.POST("/employees", ec.CreateEmployee)
	r.GET("/employees/:id", ec.FindEmployeeByID)
	r.GET("/employees", ec.ListAllEmployees)
	r.DELETE("/employees/:id", ec.DeleteEmployeeByID)
}
