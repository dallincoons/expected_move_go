package prices

import (
	"encoding/json"
	"errors"
	"expected_move/alphadvantage"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type HttpClientInterface interface {
	Do(r *http.Request) (*http.Response, error)
}

type HistoricalPrices struct {
	Client alphadvantage.ClientInterface
}

func (this *HistoricalPrices) GetTodaysPricesFor(ticker string) (*TimeSeriesPrice, error) {
	return this.GetDayPrices(ticker, time.Now().Format("2006-01-02"))
}

func (this *HistoricalPrices) GetDayPrices(ticker string, date string) (*TimeSeriesPrice, error) {
	resp, _ := this.Client.GetDayPrices(ticker)

	var result HistoricalPriceSearchResult

	json.NewDecoder(resp.Body).Decode(&result)

	day, ok := result.TimeSeries[date]

	if !ok {
		return nil, errors.New("No prices were retrieved")
	}

	return &TimeSeriesPrice{
		Date: 	date,
		Open:   truncatePrice(day.Open),
		High:   truncatePrice(day.High),
		Low:    truncatePrice(day.Low),
		Close:  truncatePrice(day.Close),
		Volume: day.Volume,
	}, nil
}

type MostRecentDay struct {
	Open string `json:"1. open"`
	High string `json:"2. high"`
	Low string `json:"3. low"`
	Close string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type HistoricalPriceSearchResult struct {
	TimeSeries map[string]MostRecentDay `json:"Time Series (Daily)"`
}

type TimeSeriesPrice struct {
	Date string
	Open string
	High string
	Low string
	Close string
	Volume string
}

func truncatePrice(price string) string {
	numPrice, _ := strconv.ParseFloat(price, 64)

	return fmt.Sprintf("%.2f", numPrice)
}
