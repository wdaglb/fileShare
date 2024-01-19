package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
	"io"
	"ke.file.share/core"
	"ke.file.share/core/context"
	"ke.file.share/core/http"
	"ke.file.share/domain"
	"log"
	"os"
	"time"
)

// 获取文件列表
func GetFileList(c *context.Context) {
	var (
		list  []domain.File
		total int64
	)
	offset, limit := c.GetPagination()

	core.DB.Where("user_id=?", c.UserId).Offset(offset).Limit(limit).Order("created_at desc").Find(&list)
	core.DB.Where("user_id=?", c.UserId).Count(&total)
	http.SUCCESS(c, gin.H{
		"list":  list,
		"total": total,
	})
}

// 获取文件信息
func GetFileInfo(c *context.Context) {
	var req struct {
		Code string `form:"code" binding:"required"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		http.ERR(c, err)
		return
	}
	var file domain.File
	core.DB.Where("password=?", req.Code).First(&file)
	if file.ID == 0 {
		http.ERROR(c, "文件不存在")
		return
	}
	http.SUCCESS(c, gin.H{
		"file": file,
	})
}

// 文件上传
func FileUpload(c *context.Context) {
	var req struct {
		ExpireType  uint
		ExpireValue uint
	}
	if err := c.ShouldBind(&req); err != nil {
		http.ERR(c, err)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		http.ERR(c, err)
		return
	}
	fs, err := file.Open()
	if err != nil {
		http.ERR(c, err)
		return
	}
	defer fs.Close()

	rootPath := core.Config.Storage.FilePath
	hashName := strutil.Md5(fmt.Sprintf("%d%s%s", time.Now().Unix(), strutil.RandomChars(8), file.Filename))
	dirName := hashName[0:2] + "/"
	_ = os.Mkdir(rootPath+dirName, 0755)
	savePath := rootPath + dirName + hashName + "__" + file.Filename

	code := strutil.RandWithTpl(6, "01234567890")

	out, err := os.Create(savePath)
	if err != nil {
		http.ERR(c, err)
		return
	}
	_, _ = io.Copy(out, fs)
	_ = out.Close()

	// 将文件存入数据库
	fileData := domain.File{
		UserId:      c.UserId,
		FileName:    file.Filename,
		FileSize:    file.Size,
		MimeType:    file.Header.Get("content-type"),
		SavePath:    savePath,
		ExpireType:  req.ExpireType,
		ExpireValue: req.ExpireValue,
		ExpireTime:  0,
		Status:      0,
		Password:    code,
	}
	if req.ExpireType == 1 {
		fileData.ExpireTime = uint64(time.Now().Unix() + 86400*int64(req.ExpireValue))
	}
	if err := core.DB.Create(&fileData).Error; err != nil {
		http.ERR(c, err)
		return
	}

	http.SUCCESS(c, gin.H{
		"code": code,
	})
}

// 删除文件
func DeleteFile(c *context.Context) {
	var req struct {
		Code string `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		http.ERR(c, err)
		return
	}
	var file domain.File
	core.DB.Where("password=?", req.Code).First(&file)
	if file.ID == 0 {
		http.ERROR(c, "文件不存在")
		return
	}
	// 删除文件记录
	if err := core.DB.Where("id=?", file.ID).Delete(domain.File{}).Error; err != nil {
		http.ERR(c, err)
		return
	}
	// 文件删除
	if fsutil.FileExists(file.SavePath) {
		err := fsutil.Remove(file.SavePath)
		if err != nil {
			log.Printf("文件[%s]删除失败\n", file.SavePath)
		}
	}

	http.SUCCESS(c)
}

// 下载文件
func DownloadFile(c *context.Context) {
	code := c.Query("code")
	if code == "" {
		http.ERROR(c, "code参数为空")
		return
	}
	var data domain.File
	core.DB.Where("password=? and status=0", code).First(&data)
	if data.ID == 0 {
		http.ERROR(c, "文件不存在")
		return
	}
	updateData := domain.File{}
	updateData.DownloadNumber++
	if data.ExpireType == 2 {
		if data.DownloadNumber >= data.ExpireValue {
			updateData.Status = 1
			// 设置1小时后过期删除，防止当前下载异常
			updateData.ExpireTime = uint64(time.Now().Unix() + 3600)
		}
	}
	core.DB.Where("id=?", data.ID).Updates(updateData)
	c.File(data.SavePath)
}
