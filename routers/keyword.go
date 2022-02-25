package routers

import (
	"demo-api/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoadKeyword(g *gin.RouterGroup) {
	g.POST("/key", JWT(), createKeyHandler)
	g.DELETE("/key/:id", JWT(), deleteKeyHandler)
}

func createKeyHandler(c *gin.Context) {
	var k models.Keyword
	err := c.ShouldBindJSON(&k)
	if err != nil {
		log.Println("解析请求数据失败:", err.Error())
		failed(c, err.Error())
		return
	}
	err = models.CreateKeyword(&k)
	if err != nil {
		log.Println("创建关键词失败:", err.Error())
		failed(c, err.Error())
		return
	}
	success(c, k)
}

func deleteKeyHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("参数ID错误:", err.Error())
		failed(c, err.Error())
		return
	}
	models.DeleteKeyword(uint(id))
	success(c, nil)
}
