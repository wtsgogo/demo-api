package models

type Message struct {
	ID           uint  `gorm:"primaryKey"`
	CreatedTime  int64 `gorm:"autoCreateTime:milli"`
	UpdatedTime  int64 `gorm:"autoUpdateTime:milli"`
	Name         string
	MsgType      string
	Content      string
	MediaId      string
	Title        string
	Description  string
	MusicUrl     string
	HQMusicUrl   string
	ThumbMediaId string
	PicUrl       string
	Url          string
	Keywords     []Keyword
}

func CreateMessage(m *Message) error {
	result := db.Create(m)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindMsgById(id uint) *Message {
	var m Message
	result := db.Preload("Keywords").First(&m, id)
	if result.Error != nil {
		return nil
	}
	return &m
}

func FindMsgAll(pageNum, pageSize int) ([]Message, int64) {
	var ms []Message
	db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&ms)
	var total int64
	db.Model(&Message{}).Count(&total)
	return ms, total
}

func UpdateMessage(m *Message) error {
	result := db.Save(m)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteMessage(id uint) {
	db.Where("message_id = ?", id).Delete(&Keyword{})
	db.Delete(&Message{}, id)
}
