package offers

type OfferBody struct {
	ResponseStatus  int    `json:"responseStatus"`
	ResponseMessage string `json:"responseMessage"`
	Data            `json:"data"`
}

type Data struct {
	Data          []DataBody `json:"data"`
	TotalPages    int        `json:"totalPages"`
	TotalElements int        `json:"totalElements"`
}

type DataBody struct {
	Id                      string `json:"id"`
	PaymentMethod           `json:"paymentMethod"`
	Price                   float64 `json:"price"`
	LimitMin                float64 `json:"limitMin"`
	LimitMax                float64 `json:"limitMax"`
	Active                  bool    `json:"active"`
	User                    `json:"user"`
	Cryptocurrency          string `json:"cryptocurrency"`
	Currency                `json:"currency"`
	Type                    string  `json:"type"`
	Terms                   string  `json:"terms"`
	AutoLimit               bool    `json:"autoLimit"`
	FixedPrice              float64 `json:"fixedPrice"`
	PartnerLimitType        string  `json:"partnerLimitType"`
	PartnerLimitDealsTotal  string  `json:"partnerLimitDealsTotal"`
	PartnerLimitTradeVolume string  `json:"partnerLimitTradeVolume"`
	PartnerLimitPercent     int     `json:"partnerLimitPercent,omitempty"`
}

type PaymentMethod struct {
	Id                       int     `json:"id"`
	Name                     string  `json:"name"`
	CurrencyId               string  `json:"currencyId"`
	CurrencyDisplayNameShort string  `json:"currencyDisplayNameShort"`
	Icon                     string  `json:"icon"`
	BestPriceSell            float64 `json:"bestPriceSell"`
	BestPriceBuy             float64 `json:"bestPriceBuy"`
	Sort                     int     `json:"sort"`
	PaymentMethodId          int     `json:"paymentMethodId"`
}

type User struct {
	Id                     string  `json:"id"`
	Nickname               string  `json:"nickname"`
	ConfirmedCount         int     `json:"confirmedCount"`
	RefusedCount           int     `json:"refusedCount"`
	FirstDealDate          int64   `json:"firstDealDate"`
	MedianResponseTime     int     `json:"medianResponseTime"`
	OkReviewCount          int     `json:"okReviewCount"`
	BadReviewCount         int     `json:"badReviewCount"`
	TextReviewCount        int     `json:"textReviewCount"`
	TotalValueTrade        float64 `json:"totalValueTrade"`
	TradeVolume            `json:"tradeVolume"`
	Currency               `json:"currency"`
	TradeEnabled           bool `json:"tradeEnabled"`
	CurrentUser            bool `json:"currentUser"`
	IsNickChanged          bool `json:"isNickChanged"`
	TimeoutAccept          int  `json:"timeoutAccept"`
	TimeoutTransferConfirm int  `json:"timeoutTransferConfirm"`
	TimeoutOpenDispute     int  `json:"timeoutOpenDispute"`
	UserBlockedTimes       int  `json:"userBlockedTimes"`
	Verified               bool `json:"verified"`
	ProTrader              bool `json:"proTrader"`
	IsFavoriteUser         bool `json:"isFavoriteUser"`
	NickChanged            bool `json:"nickChanged"`
	FavoriteUser           bool `json:"favoriteUser"`
}

type Currency struct {
	Id               string `json:"id"`
	DisplayNameShort string `json:"displayNameShort"`
}

type TradeVolume struct {
	BTC  float64 `json:"BTC"`
	ETH  float64 `json:"ETH,omitempty"`
	LTC  float64 `json:"LTC,omitempty"`
	USDT float64 `json:"USDT,omitempty"`
}

type CurrencyUser struct {
	Id               string `json:"id"`
	DisplayNameShort string `json:"displayNameShort"`
}
