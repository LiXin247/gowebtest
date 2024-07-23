package model

import (
	"database/sql"
)

type CardInfo struct {
	CardId          int            `gorm:"column:card_id;type:int(11);primary_key" json:"card_id"`
	CardName        sql.NullString `gorm:"column:card_name;type:varchar(255)" json:"card_name"`
	CardLevel       sql.NullInt32  `gorm:"column:card_level;type:int(11)" json:"card_level"`
	CardImageurl    sql.NullString `gorm:"column:card_imageurl;type:varchar(255)" json:"card_imageurl"`
	CardDescription sql.NullString `gorm:"column:card_description;type:varchar(255)" json:"card_description"`
}

func (m *CardInfo) TableName() string {
	return "card_info"
}
