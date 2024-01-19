package context

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
)

type Context struct {
	*gin.Context
	UserId string
}

func (t *Context) GetPkInt() int64 {
	str := t.Query("id")
	val, _ := strutil.ToInt64(str)
	return val
}

func (t *Context) GetPagination() (int, int) {
	current, cExist := t.GetQuery("current")
	pageSize, pExist := t.GetQuery("page_size")
	var currentInt int
	var pageSizeInt int
	if cExist {
		currentInt, _ = strutil.ToInt(current)
		if currentInt < 1 {
			currentInt = 1
		}
	} else {
		currentInt = 1
	}
	if pExist {
		pageSizeInt, _ = strutil.ToInt(pageSize)
		if pageSizeInt < 1 {
			pageSizeInt = 1
		} else if pageSizeInt > 1000 {
			pageSizeInt = 1000
		}
	} else {
		pageSizeInt = 15
	}
	return (currentInt - 1) * pageSizeInt, pageSizeInt
}
