package cardsrandom

import (
	"fmt"
	"gowebtest/dbconnect"
	"gowebtest/model/datamodel"
	"log"
	"math/rand"
)

func Cardsrandom(frequency string) (cardid [10]int64, err error) {
	var cardRandomInfo []datamodel.CardAttribute
	var cardinfo_3 datamodel.CardInfo3
	var cardinfo_4 datamodel.CardInfo4
	var cardinfo_5 datamodel.CardInfo5
	db := dbconnect.DbInit()
	result := db.DB.Find(&cardRandomInfo)
	if result.Error != nil {
		fmt.Printf("Error retrieving card attributes: %v\n", result.Error)
		panic("Tool run failed")
	}
	// 打印查询结果
	for _, card := range cardRandomInfo {
		fmt.Printf("Card ID: %d, Card Probability: %d\n", card.CardId, card.CardProbability)

	}
	var randomnum [10]int
	randomnum = RandomNum(frequency)  //n
	cardlevel := Getalevel(randomnum) //random1(n)=a  limit a
	for i := 0; i < len(randomnum); i++ {
		if cardlevel[i] != 0 {
			if cardlevel[i] == 3 {
				result = db.DB.Where("card_addr = ?", AddrRandomNum(cardlevel[i])).Take(&cardinfo_3)
				if result.Error != nil {
					fmt.Printf("read failed")
				}
				cardid[i] = int64(cardinfo_3.CardId)
				cardinfo_3.CardId = 0
				cardinfo_3.CardAddr = 0

			}
			if cardlevel[i] == 4 {
				result = db.DB.Where("card_addr = ?", AddrRandomNum(cardlevel[i])).Take(&cardinfo_4)
				if result.Error != nil {
					fmt.Printf("read failed")
				}
				cardid[i] = int64(cardinfo_4.CardId)
				cardinfo_4.CardId = 0
				cardinfo_4.CardAddr = 0
			}
			if cardlevel[i] == 5 {
				result = db.DB.Where("card_addr = ?", AddrRandomNum(cardlevel[i])).Take(&cardinfo_5)
				if result.Error != nil {
					fmt.Printf("read failed")
				}
				cardid[i] = int64(cardinfo_5.CardId)
				cardinfo_5.CardId = 0
				cardinfo_5.CardAddr = 0
			}

		} else {
			break
		}

	}
	return cardid, result.Error

}
func RandomNum(frequency string) [10]int {
	var randomnum [10]int
	if frequency == "tentimes" {
		for i := 0; i < 10; i++ {
			randomnum[i] = rand.Intn(2000) + 1

		}
	}
	if frequency == "once" {
		randomnum[0] = rand.Intn(2000) + 1
		for i := 1; i < 10; i++ {
			randomnum[i] = 2001
		}
	}
	return randomnum
}
func Getalevel(randomnum [10]int) [10]int {
	var getalevel [10]int
	x := 0
	for _, i := range randomnum {
		if i <= 100 {
			getalevel[x] = 5
		}
		if i > 100 && i <= 600 {
			getalevel[x] = 4
		}
		if i > 600 && i <= 2000 {
			getalevel[x] = 3
		}
		if i == 2001 {
			getalevel[x] = 0
		}
		x++
	}
	return getalevel
}

func AddrRandomNum(cardlevel int) int {
	Cardlevel := int64(cardlevel)
	n := GetCardnum(Cardlevel)
	return rand.Intn(int(n)) + 1
}
func GetCardnum(level int64) int64 {
	db := dbconnect.DbInit()
	var tablelen int64
	switch level {
	case 3:
		{
			db.DB.Model(&datamodel.CardInfo3{}).Count(&tablelen)
			break
		}
	case 4:
		{
			db.DB.Model(&datamodel.CardInfo4{}).Count(&tablelen)
			break
		}
	case 5:
		{
			db.DB.Model(&datamodel.CardInfo5{}).Count(&tablelen)
			break
		}
	}
	log.Print(tablelen)
	return tablelen

}
