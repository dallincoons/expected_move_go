package main

import (
	"expected_move/alphadvantage"
	"expected_move/prices"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

var ticker = flag.String("ticker", "SPY", "Ticker to retrieve price from")
var date = flag.String("date", "", "Date to pull historical prices for")

func main() {
	loadEnv()
	flag.Parse()

	var dayPrice *prices.TimeSeriesPrice
	var err error

	historicalPrices := newHistoricalPrices()

	if (*date != "") {
		dayPrice, err = historicalPrices.GetDayPrices(*ticker, *date)
	} else {
		dayPrice, err = historicalPrices.GetTodaysPricesFor(*ticker)
	}

	if (err != nil) {
		log.Fatal("Could not retrieve price")
	}

	displayTable(dayPrice)
}

func displayTable(prices *prices.TimeSeriesPrice) {
	open := truncatePrice(prices.Open)
	high := truncatePrice(prices.High)
	low := truncatePrice(prices.Low)
	close := truncatePrice(prices.Close)

	fmt.Fprintf(os.Stdout,"%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", "Date", "Open", "High", "Low", "Close", "Volume")
	fmt.Fprintf(os.Stdout, "%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", prices.Date, open, high, low, close, prices.Volume)
}

func newHistoricalPrices() (*prices.HistoricalPrices) {
	return &prices.HistoricalPrices{
		Client: &alphadvantage.Client{
			HttpClient: http.DefaultClient,
			Function:   "TIME_SERIES_DAILY",
			ApiKey:     os.Getenv("API_KEY"),
		},
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func truncatePrice(price string) string {
	numPrice, _ := strconv.ParseFloat(price, 64)

	return fmt.Sprintf("%.2f", numPrice)
}
