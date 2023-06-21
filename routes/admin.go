package routes

import (
	"go-api-gateway/routes/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoute(rg *gin.RouterGroup) {
	admin := rg.Group("/back", middleware.CheckToken())
	admin.GET("/users", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		ctx.JSONP(http.StatusOK, data)
	})
}
