package jwt

import (
	"ginDemo/pkg/e"
	"ginDemo/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
*
如何编写一个gin的中间件，其实就类似于拦截器？
*/
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 也就是去验证是否包含有token
		var data interface{}
		token := c.Query("token")
		var code int = e.SUCCESS
		// 默认获取到的为空吗
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			// 如果有token
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 为什么放在这里来检查
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
