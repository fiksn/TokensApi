/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package TokensApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"TokensApi/entities"

	"github.com/golang/glog"
)

const (
	TokensBaseUrl   = "https://api.tokens.net"
	TakerFeePercent = 0.02
	MakerFeePercent = 0
)

type Interval int

const (
	DAY = iota
	HOUR
	MINUTE
)

/**
* List all existing pairs.
 */
func GetTradingPairs() (entities.TradingPairResp, error) {
	var resp entities.TradingPairResp

	jsonBlob := request(TokensBaseUrl + "/public/trading-pairs/get/all/")
	glog.V(2).Infof("GetTradingPairs resp %v", string(jsonBlob))

	err := json.Unmarshal(jsonBlob, &resp)
	if err != nil {
		glog.Warningf("GetTradingPairs unable to unmarshal json blob: %v (%v)", string(jsonBlob), err)
		return resp, err
	}

	return resp, nil
}

/**
* Get all supported currency codes.
 */
func GetAllCurrencies() ([]string, error) {
	resp, err := GetTradingPairs()
	if err != nil {
		return nil, err
	}

	set := make(map[string]bool, len(resp))

	for _, pair := range resp {
		if !set[pair.BaseCurrency] {
			set[pair.BaseCurrency] = true
		}
		if !set[pair.CounterCurrency] {
			set[pair.CounterCurrency] = true
		}
	}

	ret := make([]string, len(set))
	idx := 0
	for key := range set {
		ret[idx] = key
		idx++
	}

	return ret, nil
}

/**
* Get order book.
 */
func GetOrderBook(pair string) (entities.OrderBookResp, error) {
	var resp entities.OrderBookResp

	jsonBlob := request(TokensBaseUrl + fmt.Sprintf("/public/order-book/%s/", pair))
	if jsonBlob == nil {
		return resp, errors.New("No response")
	}

	glog.V(2).Infof("GetOrderBook resp %v", string(jsonBlob))

	err := json.Unmarshal(jsonBlob, &resp)

	if err != nil {
		glog.Warningf("Unable to unmarshal json blob: %v (%v)", string(jsonBlob), err)
		return resp, err
	}

	if resp.Status != "ok" {
		return resp, errors.New(resp.Status)
	}

	sort.Sort(entities.AskOrder(resp.Asks))
	sort.Sort(sort.Reverse(entities.AskOrder(resp.Bids)))

	return resp, nil
}

/**
* Get balance.
 */
func GetBalance(currency string) (entities.BalanceResp, error) {
	var resp entities.BalanceResp

	jsonBlob := requestAuth(TokensBaseUrl + fmt.Sprintf("/private/balance/%s/", currency))
	if jsonBlob == nil {
		return resp, errors.New("No response")
	}

	glog.V(2).Infof("GetBalance resp %v", string(jsonBlob))

	err := json.Unmarshal(jsonBlob, &resp)

	if err != nil {
		glog.Warningf("Unable to unmarshal json blob: %v (%v)", string(jsonBlob), err)
		return resp, err
	}

	if resp.Status != "ok" {
		return resp, errors.New(resp.Status)
	}

	return resp, nil
}

/**
* Get ticker for last day or hour.
 */
func GetTicker(pair string, interval Interval) (entities.TickerResp, error) {
	var (
		resp entities.TickerResp
		url  string
	)

	switch interval {
	case HOUR:
		url = fmt.Sprintf("/public/ticker/hour/%s/", pair)
	case DAY:
		url = fmt.Sprintf("/public/ticker/%s/", pair)
	default:
		return resp, errors.New("Illegal interval specified")
	}

	jsonBlob := request(TokensBaseUrl + url)
	if jsonBlob == nil {
		return resp, errors.New("No response")
	}

	glog.V(2).Infof("GetTicker resp %v", string(jsonBlob))

	err := json.Unmarshal(jsonBlob, &resp)

	if err != nil {
		glog.Warningf("Unable to unmarshal json blob: %v (%v)", string(jsonBlob), err)
		return resp, err
	}

	if resp.Status != "ok" {
		return resp, errors.New(resp.Status)
	}

	return resp, nil
}

/**
* List trades, which occured in last minute, hour or day.
 */
func GetTrades(pair string, interval Interval) (entities.TradesResp, error) {
	var (
		resp entities.TradesResp
		url  string
	)

	switch interval {
	case HOUR:
		url = fmt.Sprintf("/public/trades/hour/%s/", pair)
	case DAY:
		url = fmt.Sprintf("/public/trades/day/%s/", pair)
	case MINUTE:
		url = fmt.Sprintf("/public/trades/minute/%s/", pair)
	default:
		return resp, errors.New("Illegal interval specified")
	}

	jsonBlob := request(TokensBaseUrl + url)
	if jsonBlob == nil {
		return resp, errors.New("No response")
	}

	glog.V(2).Infof("GetTrades resp %v", string(jsonBlob))

	err := json.Unmarshal(jsonBlob, &resp)

	if err != nil {
		glog.Warningf("Unable to unmarshal json blob: %v (%v)", string(jsonBlob), err)
		return resp, err
	}

	if resp.Status != "ok" {
		return resp, errors.New(resp.Status)
	}

	return resp, nil
}

/**
 * List all currencies participating in voting and number of votes for each currency.
 */
func GetVotes() (entities.VotesResp, error) {
	var resp entities.VotesResp

	jsonBlob := request(TokensBaseUrl + "/public/voting/get/all/")
	if jsonBlob == nil {
		return resp, errors.New("No response")
	}

	glog.V(2).Infof("GetVotes resp %v", string(jsonBlob))

	err := json.Unmarshal(jsonBlob, &resp)

	if err != nil {
		glog.Warningf("Unable to unmarshal json blob: %v (%v)", string(jsonBlob), err)
		return resp, err
	}

	if resp.Status != "ok" {
		return resp, errors.New(resp.Status)
	}

	return resp, nil
}
