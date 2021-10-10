# tvctl

📈 A command-line utility to interact with [TradingView](https://www.tradingview.com).

This utility aims to improve usability of tedius and repetative tasks, such as **watchlist** generation and hopefully in the future, **alert** management. New features will be added as and when an TradingView expose an endpoint to manage these features.

Currently supported features are:

* Watchlists
  * Binance: Spot, Margin, Futures (USD-M) watchlists
  * FTX: in progress
  * Uniwap: in progress

## Getting Started

Running `tvctl` is availabile through several methods. You can download it as a binary from GitHub releases, running it as a distroless docker image or building it from source.

```sh
NAME:
   tvctl - A command-line utility to interact with TradingView

USAGE:
   tvctl [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

COMMANDS:
   watchlist, wl  Commands related to watchlist features
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Run

### Watchlist

To generate (export) a watchlist, use the following command:

```sh
# Long
tvctl watchlist generate

# Short
tvctl wl g
```

For an example of the generated output, view the [export](https://github.com/krzko/tvctl/tree/main/export) directory.

## Download Binary

Download the latest [release](https://github.com/krzko/tvctl/releases).

## Docker

Run it via a docker distroless image:

```sh
docker run --rm \
    -v "$(pwd)":/save \
    ghcr.io/krzko/tvctl:latest watchlist generate -directory /save
```
