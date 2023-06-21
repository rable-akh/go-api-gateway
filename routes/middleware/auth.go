package middleware

import (
	pb "go-api-gateway/protos/go-api-gateway/auth"
	"log"
	"net/http"
	"strings"

	auth "go-api-gateway/domains/auth"

	"github.com/gin-gonic/gin"
)

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req pb.CheckTokenRequest
		headerToken := ctx.Request.Header.Get("Authorization")
		req.Token = strings.Replace(headerToken, "Bearer ", "", -1)

		res, err := auth.CheckToken(&req)

		if err != nil {
			log.Println(err)
		}
		if !res.Status {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			ctx.Abort()
		}

		ctx.Next()
	}
}
