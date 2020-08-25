package expected_moves

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type HttpEMClient interface {
	getATMOptions(symbol string, from string, to string) http.Response
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

func (p ExpectedMovePuller) getExpectedMove(symbol string, from string, to string) ExpectedMove {
	response := p.HttpClient.getATMOptions(symbol, from, to)
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

func TestPullExpectedMoveForUpcomingWeek (t *testing.T) {
	puller := &ExpectedMovePuller{
		HttpClient: &FakeHttpClient{},
	}

	em := puller.getExpectedMove("spy", "2020-08-28", "2020-08-29")

	if em.HighPrice == "" {
		t.Errorf("Missing high price")
	}

	if em.HighPrice != "344.59" {
		t.Errorf("Wanted an expected move of 344.59, got %v", em.HighPrice)
	}

	if em.LowPrice != "334.69" {
		t.Errorf("Wanted an expected move of 334.69, got %v", em.LowPrice)
	}
}

type FakeHttpClient struct {

}

func (*FakeHttpClient) getATMOptions(symbol string, from string, to string) http.Response {
	return http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(`
			{
				"symbol":"SPY",
				"status":"SUCCESS",
				"underlying":null,
				"strategy":"SINGLE",
				"interval":0.0,
				"isDelayed":true,
				"isIndex":false,
				"interestRate":0.1,
				"underlyingPrice":339.64,
				"volatility":29.0,
				"daysToExpiration":0.0,
				"numberOfContracts":2,
				"callExpDateMap":{
				"2020-08-28:6":{
					"340.0":[
						{
							"putCall":"CALL",
							"symbol":"SPY_082820C340",
							"description":"SPY Aug 28 2020 340 Call (Weekly)",
							"exchangeName":"OPR",
							"bid":2.15,
							"ask":2.17,
							"last":2.18,
							"mark":2.16,
							"bidSize":25,
							"askSize":105,
							"bidAskSize":"25X105",
							"lastSize":0,
							"highPrice":2.44,
							"lowPrice":1.76,
							"openPrice":0.0,
							"closePrice":2.16,
							"totalVolume":18511,
							"tradeDate":null,
							"tradeTimeInLong":1598040897973,
							"quoteTimeInLong":1598040899947,
							"netChange":0.25,
							"volatility":12.701,
							"delta":0.473,
							"gamma":0.067,
							"theta":-0.163,
							"vega":0.187,
							"rho":0.029,
							"openInterest":14691,
							"timeValue":2.18,
							"theoreticalOptionValue":2.16,
							"theoreticalVolatility":29.0,
							"optionDeliverablesList":null,
							"strikePrice":340.0,
							"expirationDate":1598644800000,
							"daysToExpiration":6,
							"expirationType":"S",
							"lastTradingDay":1598659200000,
							"multiplier":100.0,
							"settlementType":" ",
							"deliverableNote":"",
							"isIndexOption":null,
							"percentChange":11.34,
							"markChange":0.0,
							"markPercentChange":0.0,
							"nonStandard":false,
							"inTheMoney":false,
							"mini":false
						}]}},
				"putExpDateMap":{"2020-08-28:6":{
					"340.0":[{
						"putCall":"PUT",
						"symbol":"SPY_082820P340",
						"description":"SPY Aug 28 2020 340 Put (Weekly)",
						"exchangeName":"OPR",
						"bid":2.8,
						"ask":2.82,
						"last":2.8,
						"mark":2.81,
						"bidSize":10,
						"askSize":42,
						"bidAskSize":"10X42",
						"lastSize":0,
						"highPrice":4.19,
						"lowPrice":2.63,
						"openPrice":0.0,
						"closePrice":2.81,
						"totalVolume":9204,
						"tradeDate":null,
						"tradeTimeInLong":1598040880018,
						"quoteTimeInLong":1598040900009,
						"netChange":-0.96,
						"volatility":13.706,
						"delta":-0.525,
						"gamma":0.062,
						"theta":-0.191,
						"vega":0.187,
						"rho":-0.035,
						"openInterest":3666,
						"timeValue":2.28,
						"theoreticalOptionValue":2.81,
						"theoreticalVolatility":29.0,
						"optionDeliverablesList":null,
						"strikePrice":340.0,
						"expirationDate":1598644800000,
						"daysToExpiration":6,
						"expirationType":"S",
						"lastTradingDay":1598659200000,
						"multiplier":100.0,
						"settlementType":" ",
						"deliverableNote":"",
						"isIndexOption":null,
						"percentChange":-34.28,
						"markChange":0.0,
						"markPercentChange":-0.01,
						"nonStandard":false,
						"inTheMoney":true,
						"mini":false}]}}}
		`)),
	}
}

func TestGetDateForNextFriday (t *testing.T) {
	today, _ := time.ParseInLocation("2006-01-02", "2020-08-22", time.Local)
	expected_friday, _ := time.ParseInLocation("2006-01-02", "2020-08-28", time.Local)
	finder := &WeekendFinder{FakeCalendar {
		T: today,
	}}

	friday := finder.GetNextFriday()

	if !friday.Equal(expected_friday) {
		t.Errorf("Friday not found: expected 2020-08-28, found %s", friday)
	}
}

type FakeCalendar struct {
	T time.Time
}

func (clock FakeCalendar) Today() time.Time {
	return clock.T
}
