package exchange

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"

	log "github.com/sirupsen/logrus"
)

// https://thegraph.com/legacy-explorer/subgraph/uniswap/uniswap-v2
// TODO: krzko: https://thegraph.com/legacy-explorer/subgraph/uniswap/uniswap-v3, not supported yet by TradigView
// TODO: krzko: Base assets, WETH, USDT, USDC, DAI, TRU

type UniSwap struct{}

func (un UniSwap) GetSpotSymbols(url, name string) []BinanceSpotSymbol {
	path := "/subgraphs/name/uniswap/uniswap-v2"

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

func (v UniSwap) Ping(url, name string) (bool, error) {
	path := "/subgraphs/name/uniswap/uniswap-v2"

	log.WithFields(log.Fields{
		"url": url + path,
	}).Info("Pinging " + name)

	// Create a Resty Client
	client := resty.New()

	// query := `{"query": "{ pairs(first: 1) { id } }"}`
	// query := `{"query": "{ pairs(first: 1, where: {reserveUSD_gt: "1000000", volumeUSD_gt: "50000"}, orderBy: reserveUSD, orderDirection: desc) { token0 { symbol } token1 { symbol } } }" }`
	query := `{"query": "{ pairs(first: 10, where: {reserveUSD_gt: "1000000", volumeUSD_gt: "50000"}, orderBy: reserveUSD, orderDirection: desc) { id token0 { id symbol } token1 { id symbol } reserveUSD volumeUSD } }"}`

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(query).
		Post(url + path)

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
