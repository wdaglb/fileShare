package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type AppServer struct {
	handle *gin.Engine
}

var App *AppServer

func InitAppServer() {
	App = &AppServer{}
	loadConfig()
	initDb()
	App.handle = gin.New()
}

func RunAppServer() {
	if err := App.handle.Run(fmt.Sprintf(":%d", Config.Server.Port)); err != nil {
		log.Panicln(err)
	}
}
