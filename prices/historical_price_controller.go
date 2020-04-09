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
	WriteStrategy DisplayStrategyInterface
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

	this.WriteStrategy.Write(dayPrice)
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
