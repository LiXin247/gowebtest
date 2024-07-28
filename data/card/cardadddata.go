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
	switch newCard.CardLevel {
	case 3:
		{
			var lastCard3 datamodel.CardInfo3
			db.DB.Last(&lastCard3)
			CardAddr := lastCard3.CardAddr + 1
			newCard3 := datamodel.CardInfo3{CardID, CardAddr}
			result23 := db.DB.Create(&newCard3)
			if result23.Error != nil {
				return "Card image added failed"
			}
			return "nil"
		}
	case 4:
		{
			var lastCard4 datamodel.CardInfo4
			db.DB.Last(&lastCard4)
			CardAddr := lastCard4.CardAddr + 1
			newCard4 := datamodel.CardInfo4{CardID, CardAddr}
			result24 := db.DB.Create(&newCard4)
			if result24.Error != nil {
				return "Card image added failed"
			}
			return "nil"
		}
	case 5:
		{
			var lastCard5 datamodel.CardInfo5
			db.DB.Last(&lastCard5)
			CardAddr := lastCard5.CardAddr + 1
			newCard5 := datamodel.CardInfo5{CardID, CardAddr}
			result25 := db.DB.Create(&newCard5)
			if result25.Error != nil {
				return "Card image added failed"
			}
			return "nil"
		}
	}
	if result2.Error != nil {
		return "Card image added failed"
	}
	return "nil"

}
