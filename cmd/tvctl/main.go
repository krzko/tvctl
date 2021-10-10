package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/krzko/tvctl/pkg/exchange"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"

	log "github.com/sirupsen/logrus"
)

var (
	buildVersion string
	// commit       string

	// binanceFuturesCoinApiURL = "https://dapi.binance.com"
	binanceFuturesUSDApiURL = "https://fapi.binance.com"
	binanceSpotApiURL       = "https://api.binance.com"
	ftxApiURL               = "https://ftx.com/api"
	uniSwapApiURL           = "https://api.thegraph.com"
)

func init() {
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	// Rainbow
	c := []color.Attribute{color.FgRed, color.FgGreen, color.FgYellow, color.FgMagenta, color.FgCyan, color.FgWhite, color.FgHiRed, color.FgHiGreen, color.FgHiYellow, color.FgHiBlue, color.FgHiMagenta, color.FgHiCyan, color.FgHiWhite}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c), func(i, j int) { c[i], c[j] = c[j], c[i] })
	c0 := color.New(c[0]).SprintFunc()
	c1 := color.New(c[1]).SprintFunc()
	c2 := color.New(c[2]).SprintFunc()
	c3 := color.New(c[3]).SprintFunc()
	c4 := color.New(c[4]).SprintFunc()
	appName := fmt.Sprintf("%s%s%s%s%s", c0("t"), c1("v"), c2("c"), c3("t"), c4("l"))

	app := &cli.App{
		Name:      appName,
		Usage:     "A command-line utility to interact with TradingView",
		UsageText: appName + " [global options] command [command options] [arguments...]",
		Version:   buildVersion,
		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Fprintf(c.App.Writer, "tvctl: Command not found: %q\n", command)
		},
	}

	app.Commands = cli.Commands{
		{
			Name:    "watchlist",
			Aliases: []string{"wl"},
			Usage:   "Commands related to watchlist features",
			Subcommands: []*cli.Command{
				{
					Name:    "generate",
					Aliases: []string{"g"},
					Usage:   "Generates a watchlist from exchanges",
					Flags: []cli.Flag{
						&cli.StringSliceFlag{
							Name:     "exchange",
							Usage:    "Comma seperated list of exchanges to use (default: \"binance\"",
							Aliases:  []string{"ex"},
							Required: false,
							Value:    cli.NewStringSlice("binance"),
						},
						&cli.StringFlag{
							Name:        "directory",
							Usage:       "The directory to save the files to",
							DefaultText: "./",
							Aliases:     []string{"d"},
							Required:    false,
							Value:       ".",
						},
					},
					Action: func(c *cli.Context) error {

						log.WithFields(log.Fields{
							"version": buildVersion,
						}).Info("Starting tvctl")

						log.Debug(len(c.String("exchange")))

						log.WithFields(log.Fields{
							"directory": c.String("directory"),
							"exchange":  c.String("exchange"),
						}).Info("Generating watchlist")

						if err := ensureDir(c.String("directory")); err != nil {
							log.Fatal("Directory creation failed with error: " + err.Error())
							os.Exit(1)
						}

						dir := ""
						if !strings.HasSuffix(c.String("directory"), "/") {
							dir = fmt.Sprintf("%v/", c.String("directory"))
						} else {
							dir = c.String("directory")
						}

						exBinName := "Binance"
						exBin := exchange.Binance{}

						// exBinPingFuturesCoin, err := exBin.PingFuturesCoin(binanceFuturesCoinApiURL, exBinName)
						// if err != nil {
						// 	log.WithFields(log.Fields{
						// 		"error": err,
						// 	}).Fatalln("Cannot contact " + exBinName)
						// }

						// if !exBinPingFuturesCoin {
						// 	log.WithFields(log.Fields{
						// 		"error": err,
						// 	}).Fatalln("Cannot contact " + exBinName)
						// }
						// log.Infof("Connected to %s: %v", exBinName, exBinPingFuturesCoin)

						exBinPingFuturesUSD, err := exBin.PingFuturesUSD(binanceFuturesUSDApiURL, exBinName)
						if err != nil {
							log.WithFields(log.Fields{
								"error": err,
							}).Fatalln("Cannot contact " + exBinName)
						}

						if !exBinPingFuturesUSD {
							log.WithFields(log.Fields{
								"error": err,
							}).Fatalln("Cannot contact " + exBinName)
						}
						log.Infof("Connected to %s: %v", exBinName, exBinPingFuturesUSD)

						exBinPingSpot, err := exBin.PingSpot(binanceSpotApiURL, exBinName)
						if err != nil {
							log.WithFields(log.Fields{
								"error": err,
							}).Fatalln("Cannot contact " + exBinName)
						}

						if !exBinPingSpot {
							log.WithFields(log.Fields{
								"error": err,
							}).Fatalln("Cannot contact " + exBinName)
						}
						log.Infof("Connected to %s: %v", exBinName, exBinPingSpot)

						// exFtxName := "FTX"
						// exFtx := exchange.Ftx{}
						// exFtxPing, err := exFtx.Ping(ftxApiURL, exFtxName)
						// if err != nil {
						// 	log.WithFields(log.Fields{
						// 		"error": err,
						// 	}).Error("Cannot contact " + exFtxName)
						// }
						// log.Infof("Connected to %s: %v", exFtxName, exFtxPing)

						// exUniSwapName := "UniSwap"
						// exUniSwap := exchange.UniSwap{}
						// exUniSwapPing, err := exUniSwap.Ping(uniSwapApiURL, exUniSwapName)
						// if err != nil {
						// 	log.WithFields(log.Fields{
						// 		"error": err,
						// 	}).Error("Cannot contact " + exUniSwapName)
						// }
						// log.Infof("Connected to %s: %v", exUniSwapName, exUniSwapPing)

						exBinSpotSymbols := exBin.GetSpotSymbols(binanceSpotApiURL, exBinName)
						exBin.ExportMarginSymbolsToDirectory(dir, exBinName, "USDT", exBinSpotSymbols)
						exBin.ExportSpotSymbolsToDirectory(dir, exBinName, "USDT", exBinSpotSymbols)
						exBin.ExportSpotSymbolsToDirectory(dir, exBinName, "BTC", exBinSpotSymbols)
						exBin.ExportSpotSymbolsToDirectory(dir, exBinName, "ETH", exBinSpotSymbols)

						// exBinFuturesCoinSymbols := exBin.GetFuturesCoinSymbols(binanceFuturesCoinApiURL, exBinName)
						// exBin.ExportFuturesCoinPerpSymbolsToDirectory(dir, exBinName, "USD", exBinFuturesCoinSymbols)
						// exBin.ExportFuturesCoinCurrentQuarterlySymbolsToDirectory(dir, exBinName, "USD", exBinFuturesCoinSymbols)
						// exBin.ExportFuturesCoinNextQuarterlySymbolsToDirectory(dir, exBinName, "USD", exBinFuturesCoinSymbols)

						exBinFuturesUSDSymbols := exBin.GetFuturesUSDSymbols(binanceFuturesUSDApiURL, exBinName)
						exBin.ExportFuturesUSDPerpSymbolsToDirectory(dir, exBinName, "USDT", exBinFuturesUSDSymbols)
						exBin.ExportFuturesUSDQuarterlySymbolsToDirectory(dir, exBinName, "USDT", exBinFuturesUSDSymbols)

						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ensureDir(dirName string) error {
	err := os.Mkdir(dirName, 0755)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}
