package routes

import (
	"github.com/PhilipFelipe/golang-alura-course/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/index", controllers.ShowIndexPage)
	// Students routes
	r.POST("/students", controllers.CreateStudent)
	r.GET("/:name", controllers.Welcome)
	r.GET("/students", controllers.ListStudents)
	r.GET("/students/cpf/:cpf", controllers.RetrieveByCPF)
	r.GET("/students/:id", controllers.RetrieveStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.NoRoute(controllers.NotFoundRoute)
	
	r.Run(":8000")
}
