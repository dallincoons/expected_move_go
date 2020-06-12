package prices

import (
	"log"
	"testing"
)

var output = map[string]*TimeSeriesPrice{}

func TestGetMultiplePricesForDate(t *testing.T) {
	historical_price_controller := &HistoricalPriceController{
		Client: &FakeHttpClient{},
	}

	writers := map[string]DisplayStrategyInterface{
		"SPY": &FakeDisplay{},
		"QQQ": &FakeDisplay{},
	}

	 historical_price_controller.GetPrices("2020-03-26", writers)

	_, spy_found := output["SPY"]

	if (!spy_found) {
		log.Fatal("SPY stock not found when retrieving multiple prices")
	}

	_, qqq_found := output["QQQ"]

	if (!qqq_found) {
		log.Fatal("QQQ stock not found when retrieving multiple prices")
	}
}

type FakeDisplay struct {
}

func (this *FakeDisplay) Write(prices *TimeSeriesPrice) (error) {
	output[prices.Ticker] = prices

	return nil
}

//type FakeHttpClient struct {
//	TimesRan int
//}

//func (this *FakeHttpClient) Do(r *http.Request) (*http.Response, error) {
//	fmt.Println(r)
//
//
//	return &http.Response{
//		Body: ioutil.NopCloser(bytes.NewBufferString(`{
//	"Meta Data": {
//		"1. Information": "Daily Prices (open, high, low, close) and Volumes",
//		"2. Symbol": "SPY",
//		"3. Last Refreshed": "2020-03-26",
//		"4. Output Size": "Compact",
//		"5. Time Zone": "US/Eastern"
//	},
//	"Time Series (Daily)": {
//		"2020-03-26": {
//			"1. open": "249.5200",
//			"2. high": "262.8000",
//			"3. low": "249.0500",
//			"4. close": "261.2000",
//			"5. volume": "245530812"
//		},
//		"2020-03-25": {
//			"1. open": "244.8700",
//			"2. high": "256.3500",
//			"3. low": "239.7500",
//			"4. close": "246.7900",
//			"5. volume": "297989659"
//		},
//		"2020-03-24": {
//			"1. open": "234.4200",
//			"2. high": "244.1000",
//			"3. low": "233.8000",
//			"4. close": "243.1500",
//			"5. volume": "233038623"
//		}
//	}
//	}`)),
//		StatusCode: http.StatusOK,
//	}, nil
//}
