package models

import (
	"strings"
)

type Keyword struct {
	ID          uint   `gorm:"primaryKey"`
	CreatedTime int64  `gorm:"autoCreateTime:milli"`
	UpdatedTime int64  `gorm:"autoUpdateTime:milli"`
	Value       string `gorm:"unique;index"`
	MatchType   string
	MessageID   uint
}

func CreateKeyword(k *Keyword) error {
	result := db.Create(k)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteKeyword(id uint) {
	db.Delete(&Keyword{}, id)
}

func FindMsgByKey(key string) *Message {
	var k Keyword
	db.Where("value = ?", key).First(&k)
	if k.ID > 0 {
		return FindMsgById(k.MessageID)
	}
	var ks []Keyword
	db.Where("match_type = ?", "half").Order("id desc").Find(&ks)
	for _, k := range ks {
		if strings.Contains(key, k.Value) {
			return FindMsgById(k.MessageID)
		}
	}
	return nil
}
