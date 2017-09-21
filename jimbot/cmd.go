package jimbot

import (
	"strings"
)

const info = "Under development"

// ProcessCmd : handles bot commands
func ProcessCmd(command string, userID int64) string {
	switch command {
	case "start":
		msg := start(userID)
		return msg
	case "google":
		msg := googleSearch(command)
		return msg
	case "stat":
		return info
	case "translate":
		return info
	case "pic":
		return info
	case "3_day_forecast":
		return info
	case "weather":
		return info
	case "prices":
		msg := prices()
		return msg
	default:
		return "Unknown command"
	}
}

func start(userID int64) string {
	var msg string
	switch userID {
	case ReadConfig().GFID:
		msg = "Hi, I'm your Telegram bot,\n"
		msg += "hope I'll be loved,\n"
		msg += "if not, well... blame him,"
		msg += "and... I love you two\n"
	case ReadConfig().BFID:
		msg = "Hi, I'm your Telegram bot, and...\n"
		msg += "I'll always be here with you,\n"
		msg += "and... I love you two\n"
	default:
		msg = "There must be something wrong...\n"
	}
	msg += KISS
	return msg
}

func prices() string {
	btcPrice := GetPrice("BTC")
	msg := HII + " I got this list\n`"
	msg += strings.Repeat("-", 35)
	msg += "\n"
	msg += btcPrice.CoinName + " -> USD : " + btcPrice.PriceInUSD + "\n"
	msg += getAltcoinPrices("XMR")
	msg += getAltcoinPrices("ETH")
	msg += getAltcoinPrices("ETC")
	msg += getAltcoinPrices("BCC")
	msg += "`"
	return msg
}

func getAltcoinPrices(coin string) string {
	coinPrice := GetPrice(coin)
	var msg string
	msg += coinPrice.CoinName + " -> USD : " + coinPrice.PriceInUSD + "\n"
	msg += coinPrice.CoinName + " -> BTC : " + coinPrice.PriceInBTC + "\n"
	return msg
}

func googleSearch(command string) string {
	query := "test"
	//query := strings.Split(command, "/google ")[0]
	return Search(query, false)
}
