package index_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skeleton/app/entities/models/user"
	"skeleton/app/entities/request"
	"skeleton/app/entities/response"
)

type IndexController struct {
}

func (c IndexController) Reg(ctx *gin.Context) {
	httpRequest := request.RegRequest{}

	err := ctx.ShouldBindJSON(&httpRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.New(err.Error(), nil))
		return
	}

	user4add, result := user.User{Username: httpRequest.Username, Password: httpRequest.Password}.Create()
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, response.New(result.Error.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, response.New("注册成功", user4add))
}
