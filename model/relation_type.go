package model

type RelationType struct {
	RelationTypeID   int    `gorm:"primaryKey;autoIncrement" json:"relation_type_id"`
	RelationTypeName string `gorm:"not null" json:"relation_type_name"`
	OwnerID          int    `gorm:"not null" json:"owner_id"`
}

func (RelationType) TableName() string {
	return "tb_relation_type"
}
