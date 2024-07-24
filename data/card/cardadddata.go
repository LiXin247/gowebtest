package card

import (
	"gorm.io/gorm"
	"gowebtest/dbconnect"
	"gowebtest/model/datamodel"
	"gowebtest/tool/datacopy"
)

func CardAddData(name string, cardlevel int, cardimageurl string, description string) (err string) {
	db := dbconnect.DbInit()
	var alreadyexists datamodel.CardInfo
	result1 := db.DB.Where("card_name = ?", name).First(&alreadyexists)
	if result1.Error != gorm.ErrRecordNotFound {
		return "The card already exists"
	}
	var lastcard datamodel.CardInfo
	db.DB.Last(&lastcard)
	CardID := lastcard.CardId + 1
	cardimageurlcopy, err_1 := datacopy.CardCopy(CardID, cardimageurl)
	if err_1 != nil {
		return "Card image copy failed"
	}
	newCard := datamodel.CardInfo{CardID, name, cardlevel, cardimageurlcopy, description}
	result2 := db.DB.Create(&newCard)
	if result2.Error != nil {
		return "Card image added failed"
	}
	return "nil"

}
