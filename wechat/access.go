package wechat

import (
	"demo-api/cache"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AccessToken struct {
	Access_token string `json:"access_token"` // 令牌
	Expires_in   int    `json:"expires_in"`   // 过期时间
}

func GetAccessToken(appID, appSecret string) (string, error) {
	cacheAccessToken, _ := cache.GetValue("accessToken")
	if cacheAccessToken != "" {
		return cacheAccessToken, nil
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appID, appSecret)
	res, err := http.Get(url)
	if err != nil {
		log.Println("请求获取AccessToken失败:", err.Error())
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("无法读取AccessToken:", err.Error())
		return "", err
	}
	var ac AccessToken
	err = json.Unmarshal(body, &ac)
	if err != nil {
		log.Println("JSON转换失败,请检查错误码和IP白名单设置:", err.Error())
		return "", err
	}
	cache.SetValue("accessToken", ac.Access_token, ac.Expires_in-(ac.Expires_in/5))
	return ac.Access_token, nil
}
