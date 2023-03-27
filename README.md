# tvctl

📈 A command-line utility to interact with [TradingView](https://www.tradingview.com).

This utility aims to improve usability of tedius and repetative tasks, such as **watchlist** generation and hopefully in the future, **alert** management. New features will be added as and when an TradingView expose an endpoint to manage these features.

Currently supported features are:

* Watchlists
  * Binance: Spot, Margin, Futures (USD-M) watchlists
  * ~~FTX: Spot, Futures, Stocks~~ 🤡🌎
  * KuCoin: Spot

## Getting Started

Running `tvctl` is availabile through several methods. You can using `brew`, download it as a binary from GitHub releases, or running it as a distroless docker image.

### brew

Install [brew](https://brew.sh/) and then run:

```sh
brew install krzko/tap/tvctl
```

### Download Binary

Download the latest version from the [Releases](https://github.com/krzko/tvctl/releases) page.

### Docker

Run it via a docker distroless image:

```sh
docker run --rm \
    -v "$(pwd)":/save \
    ghcr.io/krzko/tvctl:latest watchlist generate -directory /save
```

To see all the tags view the [Packages](https://github.com/krzko/tvctl/pkgs/container/tvctl) page.

## Run

```sh
NAME:
   tvctl - A command-line utility to interact with TradingView

USAGE:
   tvctl [global options] command [command options] [arguments...]

VERSION:
   v0.0.6-625fe68 (2022-07-21T04:57:56Z)

COMMANDS:
   watchlist, wl  Commands related to watchlist features
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### Watchlist

To generate (export) a watchlist, use the following command:

```sh
# Long
tvctl watchlist generate

# Short
tvctl wl g
```

For an example of the generated output, view the [export](https://github.com/krzko/tvctl/tree/main/export) directory. The `watchlist generate` command will export all the pairs for each supported exchange and market instrument, into a file that TradingView recognises as a watchlist.

To import the file, simply go to your chart view, select the `...` in the top right-hand corner and select **Import list...**.


