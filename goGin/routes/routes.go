package routes

import (
	"github.com/PhilipFelipe/golang-alura-course/controllers"
	docs "github.com/PhilipFelipe/golang-alura-course/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequest() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/index", controllers.ShowIndexPage)
	// Students routes
	r.POST("/students", controllers.CreateStudent)
	// r.GET("/:name", controllers.Welcome)
	r.GET("/students", controllers.ListStudents)
	r.GET("/students/cpf/:cpf", controllers.RetrieveByCPF)
	r.GET("/students/:id", controllers.RetrieveStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 404
	r.NoRoute(controllers.NotFoundRoute)
	
	r.Run(":8000")
}
