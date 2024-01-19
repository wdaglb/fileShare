package api

import (
	"github.com/gin-gonic/gin"
	"ke.file.share/core"
	"ke.file.share/core/context"
	"ke.file.share/core/http"
)

// 获取系统配置
func GetConfig(c *context.Context) {
	http.SUCCESS(c, gin.H{
		"config": gin.H{
			"part_domain":   core.Config.PartDomain,
			"public_domain": core.Config.PublicDomain,
		},
	})
}
