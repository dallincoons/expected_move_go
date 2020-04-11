package main

import (
	"expected_move/prices"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ticker = flag.String("ticker", "SPY", "Ticker to retrieve price from")
var date = flag.String("date", "", "Date to pull historical prices for")
var file = flag.String("file", "", "file to write prices")

func main() {
	loadEnv()
	flag.Parse()

	openFile := getCsvWriter()

	pricesController := &prices.HistoricalPriceController{
		Ticker: *ticker,
		Date:   *date,
		WriteStrategy: &prices.WriteCSV{
			Writer: openFile,
			DisplayToUser: os.Stdout,
		},
	}

	pricesController.GetPrices()
}

func getCsvWriter() *os.File {

	file, err := os.OpenFile(*file, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatalf("failed opening fle, %s", err)
	}

	return file
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
