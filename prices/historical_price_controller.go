package prices

import (
	"expected_move/alphadvantage"
	"github.com/spf13/viper"
	"log"
	"net/http"
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
		log.Fatal(err)
	}

	err = this.WriteStrategy.Write(dayPrice)

	if err != nil {
		log.Fatalf(err.Error())
	}
}

func newHistoricalPrices() *HistoricalPrices {
	return &HistoricalPrices{
		Client: &alphadvantage.Client{
			HttpClient: http.DefaultClient,
			Function:   "TIME_SERIES_DAILY",
			ApiKey:     viper.GetString("API_KEY"),
		},
	}
}
