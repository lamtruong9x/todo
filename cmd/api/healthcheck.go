package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) healthCheckHandler(c *gin.Context) {
	data := map[string]string{
		"status":      "available",
		"environment": app.cfg.env,
	}
	c.JSON(http.StatusOK, data)
}
