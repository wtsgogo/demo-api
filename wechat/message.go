package wechat

import (
	"demo-api/config"
	"demo-api/models"
	"encoding/xml"
	"time"
)

// 基础消息类型
type BasicMessage struct {
	ToUserName   string `xml:"ToUserName" binding:"required"`   // 接收方帐号
	FromUserName string `xml:"FromUserName" binding:"required"` // 发送方账号
	CreateTime   int    `xml:"CreateTime" binding:"required"`   // 创建时间
	MsgType      string `xml:"MsgType" binding:"required"`      // 消息类型
}

// 接收消息类型
type ReceiveMessage struct {
	BasicMessage
	MsgId        int64   `xml:"MsgId"`        // 消息ID
	Content      string  `xml:"Content"`      // 文本内容
	PicUrl       string  `xml:"PicUrl"`       // 图片链接
	MediaId      string  `xml:"MediaId"`      // 消息媒体ID
	Format       string  `xml:"Format"`       // 语音格式
	Recognition  string  `xml:"Recognition"`  // 语音识别结果
	ThumbMediaId string  `xml:"ThumbMediaId"` // 缩略图媒体ID
	Location_X   float32 `xml:"Location_X"`   // 地理位置纬度
	Location_Y   float32 `xml:"Location_Y"`   // 地理位置经度
	Scale        int     `xml:"Scale"`        // 地图缩放大小
	Label        string  `xml:"Label"`        // 地理位置信息
	Url          string  `xml:"Url"`          // 链接地址
	Title        string  `xml:"Title"`        // 链接标题
	Description  string  `xml:"Description"`  // 链接描述
	Event        string  `xml:"Event"`        // 事件类型
	EventKey     string  `xml:"EventKey"`     // 事件KEY值
	Ticket       string  `xml:"Ticket"`       // 二维码的TICKET
	Latitude     float32 `xml:"Latitude"`     // 事件地理位置纬度
	Longitude    float32 `xml:"Longitude"`    // 事件地理位置经度
	Precision    float32 `xml:"Precision"`    // 地理位置精度
}

// 回复消息类型
type ReplyMessage struct {
	BasicMessage
	XMLName      xml.Name         `xml:"xml"`          // 根节点名
	Content      string           `xml:"Content"`      // 文本消息内容
	Image        ImageMessage     `xml:"Image"`        // 图片消息
	Voice        VoiceMessage     `xml:"Voice"`        // 语音消息
	Video        VideoMessage     `xml:"Video"`        // 视频消息
	Music        MusicMessage     `xml:"Music"`        // 音乐消息
	ArticleCount int              `xml:"ArticleCount"` // 图文消息个数
	Articles     ArticleContainer `xml:"Articles"`     // 图文消息
}

type ArticleContainer struct {
	XMLName xml.Name         `xml:"Articles"`
	Data    []ArticleMessage `xml:"item"`
}

// 图片消息
type ImageMessage struct {
	MediaId string `xml:"MediaId"` // 素材ID
}

// 语音消息
type VoiceMessage struct {
	MediaId string `xml:"MediaId"` // 素材ID
}

// 视频消息
type VideoMessage struct {
	MediaId     string `xml:"MediaId"`     // 素材ID
	Title       string `xml:"Title"`       // 标题
	Description string `xml:"Description"` // 描述
}

// 音乐消息
type MusicMessage struct {
	Title        string `xml:"Title"`        // 标题
	Description  string `xml:"Description"`  // 描述
	MusicUrl     string `xml:"MusicUrl"`     // 音乐链接
	HQMusicUrl   string `xml:"HQMusicUrl"`   // 高品质音乐链接
	ThumbMediaId string `xml:"ThumbMediaId"` // 缩略图素材ID
}

// 图文消息
type ArticleMessage struct {
	XMLName     xml.Name `xml:"item"`        // 根节点名
	Title       string   `xml:"Title"`       // 标题
	Description string   `xml:"Description"` // 描述
	PicUrl      string   `xml:"PicUrl"`      // 图片链接
	Url         string   `xml:"Url"`         // 图文消息链接
}

func Reply(toUserName string, m *models.Message) *ReplyMessage {
	bm := BasicMessage{
		ToUserName:   toUserName,
		FromUserName: config.WechatId,
		CreateTime:   int(time.Now().Unix()),
		MsgType:      m.MsgType,
	}
	switch m.MsgType {
	case "text":
		return &ReplyMessage{
			BasicMessage: bm,
			Content:      m.Content,
		}
	case "image":
		return &ReplyMessage{
			BasicMessage: bm,
			Image: ImageMessage{
				MediaId: m.MediaId,
			},
		}
	case "voice":
		return &ReplyMessage{
			BasicMessage: bm,
			Voice: VoiceMessage{
				MediaId: m.MediaId,
			},
		}
	case "video":
		return &ReplyMessage{
			BasicMessage: bm,
			Video: VideoMessage{
				MediaId:     m.MediaId,
				Title:       m.Title,
				Description: m.Description,
			},
		}
	case "music":
		return &ReplyMessage{
			BasicMessage: bm,
			Music: MusicMessage{
				Title:        m.Title,
				Description:  m.Description,
				MusicUrl:     m.MusicUrl,
				HQMusicUrl:   m.HQMusicUrl,
				ThumbMediaId: m.ThumbMediaId,
			},
		}
	case "news":
		return &ReplyMessage{
			BasicMessage: bm,
			ArticleCount: 1,
			Articles: ArticleContainer{
				Data: []ArticleMessage{
					{
						Title:       m.Title,
						Description: m.Description,
						PicUrl:      m.PicUrl,
						Url:         m.Url,
					},
				},
			},
		}
	default:
		return nil
	}
}
