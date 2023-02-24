package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skeleton/app/http/middlewares/logger"
	zapLogger "skeleton/bootstrap/logger"
	"skeleton/config"
	"skeleton/route"
)

var Engine *gin.Engine

func Init() {
	gin.SetMode(gin.ReleaseMode)
	Engine = gin.New()
	Engine.Use(logger.Middleware(), gin.Recovery())
	Engine.StaticFS("public/", http.Dir("public"))
	route.Load(Engine)
}

func Run() {
	err := Engine.Run(config.YAML.Server.Addr)
	if err != nil {
		zapLogger.SugarLogger.Errorf("HTTP服务启动失败 -> %v\n", err)
		return
	}
}
