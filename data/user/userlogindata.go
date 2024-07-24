package user

import (
	"gowebtest/dbconnect"
	"gowebtest/model/datamodel"
)

func UserLoginData(name string, password string) error {
	db := dbconnect.DbInit()
	var user datamodel.UserInfo
	result := db.DB.Where("user_name = ? AND user_password = ?", name, password).First(&user)
	return result.Error
}
