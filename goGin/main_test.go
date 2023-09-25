package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/PhilipFelipe/golang-alura-course/controllers"
	"github.com/PhilipFelipe/golang-alura-course/database"
	"github.com/PhilipFelipe/golang-alura-course/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func createStudentMock() {
	student := models.Student{Name: "Test Student", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func deleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifyWelcomeEndpointResponseStatusCodeAndBody(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controllers.Welcome)
	req, _ := http.NewRequest("GET", "/felipe", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")
	responseMock := `{"API says":"Hey! How is it going felipe?"}`
	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, responseMock, string(responseBody))
}

func TestListStudentsHandler(t *testing.T) {
	database.DbConnect()
	createStudentMock()
	defer deleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students", controllers.ListStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")
}

func TestSearchStudentByCPFHandler(t *testing.T) {
	database.DbConnect()
	createStudentMock()
	defer deleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students/cpf/:cpf", controllers.RetrieveByCPF)
	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchStudentByIDHandler(t *testing.T) {
	database.DbConnect()
	createStudentMock()
	defer deleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students/:id", controllers.RetrieveStudent)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var mockStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &mockStudent)
	fmt.Println(mockStudent.Name)
	assert.Equal(t, http.StatusOK, response.Code, "Status Code should be OK")
	assert.Equal(t, "Test Student", mockStudent.Name, "Name should be equal")
	assert.Equal(t, "12345678901", mockStudent.CPF, "CPF should be equal")
	assert.Equal(t, "123456789", mockStudent.RG, "RG should be equal")
}

func TestDeleteStudentHandler(t *testing.T) {
	database.DbConnect()
	createStudentMock()

	r := SetupTestRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Status Code should be OK")
}

func TestUpdateStudentHandler(t *testing.T) {
	database.DbConnect()
	createStudentMock()
	defer deleteStudentMock()

	student := models.Student{Name: "Test Student", CPF: "54321678901", RG: "543216789"}
	studentJson, _ := json.Marshal(student)

	r := SetupTestRoutes()
	r.PATCH("/students/:id", controllers.UpdateStudent)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(studentJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Status Code should be OK")
	updatedStudent := models.Student{}
	json.Unmarshal(response.Body.Bytes(), &updatedStudent)
	assert.Equal(t, "Test Student", updatedStudent.Name, "Name shoulb be equal")
	assert.Equal(t, "54321678901", updatedStudent.CPF, "CPF should be equal")
	assert.Equal(t, "543216789", updatedStudent.RG, "RG shoulb be equal")
}
