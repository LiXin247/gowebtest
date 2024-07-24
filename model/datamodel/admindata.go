package datamodel

import (
	"database/sql"
)

type AdminInfo struct {
	AdminId       int            `gorm:"column:admin_id;type:int(11);primary_key" json:"admin_id"`
	AdminName     sql.NullString `gorm:"column:admin_name;type:varchar(255)" json:"admin_name"`
	AdminPassword sql.NullString `gorm:"column:admin_password;type:varchar(255)" json:"admin_password"`
	LastloginTime sql.NullTime   `gorm:"column:lastlogin_time;type:datetime" json:"lastlogin_time"`
}

func (m *AdminInfo) TableName() string {
	return "admin_info"
}
