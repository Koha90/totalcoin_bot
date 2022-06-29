package main

import (
	"totalcoin_bot/internal/config"
	"totalcoin_bot/internal/offers"
)

func main() {
	//fmt.Println(offers.ChangeTrade())
	offers.GetRequest("buy", "qiwi", "pro")
	config.GetUserProfile()
}
