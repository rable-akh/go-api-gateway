package user

import (
	"encoding/json"
	"fmt"
	"go-api-gateway/domains/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Pages   string `form:"page"`
	PerPage string `form:"perPage"`
}

func UserList(ctx *gin.Context) {
	var params Params

	if err := ctx.ShouldBindQuery(&params); err != nil {
		fmt.Println("Error: ", err)
	}

	reqD, _ := json.Marshal(params)
	// fmt.Println(params)
	// fmt.Println(reqD)
	uReq, _ := auth.UsersRequest(reqD)
	results, err := auth.Users(uReq)
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
