package datamodel

type UserInfo struct {
	UserId       int    `gorm:"column:user_id;type:int(11);primary_key" json:"user_id"`
	UserName     string `gorm:"column:user_name;type:varchar(255);NOT NULL" json:"user_name"`
	UserPassword string `gorm:"column:user_password;type:varchar(255);NOT NULL" json:"user_password"`
	UserLevel    int    `gorm:"column:user_level;type:int(11);NOT NULL" json:"user_level"`
	UserBalance  int    `gorm:"column:user_balance;type:int(11);NOT NULL" json:"user_balance"`
}

func (m *UserInfo) TableName() string {
	return "user_info"
}

type UserCardLibrary struct {
	UserId  int `gorm:"column:user_id;type:int(11);primary_key" json:"user_id"`
	CardId  int `gorm:"column:card_id;type:int(11);NOT NULL" json:"card_id"`
	CardNum int `gorm:"column:card_num;type:int(11);NOT NULL" json:"card_num"`
}

func (m *UserCardLibrary) TableName() string {
	return "user_card_library"
}
