package main

import (
	"os"
	_ "insur-box/src/db"
	"insur-box/src/router"
	"github.com/gin-gonic/gin"
	"io"
	"insur-box/src/config"
)

func main() {
	app := gin.New()
	f, _ := os.Create("./logs/api.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.SetMode(gin.ReleaseMode)
	app.Use(gin.Logger(), gin.Recovery())
	router.Route(app)
	app.Run(config.ServerPort)

}
