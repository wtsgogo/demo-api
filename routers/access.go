package routers

import (
	"demo-api/config"
	"demo-api/wechat"

	"github.com/gin-gonic/gin"
)

func LoadAccess(g *gin.RouterGroup) {
	g.GET("/access", JWT(), getAccessTokenHandler)
}

func getAccessTokenHandler(c *gin.Context) {
	accessToken, err := wechat.GetAccessToken(config.WechatAppId, config.WechatAppSecret)
	if err != nil {
		// 这里没必要再打印多余的日志了
		failed(c, err.Error())
		return
	}
	success(c, accessToken)
}
