package user

import (
	"gowebtest/dbconnect"
	"gowebtest/model/datamodel"
)

func UserRegisterData(name string, password string) (err string) {
	db := dbconnect.DbInit()
	var lastuser datamodel.UserInfo
	db.DB.Last(&lastuser)
	UID := lastuser.UserId + 1
	var alreadyexists datamodel.UserInfo
	result_1 := db.DB.Where("user_name = ?", name).First(&alreadyexists)
	if result_1.Error == nil {
		return "The user already exists"
	}
	newUser := datamodel.UserInfo{UID, name, password, 0, 0}
	result_2 := db.DB.Create(&newUser)
	if result_2.Error != nil {
		return "Could not create user"
	}

	return "nil"
}
