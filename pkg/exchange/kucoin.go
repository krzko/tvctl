package exchange

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/go-resty/resty/v2"

	log "github.com/sirupsen/logrus"
)

// https://docs.kucoin.com/#base-url

type Kucoin struct{}

type KucoinSpotData struct {
	Code         string         `json:"code"`
	SymbolEntity []KucoinSymbol `json:"data"`
}

type KucoinSymbol struct {
	Symbol          string `json:"symbol"`
	Name            string `json:"name"`
	BaseCurrency    string `json:"baseCurrency"`
	QuoteCurrency   string `json:"quoteCurrency"`
	FeeCurrency     string `json:"feeCurrency"`
	Market          string `json:"market"`
	BaseMinSize     string `json:"baseMinSize"`
	QuoteMinSize    string `json:"quoteMinSize"`
	BaseMaxSize     string `json:"baseMaxSize"`
	QuoteMaxSize    string `json:"quoteMaxSize"`
	BaseIncrement   string `json:"baseIncrement"`
	QuoteIncrement  string `json:"quoteIncrement"`
	PriceIncrement  string `json:"priceIncrement"`
	PriceLimitRate  string `json:"priceLimitRate"`
	IsMarginEnabled bool   `json:"isMarginEnabled"`
	EnableTrading   bool   `json:"enableTrading"`
}

func (ku Kucoin) ExportSpotSymbolsToDirectory(dir, name, baseAsset string, sym []KucoinSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Spot symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.EnableTrading && x.QuoteCurrency == baseAsset {
			val := fmt.Sprintf("%v:%v%v\n", strings.ToUpper(name), strings.ToUpper(x.BaseCurrency), strings.ToUpper(x.QuoteCurrency))
			log.Debug(val)
			symbList = append(symbList, val)
		}
	}
	sort.Strings(symbList)
	log.Debug(strings.Join(symbList, " "))

	fn := fmt.Sprintf("%v-Spot-%v.txt", name, baseAsset)
	f, err := os.Create(dir + fn)
	log.WithFields(log.Fields{
		"file": dir + fn,
	}).Infof("Writing %v Spot symbols to file ", name)

	if err != nil {
		log.WithFields(log.Fields{
			"dir":   dir,
			"error": err,
		}).Fatal("Cannot write to file")
	}

	for _, sle := range symbList {
		sym := fmt.Sprintf("%v,", strings.Trim(sle, " \r\n"))
		log.Debug(sym)
		_, err := f.WriteString(sym + "\n")
		if err != nil {
			log.WithFields(log.Fields{
				"symbol": sle,
				"error":  err,
			}).Fatal("Cannot write symbol to file")
		}
	}
}

func (ku Kucoin) GetSymbols(url, name string) []KucoinSymbol {
	path := "/api/v1/symbols"

	log.WithFields(log.Fields{
		"url": url + path,
	}).Infof("Getting %v symbols ", name)

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get(url + path)

	if err != nil {
		log.WithFields(log.Fields{
			"status": resp.Status(),
			"time":   resp.Time(),
		}).Error("Cannot get symbols for " + name)
	}

	var dynamic KucoinSpotData
	json.Unmarshal([]byte(resp.Body()), &dynamic)

	log.Infof("%v symbols found: %v", name, len(dynamic.SymbolEntity))

	return dynamic.SymbolEntity
}

func (ku Kucoin) Ping(url, name string) (bool, error) {
	path := "/api/v1/market/orderbook/level1?symbol=BTC-USDT"

	log.WithFields(log.Fields{
		"url": url + path,
	}).Info("Pinging " + name)

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get(url + path)

	if err != nil {
		log.WithFields(log.Fields{
			"url": url + path,
		}).Error("Cannot connect to " + name)
		return false, err
	}

	log.WithFields(log.Fields{
		"status": resp.Status(),
		"time":   resp.Time(),
	}).Info("Successfully pinged " + name)

	return true, nil
}
