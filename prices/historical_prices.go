package prices

import (
	"encoding/json"
	"errors"
	"expected_move/alphadvantage"
	"fmt"
	"log"
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
	resp, err := this.Client.GetDayPrices(ticker)

	if (err != nil) {
		return nil, err
	}

	var result HistoricalPriceSearchResult

	json.NewDecoder(resp.Body).Decode(&result)

	day, ok := result.TimeSeries[date]

	if !ok {
		return nil, errors.New(fmt.Sprintf("No prices were found for date %s", date))
	}

	return &TimeSeriesPrice{
		Ticker: ticker,
		Date: 	date,
		Open:   truncatePrice(day.Open),
		High:   truncatePrice(day.High),
		Low:    truncatePrice(day.Low),
		Close:  truncatePrice(day.Close),
		Volume: day.Volume,
	}, nil
}

func (this *HistoricalPrices) GetAllDayPrices(tickers []string, date string) []*TimeSeriesPrice {
	results := make(chan *TimeSeriesPrice, len(tickers))
	var timeSeries []*TimeSeriesPrice

	for _, ticker := range tickers {
		go func(t string) {
			result, err := this.GetDayPrices(t, date)
			if err != nil {
				log.Println(fmt.Sprintf("Could not retrieve price for ticker %s on date %s", t, date))
			}
			results<-result
		}(ticker)
	}

	for range tickers {
		timeSeries = append(timeSeries, <-results)
	}

	return timeSeries
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
	Ticker string
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
