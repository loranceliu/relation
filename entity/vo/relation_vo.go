package vo

import "time"

type RelationVO struct {
	RelationID       int       `json:"relationId"`
	Money            float64   `json:"money"`
	RelationTypeName string    `json:"relationTypeName"`
	RelationTypeId   int       `json:"relationTypeId"`
	RelationUserName string    `json:"relationUserName"`
	RelationUserId   int       `json:"relationUserId"`
	TransactionType  int       `json:"transactionType"`
	Date             time.Time `json:"date"`
	Remark           string    `json:"remark"`
}

type RelationUserVO struct {
	RelationUserID   int     `json:"relationUserId"`
	RelationUserName string  `json:"relationUserName"`
	Sex              int8    `json:"sex"`
	Status           int8    `json:"status"`
	Prefix           string  `json:"prefix"`
	Income           float64 `json:"income"`
	Expend           float64 `json:"expend"`
	Remark           string  `json:"remark"`
}

type RelationUserItemVO struct {
	Type string           `json:"type"`
	Data []RelationUserVO `json:"data"`
}

type RelationUserIndexVO struct {
	Index []string             `json:"index"`
	Item  []RelationUserItemVO `json:"item"`
}

type SystemUserVO struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Status   int8   `json:"status"`
}

type RelationTypeVO struct {
	RelationTypeID   int    `json:"relationTypeId"`
	RelationTypeName string `json:"relationTypeName"`
}
