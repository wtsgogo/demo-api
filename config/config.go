package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

var (
	AppMode         string
	AppPassword     string
	AppJwtSecret    string
	ServerPort      int
	WechatId        string
	WechatAppId     string
	WechatAppSecret string
)

func init() {
	// 设置日志输出文件
	logFile, err := os.OpenFile("demo-api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("无法打开日志文件:", err.Error())
		os.Exit(1)
	}
	log.SetOutput(logFile)
	// 导入配置信息
	cfg, err := ini.Load("demo-api.ini")
	if err != nil {
		log.Fatalln("无法加载配置文件:", err.Error())
	}
	AppMode = cfg.Section("app").Key("mode").MustString("release")
	AppPassword = cfg.Section("app").Key("password").MustString("123456")
	AppJwtSecret = cfg.Section("app").Key("jwtsecret").MustString("hard_to_guess_string")
	ServerPort = cfg.Section("server").Key("port").MustInt(8081)
	WechatId = cfg.Section("wechat").Key("id").String()
	WechatAppId = cfg.Section("wechat").Key("appid").String()
	WechatAppSecret = cfg.Section("wechat").Key("appsecret").String()
}
