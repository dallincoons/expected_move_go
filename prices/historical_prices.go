package prices

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

type HttpClientInterface interface {
	Do(r *http.Request) (*http.Response, error)
}

type HistoricalPrices struct {
	Client HttpClientInterface
	APIKey string
}

func (this *HistoricalPrices) GetTodaysPricesFor(ticker string) (*TodaysPrices, error) {
	return this.GetDayPrices(ticker, time.Now().Format("2006-01-02"))
}

func (this *HistoricalPrices) GetDayPrices(ticker string, date string) (*TodaysPrices, error) {
	query := make(url.Values)
	query.Set("function", "TIME_SERIES_DAILY")
	query.Set("symbol", ticker)
	query.Set("apikey", this.APIKey)
	query.Set("outputsize", "compact")

	request, _ := http.NewRequest("GET", "https://www.alphavantage.co/query?"+query.Encode(), nil)
	resp, _ := this.Client.Do(request)

	var result HistoricalPriceSearchResult

	json.NewDecoder(resp.Body).Decode(&result)

	day, ok := result.TimeSeries[date]

	if !ok {
		return nil, errors.New("No prices were retrieved")
	}

	return &TodaysPrices{
		Date: 	date,
		Open:   day.Open,
		High:   day.High,
		Low:    day.Low,
		Close:  day.Close,
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

type TodaysPrices struct {
	Date string
	Open string
	High string
	Low string
	Close string
	Volume string
}
