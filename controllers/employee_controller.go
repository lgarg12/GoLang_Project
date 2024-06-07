package controllers

import (
	"context"
	"net/http"
	
	"GoLang_Project/models" 
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeController struct {
	Collection *mongo.Collection
}

func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
    var employee models.Employee
    if err := c.ShouldBindJSON(&employee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := ec.Collection.InsertOne(context.Background(), employee)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Employee created successfully", "employee": employee})
}

func (ec *EmployeeController) FindEmployeeByID(c *gin.Context) {
    empID := c.Param("id")


    empIDPrimitive, err := primitive.ObjectIDFromHex(empID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _id format"})
        return
    }

    var emp models.Employee
    newerr := ec.Collection.FindOne(context.Background(), bson.D{{"_id", empIDPrimitive}}).Decode(&emp)
    if newerr != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }

    c.JSON(http.StatusOK, emp)
}

func (ec *EmployeeController) ListAllEmployees(c *gin.Context) {
    cursor, err := ec.Collection.Find(context.Background(), bson.D{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cursor.Close(context.Background())

    var employees []models.Employee
    err = cursor.All(context.Background(), &employees)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, employees)
}

func (ec *EmployeeController) DeleteEmployeeByID(c *gin.Context) {
    empID := c.Param("id")

    empIDPrimitive, err := primitive.ObjectIDFromHex(empID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid _id format"})
        return
    }

    _, newerr := ec.Collection.DeleteOne(context.Background(), bson.D{{"_id", empIDPrimitive}})
    if newerr != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

