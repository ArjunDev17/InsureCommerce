// cmd/myapp/main.go
package main

import (
	"insurecommerce/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/register", handlers.RegisterStudent)
		api.POST("/login", handlers.LoginStudent)
	}

	router.Run(":8089")
}
