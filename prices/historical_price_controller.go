package prices

import (
	"expected_move/alphadvantage"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type HistoricalPriceController struct {
	Client alphadvantage.HttpClientInterface
}

func (this *HistoricalPriceController) GetPrices(date string, writers map[string]DisplayStrategyInterface) {
	historicalPrices := newHistoricalPrices(this.Client)

	tickers := make([]string, len(writers))
	i := 0
	for ticker := range writers {
		tickers[i] = ticker
		i++
	}

	timeSeries := historicalPrices.GetAllDayPrices(tickers, date)

	for _, t := range timeSeries {
		writer, _ := writers[t.Ticker]

		writer.Write(t)
	}
}

func (this *HistoricalPriceController) GetPrice(ticker, date string, writer DisplayStrategyInterface) {
	var dayPrice *TimeSeriesPrice
	var err error

	historicalPrices := newHistoricalPrices(this.Client)

	if date != "" {
		dayPrice, err = historicalPrices.GetDayPrices(ticker, date)
	} else {
		dayPrice, err = historicalPrices.GetTodaysPricesFor(ticker)
	}

	if err != nil {
		log.Fatal(err)
	}

	err = writer.Write(dayPrice)

	if err != nil {
		log.Fatalf(err.Error())
	}
}

func newHistoricalPrices(client alphadvantage.HttpClientInterface) *HistoricalPrices {
	return &HistoricalPrices{
		Client: &alphadvantage.Client{
			HttpClient: http.DefaultClient,
			Function:   "TIME_SERIES_DAILY",
			ApiKey:     viper.GetString("API_KEY"),
		},
	}
}
