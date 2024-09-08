package model

import (
	"gorm.io/gorm"
	"time"
)

type Relation struct {
	RelationID      int        `gorm:"primaryKey;autoIncrement" json:"relation_id"`
	RelationUserID  int        `gorm:"not null" json:"relation_user_id"`
	RelationTypeID  int8       `gorm:"not null" json:"relation_type_id"`
	TransactionType int8       `gorm:"not null;default:1" json:"transaction_type"`
	Money           float64    `gorm:"not null" json:"money"`
	Date            string     `gorm:"not null" json:"date"`
	Remark          string     `gorm:"not null" json:"remark"`
	OwnerID         int        `gorm:"not null" json:"owner_id"`
	CreateTime      time.Time  `gorm:"not null" json:"create_time"`
	UpdateTime      *time.Time `json:"update_time"`
}

func (Relation) TableName() string {
	return "tb_relation"
}

func (r *Relation) BeforeCreate(tx *gorm.DB) (err error) {
	currentTime := time.Now()
	r.CreateTime = currentTime
	return
}

// BeforeUpdate 钩子会在更新记录之前调用
func (r *Relation) BeforeUpdate(tx *gorm.DB) (err error) {
	currentTime := time.Now()
	r.UpdateTime = &currentTime
	return
}
