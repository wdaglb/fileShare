package task

import (
	"github.com/gookit/goutil/fsutil"
	"ke.file.share/core"
	"ke.file.share/domain"
	"log"
	"os"
	"time"
)

func AutoDeleteFile() {
	nowTime := time.Now().Unix()
	var list []domain.File
	core.DB.Where("expire_time>0 AND expire_time<=?", nowTime).Limit(100).Order("created_at asc").Find(&list)
	for _, fileData := range list {
		// 判断文件是否存在
		if fsutil.FileExists(fileData.SavePath) {
			_ = os.Remove(fileData.SavePath)
			log.Printf("删除过期文件ID[%d] 路径[%s]\n", fileData.ID, fileData.SavePath)
		}
		core.DB.Where("id=?", fileData.ID).Delete(domain.File{})
	}
}
