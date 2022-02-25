package routers

import (
	"demo-api/auth"
	"demo-api/config"
	"log"

	"github.com/gin-gonic/gin"
)

func LoadLogin(g *gin.RouterGroup) {
	g.GET("/login/:password", loginHandler)
}

func loginHandler(c *gin.Context) {
	password := c.Param("password")
	if config.AppPassword != password {
		log.Println("无效密码:", password)
		failed(c, "密码错误")
		return
	}
	token, err := auth.GenerateToken(password)
	if err != nil {
		log.Println("生成token失败:", err.Error())
		failed(c, err.Error())
		return
	}
	success(c, token)
}
