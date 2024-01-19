package domain

import (
	"gorm.io/gorm"
	"ke.file.share/core"
)

const (
	DirPublicType_Public   = iota // 公开
	DirPublicType_Password        // 密码访问
	DirPublicType_Hide            // 隐藏
)

type Dir struct {
	gorm.Model
	UserId     string // 所属用户
	ParentId   uint   // 上级目录
	Name       string // 目录名称
	Sort       uint   // 排序
	PublicType uint   // 权限：0=公开，1=密码访问，2=隐藏
	Password   string // 密码
}

// 目录是否存在
func IsDirNameExist(parentId uint, userId, name string) bool {
	var data Dir
	core.DB.Where("parent_id=? and user_id=? and name=?", parentId, userId, name).Select("id").First(&data)
	return data.ID > 0
}
