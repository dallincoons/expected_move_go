package alphadvantage

import (
	"fmt"
	"net/http"
	"net/url"
)

const BASE_URL = "https://www.alphavantage.co";

type HttpClientInterface interface {
	Do(r *http.Request) (*http.Response, error)
}

type ClientInterface interface {
	GetDayPrices(ticker string) (*http.Response, error)
}

type Client struct {
	HttpClient HttpClientInterface
	Function string
	ApiKey string
	OutputSize string `default:"compact"`
}

func (this *Client) GetDayPrices(ticker string) (*http.Response, error) {
	fmt.Println(this.OutputSize)
	query := make(url.Values)
	query.Set("function", "TIME_SERIES_DAILY")
	query.Set("symbol", ticker)
	query.Set("apikey", this.ApiKey)
	query.Set("outputsize", "compact")

	urlString := fmt.Sprintf("%s/query?%s", BASE_URL, query.Encode())

	request, _ := http.NewRequest("GET", urlString, nil)

	return this.HttpClient.Do(request)
}
