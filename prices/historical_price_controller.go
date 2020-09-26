package prices

import (
	"expected_move/alphadvantage"
	"expected_move/utility"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type HistoricalPriceController struct {
	Client alphadvantage.HttpClientInterface
	Tickers []string
}

func (pricesController *HistoricalPriceController) GetAllDayPricesForRange(from time.Time, to time.Time, writer DisplayStrategyInterface) {
	tk := utility.NewTimeKeeper()

	for _, day := range tk.GetWeekdaysSince(from, to) {
		pricesController.GetPrices(day.Format("2006-01-02"), pricesController.Tickers, writer)
	}
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
		fmt.Println(err.Error())
		return
	}

	err = writer.Write(dayPrice)

	if err != nil {
		fmt.Println(err.Error())
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
