package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ke.file.share/core/context"
	"net/http"
)

func SUCCESS(c *context.Context, data ...any) {
	if len(data) > 0 {
		c.JSON(200, data[0])
	} else {
		c.Status(200)
	}
}

func ERR(c *context.Context, err error) {
	c.JSON(http.StatusExpectationFailed, gin.H{
		"code":  "fail",
		"error": err.Error(),
	})
}

func ERROR(c *context.Context, content string, args ...any) {
	c.JSON(http.StatusExpectationFailed, gin.H{
		"code":  "fail",
		"error": fmt.Sprintf(content, args...),
	})
}

func ERRORF(c *context.Context, code, content string, args ...any) {
	c.JSON(http.StatusExpectationFailed, gin.H{
		"code":  code,
		"error": fmt.Sprintf(content, args...),
	})
}

func CODE(c *context.Context, code int, content string, args ...any) {
	c.JSON(code, gin.H{
		"code":  "fail",
		"error": fmt.Sprintf(content, args...),
	})
}
