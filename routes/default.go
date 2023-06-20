package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func defaultRoute(rg *gin.RouterGroup) {
	def := rg.Group("/default")
	def.GET("/", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		ctx.JSONP(http.StatusOK, data)
	})
}
