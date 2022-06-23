package offers

func GetLinkOffer(trade, offer, state string) string {
	var linkOffer = "https://web.totalcoin.io/services/otc/api/v2/offer/list/" + trade + "/0/8?currency=RUB&paymentMethodCurrency=" + offer + "&pro=" + state
	return linkOffer
}
