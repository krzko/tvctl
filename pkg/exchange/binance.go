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

// https://binance-docs.github.io/apidocs/spot/en/#market-data-endpoints
// https://binance-docs.github.io/apidocs/futures/en/#market-data-endpoints
// https://binance-docs.github.io/apidocs/delivery/en/#market-data-endpoints

type Binance struct{}

type BinanceFuturesCoinData struct {
	TimeZone     string                     `json:"timezone"`
	ServerTime   int64                      `json:"serverTime"`
	SymbolEntity []BinanceFuturesCoinSymbol `json:"symbols"`
}

type BinanceFuturesCoinSymbol struct {
	Symbol         string `json:"symbol"`
	ContractStatus string `json:"contractStatus"`
	BaseAsset      string `json:"baseAsset"`
	QuoteAsset     string `json:"quoteAsset"`
	ContractType   string `json:"contractType"`
}

type BinanceFuturesUSDData struct {
	TimeZone     string                    `json:"timezone"`
	ServerTime   int64                     `json:"serverTime"`
	SymbolEntity []BinanceFuturesUSDSymbol `json:"symbols"`
}

type BinanceFuturesUSDSymbol struct {
	Symbol       string `json:"symbol"`
	Status       string `json:"status"`
	BaseAsset    string `json:"baseAsset"`
	QuoteAsset   string `json:"quoteAsset"`
	ContractType string `json:"contractType"`
}

type BinanceSpotData struct {
	TimeZone     string              `json:"timezone"`
	ServerTime   int64               `json:"serverTime"`
	SymbolEntity []BinanceSpotSymbol `json:"symbols"`
}

type BinanceSpotSymbol struct {
	Symbol                 string `json:"symbol"`
	Status                 string `json:"status"`
	BaseAsset              string `json:"baseAsset"`
	QuoteAsset             string `json:"quoteAsset"`
	IsSpotTradingAllowed   bool   `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed bool   `json:"isMarginTradingAllowed"`
}

func (bi Binance) ExportFuturesCoinCurrentQuarterlySymbolsToDirectory(dir, name, baseAsset string, sym []BinanceFuturesCoinSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Futures Current Quarterly COIN-M symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.ContractStatus == "TRADING" && x.ContractType == "CURRENT_QUARTER" && x.QuoteAsset == baseAsset {
			val := fmt.Sprintf("%v:%v\n", strings.ToUpper(name), strings.ToUpper(x.Symbol))
			log.Debug(val)
			symbList = append(symbList, val)
		}
	}
	sort.Strings(symbList)
	log.Debug(strings.Join(symbList, " "))

	fn := fmt.Sprintf("%v-Coin-Futures-Current-Quarter-%v.txt", name, baseAsset)
	f, err := os.Create(dir + fn)
	log.WithFields(log.Fields{
		"file": dir + fn,
	}).Infof("Writing %v Futures Current Quarterly COIN-M symbols to file ", name)

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

func (bi Binance) ExportFuturesCoinNextQuarterlySymbolsToDirectory(dir, name, baseAsset string, sym []BinanceFuturesCoinSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Futures Next Quarterly COIN-M symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.ContractStatus == "TRADING" && x.ContractType == "NEXT_QUARTER" && x.QuoteAsset == baseAsset {
			val := fmt.Sprintf("%v:%v\n", strings.ToUpper(name), strings.ToUpper(x.Symbol))
			log.Debug(val)
			symbList = append(symbList, val)
		}
	}
	sort.Strings(symbList)
	log.Debug(strings.Join(symbList, " "))

	fn := fmt.Sprintf("%v-Coin-Futures-Next-Quarter-%v.txt", name, baseAsset)
	f, err := os.Create(dir + fn)
	log.WithFields(log.Fields{
		"file": dir + fn,
	}).Infof("Writing %v Futures Next Quarterly COIN-M symbols to file ", name)

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

func (bi Binance) ExportFuturesCoinPerpSymbolsToDirectory(dir, name, baseAsset string, sym []BinanceFuturesCoinSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Futures Perpetual COIN-M symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.ContractStatus == "TRADING" && x.ContractType == "PERPETUAL" && x.QuoteAsset == baseAsset {
			val := fmt.Sprintf("%v:%v\n", strings.ToUpper(name), strings.ToUpper(x.Symbol))
			log.Debug(val)
			symbList = append(symbList, val)
		}
	}
	sort.Strings(symbList)
	log.Debug(strings.Join(symbList, " "))

	fn := fmt.Sprintf("%v-Coin-Futures-Perp-%v.txt", name, baseAsset)
	f, err := os.Create(dir + fn)
	log.WithFields(log.Fields{
		"file": dir + fn,
	}).Infof("Writing %v Futures Perpetual COIN-M symbols to file ", name)

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

func (bi Binance) ExportFuturesUSDQuarterlySymbolsToDirectory(dir, name, baseAsset string, sym []BinanceFuturesUSDSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Futures Quarterly USD-M symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.Status == "TRADING" && x.ContractType == "CURRENT_QUARTER" && x.QuoteAsset == baseAsset {
			val := fmt.Sprintf("%v:%v\n", strings.ToUpper(name), strings.ToUpper(x.Symbol))
			log.Debug(val)
			symbList = append(symbList, val)
		}
	}
	sort.Strings(symbList)
	log.Debug(strings.Join(symbList, " "))

	fn := fmt.Sprintf("%v-Futures-Quarter-%v.txt", name, baseAsset)
	f, err := os.Create(dir + fn)
	log.WithFields(log.Fields{
		"file": dir + fn,
	}).Infof("Writing %v Futures Quarterly USD-M symbols to file ", name)

	if err != nil {
		log.WithFields(log.Fields{
			"dir":   dir,
			"error": err,
		}).Fatal("Cannot write to file")
	}

	for _, sle := range symbList {
		sym := fmt.Sprintf("%vPERP,", strings.Trim(sle, " \r\n"))
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

func (bi Binance) ExportFuturesUSDPerpSymbolsToDirectory(dir, name, baseAsset string, sym []BinanceFuturesUSDSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Futures Perpetual USD-M symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.Status == "TRADING" && x.ContractType == "PERPETUAL" && x.QuoteAsset == baseAsset {
			val := fmt.Sprintf("%v:%v\n", strings.ToUpper(name), strings.ToUpper(x.Symbol))
			log.Debug(val)
			symbList = append(symbList, val)
		}
	}
	sort.Strings(symbList)
	log.Debug(strings.Join(symbList, " "))

	fn := fmt.Sprintf("%v-Futures-Perp-%v.txt", name, baseAsset)
	f, err := os.Create(dir + fn)
	log.WithFields(log.Fields{
		"file": dir + fn,
	}).Infof("Writing %v Futures Perpetual USD-M symbols to file ", name)

	if err != nil {
		log.WithFields(log.Fields{
			"dir":   dir,
			"error": err,
		}).Fatal("Cannot write to file")
	}

	for _, sle := range symbList {
		sym := fmt.Sprintf("%vPERP,", strings.Trim(sle, " \r\n"))
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

func (bi Binance) ExportMarginSymbolsToDirectory(dir, name, baseAsset string, sym []BinanceSpotSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Margin symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.Status == "TRADING" && x.IsMarginTradingAllowed && x.QuoteAsset == baseAsset {
			val := fmt.Sprintf("%v:%v\n", strings.ToUpper(name), strings.ToUpper(x.Symbol))
			log.Debug(val)
			symbList = append(symbList, val)
		}
	}
	sort.Strings(symbList)
	log.Debug(strings.Join(symbList, " "))

	fn := fmt.Sprintf("%v-Margin-%v.txt", name, baseAsset)
	f, err := os.Create(dir + fn)
	log.WithFields(log.Fields{
		"file": dir + fn,
	}).Infof("Writing %v Margin symbols to file ", name)

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

func (bi Binance) ExportSpotSymbolsToDirectory(dir, name, baseAsset string, sym []BinanceSpotSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Spot symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.Status == "TRADING" && x.IsSpotTradingAllowed && x.QuoteAsset == baseAsset {
			val := fmt.Sprintf("%v:%v\n", strings.ToUpper(name), strings.ToUpper(x.Symbol))
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

func (bi Binance) GetFuturesCoinSymbols(url, name string) []BinanceFuturesCoinSymbol {
	path := "/dapi/v1/exchangeInfo"

	log.WithFields(log.Fields{
		"url": url + path,
	}).Infof("Getting %v Futures COIN-M symbols ", name)

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

	var dynamic BinanceFuturesCoinData
	json.Unmarshal([]byte(resp.Body()), &dynamic)

	log.Infof("%v Futures server time %v", name, dynamic.ServerTime)
	log.Infof("%v Futures symbols found: %v", name, len(dynamic.SymbolEntity))

	return dynamic.SymbolEntity
}

func (bi Binance) GetFuturesUSDSymbols(url, name string) []BinanceFuturesUSDSymbol {
	path := "/fapi/v1/exchangeInfo"

	log.WithFields(log.Fields{
		"url": url + path,
	}).Infof("Getting %v Futures USD-M symbols ", name)

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

	var dynamic BinanceFuturesUSDData
	json.Unmarshal([]byte(resp.Body()), &dynamic)

	log.Infof("%v Futures server time %v", name, dynamic.ServerTime)
	log.Infof("%v Futures symbols found: %v", name, len(dynamic.SymbolEntity))

	return dynamic.SymbolEntity
}

func (bi Binance) GetSpotSymbols(url, name string) []BinanceSpotSymbol {
	path := "/api/v3/exchangeInfo"

	log.WithFields(log.Fields{
		"url": url + path,
	}).Infof("Getting %v Spot symbols ", name)

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

	var dynamic BinanceSpotData
	json.Unmarshal([]byte(resp.Body()), &dynamic)

	log.Infof("%v Spot server time %v", name, dynamic.ServerTime)
	log.Infof("%v Spot symbols found: %v", name, len(dynamic.SymbolEntity))

	return dynamic.SymbolEntity
}

func (bi Binance) PingFuturesCoin(url, name string) (bool, error) {
	path := "/dapi/v1/ping"

	log.WithFields(log.Fields{
		"url": url + path,
	}).Infof("Pinging %v Futures COIN-M", name)

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get(url + path)

	if err != nil {
		log.WithFields(log.Fields{
			"status": resp.Status(),
			"time":   resp.Time(),
		}).Error("Cannot connect to " + name)
		return false, err
	}

	if resp.StatusCode() != 200 {
		log.WithFields(log.Fields{
			"status": resp.Status(),
			"time":   resp.Time(),
		}).Fatalln("Cannot connect to " + name)
		return false, err
	}

	log.WithFields(log.Fields{
		"status": resp.Status(),
		"time":   resp.Time(),
	}).Infof("Successfully pinged %v Futures COIN-M", name)

	return true, nil
}

func (bi Binance) PingFuturesUSD(url, name string) (bool, error) {
	path := "/fapi/v1/ping"

	log.WithFields(log.Fields{
		"url": url + path,
	}).Infof("Pinging %v Futures USD-M", name)

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get(url + path)

	if err != nil {
		log.WithFields(log.Fields{
			"status": resp.Status(),
			"time":   resp.Time(),
		}).Error("Cannot connect to " + name)
		return false, err
	}

	if resp.StatusCode() != 200 {
		log.WithFields(log.Fields{
			"status": resp.Status(),
			"time":   resp.Time(),
		}).Fatalln("Cannot connect to " + name)
		return false, err
	}

	log.WithFields(log.Fields{
		"status": resp.Status(),
		"time":   resp.Time(),
	}).Infof("Successfully pinged %v Futures USD-M", name)

	return true, nil
}

func (bi Binance) PingSpot(url, name string) (bool, error) {
	path := "/api/v3/ping"

	log.WithFields(log.Fields{
		"url": url + path,
	}).Infof("Pinging %v Spot", name)

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get(url + path)

	if err != nil {
		log.WithFields(log.Fields{
			"status": resp.Status(),
			"time":   resp.Time(),
		}).Error("Cannot connect to " + name)
		return false, err
	}

	if resp.StatusCode() != 200 {
		log.WithFields(log.Fields{
			"status": resp.Status(),
			"time":   resp.Time(),
		}).Fatalln("Cannot connect to " + name)
		return false, err
	}

	log.WithFields(log.Fields{
		"status": resp.Status(),
		"time":   resp.Time(),
	}).Infof("Successfully pinged %v Spot", name)

	return true, nil
}
