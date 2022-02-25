package routers

import (
	"demo-api/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}

func failed(c *gin.Context, msg string) {
	// 只要是应用返回的,http状态都设为200
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  msg,
	})
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")
		if token != "" {
			claims, err := auth.ParseToken(token)
			if err == nil && time.Now().Unix() <= claims.ExpiresAt {
				c.Next()
				return
			}
		}
		// code = 2 表示需要授权,等价于http状态码401
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "未成功获取授权信息",
		})
		c.Abort()
	}
}

// 跨域访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "7200")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
