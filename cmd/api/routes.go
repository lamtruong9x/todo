package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	r := gin.Default()
	r.GET("/healthcheck", app.healthCheckHandler)
	// tasks
	r.POST("tasks/", app.createTaskHandler)
	r.DELETE("tasks/:id", app.deleteTaskHandler)
	r.GET("tasks/:id", app.getTaskHandler)
	r.PATCH("tasks/:id", app.updateTaskHandler)
	return r
}
