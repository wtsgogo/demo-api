package routers

import (
	"demo-api/cache"
	"demo-api/models"
	"demo-api/wechat"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoadWechat(g *gin.RouterGroup) {
	g.POST("/wechat", wechatHandler)
}

func wechatHandler(c *gin.Context) {
	var rm wechat.ReceiveMessage
	// Postman格式化的xml需要在这里去掉字符串前后空格
	err := c.ShouldBindXML(&rm)
	if err != nil {
		log.Println("无法解析该用户消息:", err.Error())
		c.String(http.StatusOK, "success")
		return
	}
	// 缓存已回复的消息
	var msgId, rdbKey string
	if rm.MsgType == "event" {
		rdbKey = rm.FromUserName + strconv.Itoa(rm.CreateTime)
	} else {
		rdbKey = strconv.Itoa(int(rm.MsgId))
	}
	msgId, _ = cache.GetValue(rdbKey)
	if msgId == "" {
		cache.SetValue(rdbKey, "replied", 30)
	} else {
		c.String(http.StatusOK, "success")
		return
	}
	// 处理消息
	switch rm.MsgType {
	case "event":
		eventMsgHandler(c, &rm)
	case "text":
		textMsgHandler(c, &rm)
	default:
		c.String(http.StatusOK, "success")
	}
}

// 事件消息处理
func eventMsgHandler(c *gin.Context, rm *wechat.ReceiveMessage) {
	//dosomething()
}

// 文本消息处理
func textMsgHandler(c *gin.Context, rm *wechat.ReceiveMessage) {
	keyword := strings.TrimSpace(rm.Content)
	m := models.FindMsgByKey(keyword)
	if m == nil {
		log.Println("无法回复关键词:", keyword)
		c.String(http.StatusOK, "success")
		return
	}
	c.XML(http.StatusOK, wechat.Reply(rm.FromUserName, m))
}
