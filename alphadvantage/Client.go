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
	urlString := fmt.Sprintf("%s/query?%s", BASE_URL, this.buildQueryString(ticker))

	request, err := http.NewRequest("GET", urlString, nil)

	if (err != nil) {
		return nil, err
	}

	return this.HttpClient.Do(request)
}

func (this *Client) buildQueryString(ticker string) (string) {
	query := make(url.Values)
	query.Set("function", "TIME_SERIES_DAILY")
	query.Set("symbol", ticker)
	query.Set("apikey", this.ApiKey)
	query.Set("outputsize", "compact")

	return query.Encode()
}
