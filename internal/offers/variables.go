package offers

import (
	"fmt"
)

var Trade = map[string]string{
	"sell": "buy",
	"buy":  "sell",
}

func ChangeTrade(trade string) string {
	sell := Trade["sell"]
	buy := Trade["buy"]

	switch trade {
	case sell:
		return sell
	case buy:
		return buy
	default:
		fmt.Println("Need trade! Buy or Sell")
		return ""
	}
}

var Banks = map[string]string{
	"sberbank":    "1",
	"tinkoff":     "3",
	"visa":        "332",
	"qiwi":        "72",
	"alphaBank":   "5",
	"vtb":         "4",
	"umoney":      "6",
	"openBank":    "377",
	"gazPromBank": "423",
	"rosBank":     "434",
}

func ChangeBank(bank string) string {
	sberbank := Banks["sberbank"]
	tinkoff := Banks["tinkoff"]
	visa := Banks["visa"]
	qiwi := Banks["qiwi"]
	alphaBank := Banks["alphaBank"]
	vtb := Banks["vtb"]
	umoney := Banks["umoney"]
	openBank := Banks["openBank"]
	gazPromBank := Banks["gazPromBank"]
	rosBank := Banks["rosBank"]

	switch bank {
	case "sberbank":
		return sberbank
	case "tinkoff":
		return tinkoff
	case "visa":
		return visa
	case "qiwi":
		return qiwi
	case "alphaBank":
		return alphaBank
	case "vtb":
		return vtb
	case "umoney":
		return umoney
	case "openBank":
		return openBank
	case "gazPromBank":
		return gazPromBank
	case "rosBank":
		return rosBank
	default:
		fmt.Println("Need bank!")
		return ""
	}
}

var statusPro = map[string]string{
	"pro":   "true",
	"noPro": "false",
}

func ChangeStatus(status string) string {
	pro := statusPro["pro"]
	noPro := statusPro["noPro"]

	switch status {
	case "pro":
		return pro
	case "noPro":
		return noPro
	default:
		fmt.Println("Need status trade")
		return ""
	}
}
