package offers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetRequest(trade, bank, status string) {
	url := GetLinkOffer(ChangeTrade(trade), ChangeBank(bank), ChangeStatus(status))
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("can't get response:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("can't read body of response:", err)
	}

	var offerBody OfferBody

	err = json.Unmarshal(body, &offerBody)

	if err != nil {
		log.Fatal("this: ", err)
	}

	if offerBody.ResponseStatus == 200 {
		for i := 0; i < len(offerBody.Data.Data); i++ {
			data := offerBody.Data.Data[i]
			price := data.Price
			user := data.User.Nickname
			minLimit := data.LimitMin
			bank := data.PaymentMethod.Name
			fmt.Printf("[Price: %.2f, User: %s, LimitMin: %.2f, Bank: %s]\n\n", price, user, minLimit, bank)
		}
	} else {
		log.Fatalf("response status: %d", offerBody.ResponseStatus)
	}
}
