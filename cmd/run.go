package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"totalcoin_bot/internal/offers"
)

//todo: перенести в переменные
var (
	sell, buy = "buy", "sell"

	sberbank, tinkoff, visa = "1", "3", "332"
	qiwi                    = "72"
	alphaBank               = "5"
	vtb                     = "4"
	umoney                  = "6"
	openBank                = "377"
	gazPromBank             = "423"
	rosBank                 = "434"

	pro   = "true"
	noPro = "false"
)

func main() {
	url := offers.GetLinkOffer(sell, qiwi, pro)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("can't get response:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("can't read body of response:", err)
	}

	var offerBody offers.OfferBody

	err = json.Unmarshal(body, &offerBody)

	if err != nil {
		log.Fatal("this: ", err)
	}

	fmt.Println(offerBody.ResponseStatus, offerBody.ResponseMessage)

	// todo: перенести в переменные

	for i := 0; i < len(offerBody.Data.Data); i++ {
		data := offerBody.Data.Data[i]
		price := data.Price
		user := data.User.Nickname
		minLimit := data.LimitMin
		fmt.Printf("[Price: %.2f, User: %s, LimitMin: %.2f]\n\n", price, user, minLimit)
	}
}
