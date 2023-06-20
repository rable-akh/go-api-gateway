package auth

import (
	"encoding/json"
	"fmt"
	"go-api-gateway/config"
	"go-api-gateway/domains"
	auth "go-api-gateway/domains/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"pass" json:"pass" xml:"pass" binding:"required"`
}

type SignUpForm struct {
	Name     string `form:"name" json:"name" xml:"name" binding:"required"`
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
	Phone    string `form:"phone" json:"phone" xml:"phone" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type Response struct {
	Status  int16        `json:"status,omitempty"`
	Message string       `json:"message,omitempty"`
	Results json.Decoder `json:"results,omitempty"`
}

func Login(ctx *gin.Context) {
	// Call Service
	conn, err := domains.ServiceConn(config.AuthServiceURI())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	//////// connect with client //////////
	client := auth.AuthServiceClient(conn)
	//////// connect with client //////////
	var form LoginForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "",
			"results": gin.H{},
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	// if ctx.ShouldBindJSON(&form) != nil {
	data, _ := json.Marshal(form)
	fmt.Println(data)
	loginReq, _ := auth.CreateLoginRequest(data)

	fmt.Println("REQ", loginReq)
	results, err := auth.Login(client, loginReq)
	fmt.Println(results)
	if err != nil {
		fmt.Println("ERRRRR", err)
	}

	ctx.JSON(http.StatusOK, results)

}

func SignUp(ctx *gin.Context) {
	conn, err := domains.ServiceConn(config.AuthServiceURI())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	//////// connect with client //////////
	client := auth.AuthServiceClient(conn)
	//////// connect with client //////////

	var form SignUpForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "",
			"results": gin.H{},
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	data, _ := json.Marshal(form)
	signReq, _ := auth.CreateSignUpRequest(data)
	results, err := auth.SignUp(client, signReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "",
			"results": gin.H{},
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, results)
}
