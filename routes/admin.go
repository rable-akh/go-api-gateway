package routes

import (
	"go-api-gateway/controllers/user"
	"go-api-gateway/routes/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoute(rg *gin.RouterGroup) {
	admin := rg.Group("/back", middleware.CheckToken())
	admin.GET("/users", user.UserList)
}
