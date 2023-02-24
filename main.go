package main

import (
	"skeleton/bootstrap/config/yaml"
	"skeleton/bootstrap/database"
	"skeleton/bootstrap/http"
	"skeleton/bootstrap/logger"
	"skeleton/config"
)

func main() {

	yaml.Init()
	logger.Init()
	database.Init()

	logger.SugarLogger.Infof("HTTP Server Will Run in -> %s", config.YAML.Server.Addr)
	http.Init()
	http.Run()
}
