package model

import (
	"gorm.io/gorm"
	"time"
)

type SysUser struct {
	UserID     uint       `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	Username   string     `gorm:"column:username;not null" json:"username"`
	Password   string     `gorm:"column:password;not null" json:"-"`
	Salt       string     `gorm:"column:salt;not null" json:"-"`
	Email      string     `gorm:"column:email;not null" json:"email"`
	Name       string     `gorm:"column:name;not null" json:"name"`
	Status     int8       `gorm:"column:status;not null;default:1" json:"status"`
	IsAdmin    int8       `gorm:"column:is_admin;default:1" json:"isAdmin"`
	CreateTime time.Time  `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime *time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName设置SysUser模型的表名。
func (SysUser) TableName() string {
	return "tb_sys_user"
}

func (r *SysUser) BeforeCreate(tx *gorm.DB) (err error) {
	currentTime := time.Now()
	r.CreateTime = currentTime
	return
}

// BeforeUpdate 钩子会在更新记录之前调用
func (r *SysUser) BeforeUpdate(tx *gorm.DB) (err error) {
	currentTime := time.Now()
	r.UpdateTime = &currentTime
	return
}
