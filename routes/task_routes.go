package routes

import (
	"github.com/abik1221/Kuraz_To_Do_Test/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine) {
	api := router.Group("/api/tasks")
	{
		api.GET("/", controllers.GetTasks)
		api.POST("/", controllers.CreateTask)
		api.PUT("/:id", controllers.MarkTaskCompleted)
		api.DELETE("/:id", controllers.DeleteTask)
	}
}
