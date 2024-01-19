package main

import (
	"github.com/robfig/cron/v3"
	"ke.file.share/core"
	"ke.file.share/core/context"
	"ke.file.share/core/http"
	"ke.file.share/domain"
	"log"
)

func migrate() {
	err := core.DB.AutoMigrate(&domain.Dir{}, &domain.File{})
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	core.InitAppServer()
	core.App.Use(func(context *context.Context) {
		userId := context.GetHeader("X-UserId")
		if userId == "" {
			http.ERROR(context, "X-UserId为空")
			context.Abort()
		} else {
			context.Set("__userId", userId)
			context.Next()
		}
	})
	migrate()
	initRouter()
	go registerTask(cron.New())
	core.RunAppServer()
}
