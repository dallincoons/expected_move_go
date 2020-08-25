package expected_moves

import (
	"encoding/json"
	"fmt"
	"net/http"
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
