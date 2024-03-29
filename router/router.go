package router

import (
	"github.com/gin-gonic/gin"
)

func Inicialize() {
	// Initialize router
	router := gin.Default()
	// Initialize Routes
	InitializeRoutes(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
