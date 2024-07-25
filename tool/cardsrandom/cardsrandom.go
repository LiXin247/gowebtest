package cardsrandom

import (
	"fmt"
	"gowebtest/dbconnect"
	"gowebtest/model/datamodel"
)

func Cardsrandom(manner string) (cardid [10]int) {
	var cardRandomInfo []datamodel.CardAttribute
	db := dbconnect.DbInit()
	result := db.DB.Find(&cardRandomInfo)
	if result.Error != nil {
		fmt.Printf("Error retrieving card attributes: %v\n", result.Error)
		return
	}
	// 打印查询结果
	for _, card := range cardRandomInfo {
		fmt.Printf("Card ID: %d, Card Probability: %d\n", card.CardId, card.CardProbability)

	}
	return
}
