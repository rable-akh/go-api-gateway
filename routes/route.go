package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.New()

func RunRoutes() {

	/// Version 1 GROUP ///
	v1 := router.Group("/v1")
	defaultRoute(v1)
	authRoute(v1)
	///////////////////////

	router.Run(":5000")
}
