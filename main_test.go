package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/danitrod/go-gin-server/controllers"
	"github.com/danitrod/go-gin-server/database"
	"github.com/danitrod/go-gin-server/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func createMockStudent() uint {
	student := models.Student{Name: "John Test", RG: "123456789", CPF: "12345678910"}
	database.DB.Create(&student)
	return student.ID
}

func deleteMockStudent(id uint) {
	var student models.Student
	database.DB.Delete(&student, id)
}

func TestGetStudents(t *testing.T) {
	database.ConnectToDB()
	mockId := createMockStudent()
	defer deleteMockStudent(mockId)
	r := SetupTestRoutes()
	r.GET("/students", controllers.GetStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "status code should be 200")

	var students []models.Student
	json.Unmarshal(res.Body.Bytes(), &students)

	assert.LessOrEqual(t, 1, len(students))
}

func TestGetStudentByCPF(t *testing.T) {
	database.ConnectToDB()
	mockId := createMockStudent()
	defer deleteMockStudent(mockId)

	r := SetupTestRoutes()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	req, _ := http.NewRequest("GET", "/students/cpf/12345678910", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var student models.Student
	json.Unmarshal(res.Body.Bytes(), &student)

	assert.NotNil(t, student)
}

func TestGetStudentById(t *testing.T) {
	database.ConnectToDB()
	mockId := createMockStudent()
	defer deleteMockStudent(mockId)

	r := SetupTestRoutes()
	r.GET("/students/:id", controllers.GetStudent)
	searchPath := "/students/" + strconv.Itoa(int(mockId))
	req, _ := http.NewRequest("GET", searchPath, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var mockStudent models.Student
	json.Unmarshal(res.Body.Bytes(), &mockStudent)

	assert.Equal(t, "John Test", mockStudent.Name)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectToDB()
	mockId := createMockStudent()

	r := SetupTestRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	searchPath := "/students/" + strconv.Itoa(int(mockId))
	req, _ := http.NewRequest("DELETE", searchPath, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestUpdateStudent(t *testing.T) {
	database.ConnectToDB()
	mockId := createMockStudent()
	defer deleteMockStudent(mockId)

	r := SetupTestRoutes()
	r.PATCH("/students/:id", controllers.UpdateStudent)
	searchPath := "/students/" + strconv.Itoa(int(mockId))
	body := []byte(`{"Name": "John Update"}`)
	req, _ := http.NewRequest("PATCH", searchPath, bytes.NewBuffer(body))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var mockStudent models.Student
	json.Unmarshal(res.Body.Bytes(), &mockStudent)

	assert.Equal(t, "John Update", mockStudent.Name)
}
