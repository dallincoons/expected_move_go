package prices

import (
	"bytes"
	"errors"
	"expected_move/alphadvantage"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetPriceForDate(t *testing.T) {
	historicalPrices := NewHistoricalPrices()

	prices, error := historicalPrices.GetDayPrices("SPY", "2020-03-26")

	if (error != nil) {
		t.Errorf("Error encountered when getting todays prices")
	}

	if (prices.Open != "249.52") {
		t.Errorf("Incorrect value for open price, %s found", prices.Open)
	}

	if (prices.High != "262.80") {
		t.Errorf("Incorrect value for high price, %s found", prices.High)
	}

	if (prices.Low != "249.05") {
		t.Errorf("Incorrect value for low price, %s found", prices.Low)
	}

	if (prices.Close != "261.20") {
		t.Errorf("Incorrect value for close price, %s found", prices.Close)
	}

	if (prices.Volume != "245530812") {
		t.Errorf("Incorrect value for volume, %s found", prices.Volume)
	}
}

func TestGetErrorMessage(t *testing.T) {
	historicalPrices := NewHistoricalPrices()

	date := "2020-04-01"

	_, error := historicalPrices.GetDayPrices("SPY", date)

	if (error == nil) {
		t.Errorf("Error expected when getting todays prices for date: %s", date)
	}
}

func NewHistoricalPrices() (*HistoricalPrices){
	return &HistoricalPrices{
		Client: &alphadvantage.Client{
			HttpClient: newFakeHttpClient(),
		},
	};
}

func newFakeHttpClient() *FakeHttpClient {
	return &FakeHttpClient{}
}

type FakeHttpClient struct {

}

func (this *FakeHttpClient) Do(r *http.Request) (*http.Response, error) {
	ticker := r.URL.Query()["symbol"][0]

	if ticker == "SPY" {
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
	"Meta Data": {
		"1. Information": "Daily Prices (open, high, low, close) and Volumes",
		"2. Symbol": "SPY",
		"3. Last Refreshed": "2020-03-26",
		"4. Output Size": "Compact",
		"5. Time Zone": "US/Eastern"
	},
	"Time Series (Daily)": {
		"2020-03-26": {
			"1. open": "249.5200",
			"2. high": "262.8000",
			"3. low": "249.0500",
			"4. close": "261.2000",
			"5. volume": "245530812"
		},
		"2020-03-25": {
			"1. open": "244.8700",
			"2. high": "256.3500",
			"3. low": "239.7500",
			"4. close": "246.7900",
			"5. volume": "297989659"
		},
		"2020-03-24": {
			"1. open": "234.4200",
			"2. high": "244.1000",
			"3. low": "233.8000",
			"4. close": "243.1500",
			"5. volume": "233038623"
		},
		"2020-03-23": {
			"1. open": "333.4200",
			"2. high": "323.1000",
			"3. low": "332.8000",
			"4. close": "343.1500",
			"5. volume": "233012345"
		}
	}
	}`)),
			StatusCode: http.StatusOK,
		}, nil
	} else if ticker == "QQQ" {
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
	"Meta Data": {
		"1. Information": "Daily Prices (open, high, low, close) and Volumes",
		"2. Symbol": "QQQ",
		"3. Last Refreshed": "2020-03-26",
		"4. Output Size": "Compact",
		"5. Time Zone": "US/Eastern"
	},
	"Time Series (Daily)": {
		"2020-03-26": {
			"1. open": "249.5200",
			"2. high": "262.8000",
			"3. low": "249.0500",
			"4. close": "261.2000",
			"5. volume": "245530812"
		},
		"2020-03-25": {
			"1. open": "244.8700",
			"2. high": "256.3500",
			"3. low": "239.7500",
			"4. close": "246.7900",
			"5. volume": "297989659"
		},
		"2020-03-24": {
			"1. open": "234.4200",
			"2. high": "244.1000",
			"3. low": "233.8000",
			"4. close": "243.1500",
			"5. volume": "233038623"
		},
		"2020-03-23": {
			"1. open": "235.4200",
			"2. high": "245.1000",
			"3. low": "234.8000",
			"4. close": "244.1500",
			"5. volume": "233038624"
		}
	}
	}`)),
			StatusCode: http.StatusOK,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("%s not found", ticker))
}
