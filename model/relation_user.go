package model

import (
	"gorm.io/gorm"
	"time"
)

type RelationUser struct {
	RelationUserID   int        `gorm:"primaryKey;autoIncrement" json:"relation_user_id"`
	RelationUserName string     `gorm:"not null" json:"relation_user_name"`
	Prefix           string     `gorm:"not null" json:"prefix"`
	Sex              int8       `gorm:"not null" json:"sex"`
	Status           int8       `gorm:"not null" json:"status"`
	OwnerID          int        `gorm:"not null" json:"owner_id"`
	Remark           string     `gorm:"not null" json:"remark"`
	CreateTime       time.Time  `gorm:"not null" json:"create_time"`
	UpdateTime       *time.Time `json:"update_time"`
}

func (RelationUser) TableName() string {
	return "tb_relation_user"
}

func (r *RelationUser) BeforeCreate(tx *gorm.DB) (err error) {
	currentTime := time.Now()
	r.CreateTime = currentTime
	return
}

// BeforeUpdate 钩子会在更新记录之前调用
func (r *RelationUser) BeforeUpdate(tx *gorm.DB) (err error) {
	currentTime := time.Now()
	r.UpdateTime = &currentTime
	return
}
