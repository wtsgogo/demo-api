package routers

import (
	"demo-api/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoadMessage(g *gin.RouterGroup) {
	g.POST("/msg", JWT(), createMsgHandler)
	g.GET("/msg/:id", getMsgByIdHandler)
	g.GET("/msgs/:pageNum/:pageSize", getMsgAllHandler)
	g.PUT("/msg", JWT(), updateMsgHandler)
	g.DELETE("/msg/:id", JWT(), deleteMsgHandler)
}

func createMsgHandler(c *gin.Context) {
	var m models.Message
	err := c.ShouldBindJSON(&m)
	if err != nil {
		log.Println("解析请求数据失败:", err.Error())
		failed(c, err.Error())
		return
	}
	err = models.CreateMessage(&m)
	if err != nil {
		log.Println("创建消息回复规则失败:", err.Error())
		failed(c, err.Error())
		return
	}
	success(c, m)
}

func getMsgByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("参数ID错误:", err.Error())
		failed(c, err.Error())
		return
	}
	m := models.FindMsgById(uint(id))
	if m == nil {
		log.Println("不存在此消息ID:", id)
		failed(c, "消息ID不存在")
		return
	}
	success(c, m)
}

func getMsgAllHandler(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Param("pageNum"))
	if err != nil {
		log.Println("分页参数错误:", err.Error())
		failed(c, err.Error())
		return
	}
	pageSize, err := strconv.Atoi(c.Param("pageSize"))
	if err != nil {
		log.Println("分页参数错误:", err.Error())
		failed(c, err.Error())
		return
	}
	if pageNum < 1 || pageSize < 1 {
		log.Println("分页参数不能小于1")
		failed(c, "分页参数不能小于1")
		return
	}
	ms, total := models.FindMsgAll(pageNum, pageSize)
	success(c, gin.H{
		"list":  ms,
		"total": total,
	})
}

func updateMsgHandler(c *gin.Context) {
	var m models.Message
	err := c.ShouldBindJSON(&m)
	if err != nil {
		log.Println("解析JSON数据失败:", err.Error())
		failed(c, err.Error())
		return
	}
	err = models.UpdateMessage(&m)
	if err != nil {
		log.Println("更新消息回复规则失败:", err.Error())
		failed(c, err.Error())
		return
	}
	success(c, nil)
}

func deleteMsgHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("参数ID错误:", err.Error())
		failed(c, err.Error())
		return
	}
	models.DeleteMessage(uint(id))
	success(c, nil)
}
