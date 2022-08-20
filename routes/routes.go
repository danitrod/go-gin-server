package routes

import (
	"github.com/danitrod/go-gin-server/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", controllers.DisplayIndexPage)
	r.GET("/students", controllers.GetStudents)
	r.GET("/students/:id", controllers.GetStudent)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	r.NoRoute(controllers.NotFound)
	r.Run()
}
