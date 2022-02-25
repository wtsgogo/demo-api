package main

import (
	"demo-api/config"
	"demo-api/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.AppMode)
	r := gin.Default()
	r.Use(routers.Cors())
	v1 := r.Group("/api")
	routers.LoadLogin(v1)
	routers.LoadKeyword(v1)
	routers.LoadMessage(v1)
	routers.LoadWechat(v1)
	routers.LoadAccess(v1)
	r.Run(fmt.Sprintf("127.0.0.1:%d", config.ServerPort))
}
