# 某微信公众号自动回复API

## 配置文件

```ini
[app]
mode = 应用模式 # debug or release
password = 你的密码
jwtsecret = 你的密钥

[server]
port = 端口

[wechat]
id = 微信ID
appid = 开发者ID
appsecret = 开发者密钥
```

## 接口文档

1. 用户登录

   ```bash
   GET /login/<password>
   ```

   请求体：

   ```bash
   无
   ```

   响应：

   ```json
   {
       "code": 0,
       "data": "token字符串",
       "msg": "ok"
   }
   ```

2. 创建关键词

   ```bash
   POST /key
   ```

   请求体：

   ```json
   {
       "value": "关键词",
       "matchType": "all"
   }
   ```

   响应：

   ```json
   {
       "code": 0,
       "data": {
           "ID": 2,
           "CreatedTime": 1645529533178,
           "UpdatedTime": 1645529533178,
           "Value": "关键词",
           "MatchType": "all",
           "MessageID": 0
       },
       "msg": "ok"
   }
   ```

3. 删除关键词

   ```bash
   DELETE /key/<id>
   ```

   请求体：

   ```bash
   无
   ```

   响应：

   ```json
   {
       "code": 0,
       "data": null,
       "msg": "ok"
   }
   ```

4. 创建消息回复规则

   ```bash
   POST /msg
   ```

   请求体：

   ```json
   {
       "name": "测试",
       "msgType": "text",
       "content": "测试内容",
       "keywords": [
           {
               "value": "test",
               "matchType": "half"
           }
       ]
   }
   ```

   响应：

   ```json
   {
       "code": 0,
       "data": {
           "ID": 1,
           "CreatedTime": 1645529836427,
           "UpdatedTime": 1645529836427,
           "Name": "测试",
           "MsgType": "text",
           "Content": "测试内容",
           "MediaId": "",
           "Title": "",
           "Description": "",
           "MusicUrl": "",
           "HQMusicUrl": "",
           "ThumbMediaId": "",
           "PicUrl": "",
           "Url": "",
           "Keywords": [
               {
                   "ID": 3,
                   "CreatedTime": 1645529836428,
                   "UpdatedTime": 1645529836428,
                   "Value": "test",
                   "MatchType": "half",
                   "MessageID": 1
               }
           ]
       },
       "msg": "ok"
   }
   ```

5. 获取消息规则

   ```bash
   GET /msg/<id>
   ```

   请求体：

   ```bash
   无
   ```

   响应：

   ```json
   {
       "code": 1,
       "msg": "消息ID不存在"
   }
   ```

6. 获取全部消息

   ```bash
   GET /msgs/<pageNum>/<pageSize>
   ```

   请求体：

   ```bash
   无
   ```

   响应：

   ```json
   {
       "code": 0,
       "data": {
           "list": [
               {
                   "ID": 1,
                   "CreatedTime": 1645529836427,
                   "UpdatedTime": 1645529836427,
                   "Name": "测试",
                   "MsgType": "text",
                   "Content": "测试内容",
                   "MediaId": "",
                   "Title": "",
                   "Description": "",
                   "MusicUrl": "",
                   "HQMusicUrl": "",
                   "ThumbMediaId": "",
                   "PicUrl": "",
                   "Url": "",
                   "Keywords": null
               }
           ],
           "total": 1
       },
       "msg": "ok"
   }
   ```

7. 更新消息回复规则

   ```bash
   PUT /msg
   ```

   请求体：

   ```json
   {
       "ID": 1,
       "CreatedTime": 1645529836427,
       "UpdatedTime": 1645529836427,
       "Name": "测试",
       "MsgType": "text",
       "Content": "更改内容",
       "MediaId": "",
       "Title": "",
       "Description": "",
       "MusicUrl": "",
       "HQMusicUrl": "",
       "ThumbMediaId": "",
       "PicUrl": "",
       "Url": "",
       "Keywords": [
           {
               "ID": 3,
               "CreatedTime": 1645529836428,
               "UpdatedTime": 1645529836428,
               "Value": "test",
               "MatchType": "half",
               "MessageID": 1
           }
       ]
   }
   ```

   响应：

   ```json
   {
       "code": 0,
       "data": null,
       "msg": "ok"
   }
   ```

8. 删除消息回复规则

   ```bash
   DELETE /msg/<id>
   ```

   请求体：

   ```bash
   无
   ```

   响应：

   ```json
   {
       "code": 0,
       "data": null,
       "msg": "ok"
   }
   ```

9. 接收微信客户端消息

   ```bash
   POST /wechat
   ```

   请求体：

   ```xml
   <xml>
     <ToUserName><![CDATA[toUser]]></ToUserName>
     <FromUserName><![CDATA[fromUser]]></FromUserName>
     <CreateTime>1348831860</CreateTime>
     <MsgType><![CDATA[text]]></MsgType>
     <Content><![CDATA[this is a test]]></Content>
     <MsgId>1234567890123456</MsgId>
   </xml>
   ```

   响应：

   ```xml
   <xml>
     <ToUserName><![CDATA[toUser]]></ToUserName>
     <FromUserName><![CDATA[fromUser]]></FromUserName>
     <CreateTime>12345678</CreateTime>
     <MsgType><![CDATA[news]]></MsgType>
     <ArticleCount>1</ArticleCount>
     <Articles>
       <item>
         <Title><![CDATA[title1]]></Title>
         <Description><![CDATA[description1]]></Description>
         <PicUrl><![CDATA[picurl]]></PicUrl>
         <Url><![CDATA[url]]></Url>
       </item>
     </Articles>
   </xml>
   ```

10. 获取微信token

    ```bash
    GET /access
    ```

    请求体：

    ```bash
    无
    ```

    响应：

    ```json
    {
        "code": 0,
        "data": "token字符串",
        "msg": "ok"
    }
    ```

    
