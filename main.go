package main

import (
	"expected_move/prices"
	"flag"
	"github.com/joho/godotenv"
	"log"
)

var ticker = flag.String("ticker", "SPY", "Ticker to retrieve price from")
var date = flag.String("date", "", "Date to pull historical prices for")

func main() {
	loadEnv()
	flag.Parse()

	pricesController := &prices.HistoricalPriceController{
		Ticker: *ticker,
		Date:   *date,
		WriteStrategy: "console",
	}

	pricesController.GetPrices()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
