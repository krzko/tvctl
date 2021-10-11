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

// https://docs.ftx.com/#rest-api

type Ftx struct{}

type FtxData struct {
	Success      bool        `json:"success"`
	SymbolEntity []FtxSymbol `json:"result"`
}

type FtxSymbol struct {
	Name            string `json:"name"`
	Enabled         bool   `json:"enabled"`
	Type            string `json:"type"`
	BaseCurrency    string `json:"baseCurrency,omitempty"`
	QuoteCurrency   string `json:"quoteCurrency,omitempty"`
	Underlying      string `json:"underlying,omitempty"`
	TokenizedEquity bool   `json:"tokenizedEquity,omitempty"`
}

func (ft Ftx) ExportFuturesPerpSymbolsToDirectory(dir, name, baseAsset string, sym []FtxSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Futures Perpetual symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.Enabled && x.Type == "future" && strings.HasSuffix(x.Name, "PERP") {
			val := fmt.Sprintf("%v:%vPERP\n", strings.ToUpper(name), strings.ToUpper(x.Underlying))
			log.Debug(val)
			symbList = append(symbList, val)
		}
	}
	sort.Strings(symbList)
	log.Debug(strings.Join(symbList, " "))

	fn := fmt.Sprintf("%v-Futures-%v.txt", name, baseAsset)
	f, err := os.Create(dir + fn)
	log.WithFields(log.Fields{
		"file": dir + fn,
	}).Infof("Writing %v Futures Perpetual symbols to file ", name)

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

func (ft Ftx) ExportFuturesQuarterlySymbolsToDirectory(dir, name string, sym []FtxSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Futures Quarterly symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.Enabled && x.Type == "future" && !strings.HasSuffix(x.Name, "PERP") {
			n := strings.Replace(x.Name, "-", "", -1)
			val := fmt.Sprintf("%v:%v\n", strings.ToUpper(name), strings.ToUpper(n))
			log.Debug(val)
			symbList = append(symbList, val)
		}
	}
	sort.Strings(symbList)
	log.Debug(strings.Join(symbList, " "))

	fn := fmt.Sprintf("%v-Futures-Quarterly.txt", name)
	f, err := os.Create(dir + fn)
	log.WithFields(log.Fields{
		"file": dir + fn,
	}).Infof("Writing %v Futures Perpetual symbols to file ", name)

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

func (ft Ftx) ExportSpotSymbolsToDirectory(dir, name, baseAsset string, sym []FtxSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Spot symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.Enabled && x.Type == "spot" && x.QuoteCurrency == baseAsset && !x.TokenizedEquity {
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

func (ft Ftx) ExportStockQuarterlySymbolsToDirectory(dir, name string, sym []FtxSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Stock Quarterly symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.Enabled && x.Type == "future" && x.TokenizedEquity {
			n := strings.Replace(x.Name, "-", "", -1)
			val := fmt.Sprintf("%v:%v\n", strings.ToUpper(name), strings.ToUpper(n))
			log.Debug(val)
			symbList = append(symbList, val)
		}
	}
	sort.Strings(symbList)
	log.Debug(strings.Join(symbList, " "))

	fn := fmt.Sprintf("%v-Stock-Quarterly.txt", name)
	f, err := os.Create(dir + fn)
	log.WithFields(log.Fields{
		"file": dir + fn,
	}).Infof("Writing %v Futures Perpetual symbols to file ", name)

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

func (ft Ftx) ExportStockSpotSymbolsToDirectory(dir, name, baseAsset string, sym []FtxSymbol) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Infof("Exporting %v Stock Spot symbols", name)

	symbList := []string{}
	for _, x := range sym {
		if x.Enabled && x.Type == "spot" && x.QuoteCurrency == baseAsset && x.TokenizedEquity {
			val := fmt.Sprintf("%v:%v%v\n", strings.ToUpper(name), strings.ToUpper(x.BaseCurrency), strings.ToUpper(x.QuoteCurrency))
			log.Debug(val)
			symbList = append(symbList, val)
		}
	}
	sort.Strings(symbList)
	log.Debug(strings.Join(symbList, " "))

	fn := fmt.Sprintf("%v-Stock-Spot-%v.txt", name, baseAsset)
	f, err := os.Create(dir + fn)
	log.WithFields(log.Fields{
		"file": dir + fn,
	}).Infof("Writing %v Stock Spot symbols to file ", name)

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

func (ft Ftx) GetSymbols(url, name string) []FtxSymbol {
	path := "/markets"

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

	var dynamic FtxData
	json.Unmarshal([]byte(resp.Body()), &dynamic)

	log.Infof("%v symbols found: %v", name, len(dynamic.SymbolEntity))

	return dynamic.SymbolEntity
}

func (ft Ftx) Ping(url, name string) (bool, error) {
	path := "/markets"

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
