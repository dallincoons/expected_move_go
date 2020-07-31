package prices

import (
	"expected_move/alphadvantage"
	"github.com/spf13/viper"
	"log"
)

type HistoricalPriceController struct {
	Client alphadvantage.HttpClientInterface
}

func (this *HistoricalPriceController) GetAllDayPricesForAllDates(from string, to string, tickers []string, writer DisplayStrategyInterface) {

}

func (this *HistoricalPriceController) GetPrices(date string, tickers []string, writer DisplayStrategyInterface) {
	historicalPrices := newHistoricalPrices(this.Client)

	timeSeries := historicalPrices.GetAllDayPrices(tickers, date)

	for _, ts := range timeSeries {
		writer.Write(ts)
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
			HttpClient: client,
			Function:   "TIME_SERIES_DAILY",
			ApiKey:     viper.GetString("API_KEY"),
		},
	}
}
