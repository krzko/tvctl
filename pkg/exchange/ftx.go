package exchange

// https://docs.ftx.com/#rest-api

import (
	"github.com/go-resty/resty/v2"

	log "github.com/sirupsen/logrus"
)

type Ftx struct{}

func (v Ftx) Ping(url, name string) (bool, error) {
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
