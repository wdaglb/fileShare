package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/strutil"
	"ke.file.share/core"
	"ke.file.share/core/context"
	"ke.file.share/core/http"
	"ke.file.share/domain"
)

// 获取目录列表
func GetDirList(c *context.Context) {
	password, _ := c.GetQuery("password")
	var list []domain.Dir
	core.DB.Order("sort desc").Find(&list)
	list = arrutil.Filter(list, func(v domain.Dir) bool {
		if v.PublicType == domain.DirPublicType_Hide && v.Password != strutil.MD5(password) {
			return false
		}
		return true
	})
	list = arrutil.Map(list, func(input domain.Dir) (target domain.Dir, find bool) {
		input.Password = ""
		target = input
		find = true
		return
	})
	http.SUCCESS(c, gin.H{
		"list": list,
	})
}

// 创建目录
func CreateDir(c *context.Context) {
	var req struct {
		ParentId   uint
		Name       string
		Sort       uint
		PublicType uint
		Password   string
	}
	if domain.IsDirNameExist(req.ParentId, c.UserId, req.Name) {
		http.ERROR(c, "目录已存在")
		return
	}
	data := domain.Dir{
		UserId:     c.UserId,
		ParentId:   req.ParentId,
		Name:       req.Name,
		Sort:       req.Sort,
		PublicType: req.PublicType,
	}
	if req.Password != "" {
		data.Password = strutil.MD5(req.Password)
	}
	core.DB.Create(&data)
	http.SUCCESS(c, gin.H{
		"DirId": data.ID,
	})
}

// 修改目录
func EditDir(c *context.Context) {
	id := c.GetPkInt()
	if id <= 0 {
		http.ERROR(c, "参数错误")
		return
	}

	var req struct {
		ParentId   uint
		Name       string
		Sort       uint
		PublicType uint
		Password   string
	}
	if domain.IsDirNameExist(req.ParentId, c.UserId, req.Name) {
		http.ERROR(c, "目录已存在")
		return
	}
	var data domain.Dir
	core.DB.Where("id=? and user_id=?", id, c.UserId).First(&data)
	if data.ID == 0 {
		http.ERROR(c, "目录不存在")
		return
	}
	data.ParentId = req.ParentId
	data.Name = req.Name
	data.Sort = req.Sort
	data.PublicType = req.PublicType
	if req.Password != "" {
		data.Password = strutil.MD5(req.Password)
	}
	core.DB.Save(&data)
	http.SUCCESS(c)
}

// 删除目录
func DeleteDir(c *context.Context) {
	id := c.GetPkInt()
	if id <= 0 {
		http.ERROR(c, "参数错误")
		return
	}
	core.DB.Where("id=? and user_id=?", id, c.UserId).Delete(domain.Dir{})
	http.SUCCESS(c)
}
