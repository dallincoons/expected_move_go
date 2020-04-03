package main

import (
	"expected_move/prices"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var ticker = flag.String("ticker", "SPY", "Ticker to retrieve price from")
//var date = flag.String("date", time.Now().Format("2006-01-02"), "Date to pull historical prices for")

func main() {
	loadEnv()

	historicalPrices := newHistoricalPrices()

	dayPrice, err := historicalPrices.GetTodaysPricesFor(*ticker)

	if (err != nil) {
		log.Fatal("Could not retrieve price")
	}

	displayTable(dayPrice)
}

func displayTable(prices *prices.TodaysPrices) {
	open := truncatePrice(prices.Open)
	high := truncatePrice(prices.High)
	low := truncatePrice(prices.Low)
	close := truncatePrice(prices.Close)

	fmt.Fprintf(os.Stdout,"%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", "Date", "Open", "High", "Low", "Close", "Volume")
	fmt.Fprintf(os.Stdout, "%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", prices.Date, open, high, low, close, prices.Volume)
}

func newHistoricalPrices() (*prices.HistoricalPrices) {
	return &prices.HistoricalPrices{
		Client: http.DefaultClient,
		APIKey: os.Getenv("API_KEY"),
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
