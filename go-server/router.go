package main

import (
	"ke.file.share/api"
	"ke.file.share/core"
)

func initRouter() {
	core.App.GET("/file/list", api.GetFileList)
	core.App.GET("/file/get", api.GetFileInfo)
	core.App.POST("/upload", api.FileUpload)
	core.App.POST("/file/delete", api.DeleteFile)
	core.App.GET("/file/download", api.DownloadFile)

	core.App.GET("/dir/list", api.GetDirList)
	core.App.POST("/dir/create", api.CreateDir)
	core.App.POST("/dir/edit", api.EditDir)
	core.App.POST("/dir/delete", api.DeleteDir)

	core.App.GET("/config", api.GetConfig)
}
