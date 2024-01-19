package domain

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	UserId         string // 所属用户
	DirId          uint   // 所属目录
	FileName       string // 文件名
	FileSize       int64  // 文件大小
	MimeType       string // 媒体类型
	SavePath       string // 保存路径
	Password       string // 密码
	ExpireType     uint   // 过期类型：0=永久，1=时效性，2=次数
	ExpireValue    uint   // 过期值
	ExpireTime     uint64 // 到期时间
	Status         uint   // 状态：0=可下载，1=禁止下载
	DownloadNumber uint   // 下载次数
}
