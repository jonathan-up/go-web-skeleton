package route

import (
	"github.com/gin-gonic/gin"
	"skeleton/app/http/controllers/index_controller"
)

func Load(router *gin.Engine) {

	indexController := index_controller.IndexController{}
	userGroup := router.Group("/user")
	{
		userGroup.POST("/reg", indexController.Reg)
	}

}
