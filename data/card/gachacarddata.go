package card

import (
	"fmt"
	"gorm.io/gorm"
	"gowebtest/dbconnect"
	"gowebtest/model/datamodel"
	"gowebtest/tool/cardsrandom"
)

func GachaCardData(frequency string, userid int) error {
	db := dbconnect.DbInit()
	var (
		result   *gorm.DB
		usercard datamodel.UserCardLibrary
	)
	cardid, err := cardsrandom.Cardsrandom(frequency)
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		if cardid[i] != 0 {
			result = db.DB.Where("user_id = ? AND card_id = ?  ", userid, cardid[i]).First(&usercard)
			if result.Error != nil {
				if result.Error == gorm.ErrRecordNotFound {
					usercard = datamodel.UserCardLibrary{userid, int(cardid[i]), 1}
					db.DB.Create(&usercard)
				} else {
					fmt.Println("Error querying user card info:", result.Error)
				}
			} else {
				db.DB.Where("user_id = ? AND card_id = ?  ", userid, cardid[i]).Model(&usercard).Update("card_num", gorm.Expr("card_num + ?", 1))
			}

		} else {
			break
		}
	}
	return err
}
