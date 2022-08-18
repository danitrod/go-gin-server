package controllers

import (
	"net/http"

	"github.com/danitrod/go-gin-server/database"
	"github.com/danitrod/go-gin-server/models"
	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func GetStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(200, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	database.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"msg": "Student deleted"})
}

func UpdateStudent(c *gin.Context) {
	var toUpdateStudent, existingStudent models.Student
	id := c.Params.ByName("id")
	if err := c.ShouldBindJSON(&toUpdateStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.First(&existingStudent, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	database.DB.Model(&existingStudent).UpdateColumns(toUpdateStudent)
	c.JSON(http.StatusOK, existingStudent)
}

func GetStudentByCPF(c *gin.Context) {
	cpf := c.Params.ByName("cpf")
	var student models.Student
	if err := database.DB.Where(&models.Student{CPF: cpf}).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(200, student)
}
