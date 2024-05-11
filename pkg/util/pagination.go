package util

// 分页参数
import (
	"ginDemo/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 计算分页的页数
func GetPage(c *gin.Context) int {
	result := 0
	// 很像是一种工具类，做类型转换？
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
