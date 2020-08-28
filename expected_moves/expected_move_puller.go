package expected_moves

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const BASE_URL = "https://api.tdameritrade.com/v1/marketdata/chains"

type HttpEMClient interface {
	getATMOptions(symbol string, from string) (*http.Response, error)
}

type HttpClient struct {
	http.Client
	ApiKey string
}

func (c HttpClient) getATMOptions(symbol string, from string) (*http.Response, error) {
	urlString := fmt.Sprintf("%s?%s", BASE_URL, c.buildQueryString(symbol, from))

	request, err := http.NewRequest("GET", urlString, nil)

	if err != nil {
		return nil, err
	}

	return c.Do(request)
}

func (c HttpClient) buildQueryString(symbol string, from string) string {
	query := make(url.Values)
	query.Set("apikey", c.ApiKey)
	query.Set("symbol", symbol)
	query.Set("strikeCount", "1")
	query.Set("fromDate", from)
	query.Set("toDate", from)

	return query.Encode()
}

type ExpectedMovePuller struct {
	HttpClient HttpEMClient
}

type ExpectedMove struct {
	PeriodEndDate string
	Symbol string
	StartPrice string
	HighPrice string
	LowPrice string
}

type Call struct {
	Bid float32 `json:"bid"`
}

type Put struct {
	Bid float32 `json:"bid"`
}

type CallMap struct {
	Call map[string]map[string][]Call `json:"callExpDateMap"`
	Put map[string]map[string][]Put `json:"putExpDateMap"`
	Price float32 `json:"underlyingPrice"`
}

func (p ExpectedMovePuller) GetExpectedMove(symbol string, from string) ExpectedMove {
	response, _ := p.HttpClient.getATMOptions(symbol, from)
	defer response.Body.Close()

	var callMap CallMap

	json.NewDecoder(response.Body).Decode(&callMap)

	move := p.getCallBid(callMap) + p.getPutBid(callMap)

	return ExpectedMove{
		StartPrice: fmt.Sprintf("%.2f", callMap.Price),
		HighPrice:  fmt.Sprintf("%.2f", callMap.Price+move),
		LowPrice:   fmt.Sprintf("%.2f", callMap.Price-move),
	}
}

func (p ExpectedMovePuller) getPutBid(callMap CallMap) float32 {
	var pbid float32

	for date := range callMap.Put {
		for price := range callMap.Put[date] {
			pbid = callMap.Put[date][price][0].Bid
			break
		}
	}
	return pbid
}

func (p ExpectedMovePuller) getCallBid(callMap CallMap) float32 {
	var cbid float32

	for date := range callMap.Call {
		for price := range callMap.Call[date] {
			cbid = callMap.Call[date][price][0].Bid
			break
		}
	}
	return cbid
}
