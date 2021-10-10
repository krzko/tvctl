package exchange

// https://thegraph.com/legacy-explorer/subgraph/uniswap/uniswap-v2

import (
	"github.com/go-resty/resty/v2"

	log "github.com/sirupsen/logrus"
)

type UniSwap struct{}

func (v UniSwap) Ping(url, name string) (bool, error) {
	path := "/subgraphs/name/uniswap/uniswap-v2"

	log.WithFields(log.Fields{
		"url": url + path,
	}).Info("Pinging " + name)

	// Create a Resty Client
	client := resty.New()

	// query := `{"query": "{ pairs(first: 1, where: {reserveUSD_gt: "1000000", volumeUSD_gt: "50000"}, orderBy: reserveUSD, orderDirection: desc) { token0 { symbol } token1 { symbol } } }" }`

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"query": "{ pairs(first: 1) { id } } } "}`).
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
