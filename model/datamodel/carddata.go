package datamodel

type CardInfo struct {
	CardId          int    `gorm:"column:card_id;type:int(11);primary_key" json:"card_id"`
	CardName        string `gorm:"column:card_name;type:varchar(255);NOT NULL" json:"card_name"`
	CardLevel       int    `gorm:"column:card_level;type:int(11);NOT NULL" json:"card_level"`
	CardImageurl    string `gorm:"column:card_imageurl;type:varchar(255);NOT NULL" json:"card_imageurl"`
	CardDescription string `gorm:"column:card_description;type:varchar(255);NOT NULL" json:"card_description"`
}

func (m *CardInfo) TableName() string {
	return "card_info"
}

type CardAttribute struct {
	CardId          int `gorm:"column:card_id;type:int(11);primary_key" json:"card_id"`
	CardProbability int `gorm:"column:card_probability;type:int(11);NOT NULL" json:"card_probability"`
}

func (m *CardAttribute) TableName() string {
	return "card_attribute"
}
