package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/GoApi/api/controllers"
	"github.com/lucasvmiguel/GoApi/api/middlewares"
)

func Start(port string) {
	r := gin.Default()

	r.POST("/authorize", controllers.Authorize)

	authorized := r.Group("/")

	authorized.Use(middlewares.Authentication())
	{
		authorized.GET("/health", controllers.Health)
	}

	r.Run(port)
}
