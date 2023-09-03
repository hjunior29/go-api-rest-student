package routes

import (
	"Api-Rest-Gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/students", controllers.AllStudents)
	r.GET("/:name", controllers.Hello)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.SearchStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchCpfStudent)
	r.Run()
}
