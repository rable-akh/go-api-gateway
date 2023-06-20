package routes

import (
	"go-api-gateway/controllers/auth"

	"github.com/gin-gonic/gin"
)

func authRoute(rg *gin.RouterGroup) {

	def := rg.Group("/auth")
	def.POST("/login", auth.Login)
	def.POST("/signup", auth.Login)
}
