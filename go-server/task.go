package main

import (
	"github.com/robfig/cron/v3"
	"ke.file.share/task"
)

func registerTask(cron *cron.Cron) {
	_, _ = cron.AddFunc("*/1 * * * *", task.AutoDeleteFile)
	cron.Run()
}
