package prices

import (
	"expected_move/alphadvantage"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type HistoricalPriceController struct {
	Ticker string
	Date string
	WriteStrategy string
}

func (this *HistoricalPriceController) GetPrices() {
	var dayPrice *TimeSeriesPrice
	var err error

	historicalPrices := newHistoricalPrices()

	if this.Date != "" {
		dayPrice, err = historicalPrices.GetDayPrices(this.Ticker, this.Date)
	} else {
		dayPrice, err = historicalPrices.GetTodaysPricesFor(this.Ticker)
	}

	if err != nil {
		log.Fatal("Could not retrieve price")
	}

	displayTable(dayPrice)
}

func displayTable(prices *TimeSeriesPrice) {
	open := truncatePrice(prices.Open)
	high := truncatePrice(prices.High)
	low := truncatePrice(prices.Low)
	close := truncatePrice(prices.Close)

	fmt.Fprintf(os.Stdout,"%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", "Date", "Open", "High", "Low", "Close", "Volume")
	fmt.Fprintf(os.Stdout, "%-10s | %-8s | %-8s | %-8s | %-8s | %-8s\n", prices.Date, open, high, low, close, prices.Volume)
}

func newHistoricalPrices() *HistoricalPrices {
	return &HistoricalPrices{
		Client: &alphadvantage.Client{
			HttpClient: http.DefaultClient,
			Function:   "TIME_SERIES_DAILY",
			ApiKey:     os.Getenv("API_KEY"),
		},
	}
}

func truncatePrice(price string) string {
	numPrice, _ := strconv.ParseFloat(price, 64)

	return fmt.Sprintf("%.2f", numPrice)
}
