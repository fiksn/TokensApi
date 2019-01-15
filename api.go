/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package TokensApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"github.com/golang/glog"
)

const (
	TokensBaseUrl   = "https://api.tokens.net"
	TakerFeePercent = 0.02
	MakerFeePercent = 0
)

/**
* Obtain trading pairs.
 */
func GetTradingPairs() (TradingPairResp, error) {
	var resp TradingPairResp

	jsonBlob := request(TokensBaseUrl + "/public/trading-pairs/get/all/")
	glog.V(2).Infof("GetTradingPairs resp %v", string(jsonBlob))

	err := json.Unmarshal(jsonBlob, &resp)
	if err != nil {
		glog.Warningf("GetTradingPairs unable to unmarshal json blob: %v (%v)", string(jsonBlob), err)
		return resp, err
	}

	return resp, nil
}

func GetOrderBook(pair string) (OrderBookResp, error) {
	var resp OrderBookResp

	jsonBlob := request(TokensBaseUrl + fmt.Sprintf("/public/order-book/%s/", pair))
	if jsonBlob == nil {
		return resp, errors.New("No response")
	}

	glog.V(2).Infof("GetOrderBook resp %v", string(jsonBlob))

	err := json.Unmarshal(jsonBlob, &resp)

	fmt.Println("Status: ", resp)

	if err != nil {
		glog.Warningf("Unable to unmarshal json blob: %v (%v)", string(jsonBlob), err)
		return resp, err
	}

	if resp.Status != "ok" {
		return resp, errors.New(resp.Status)
	}

	sort.Sort(AskOrder(resp.Asks))
	sort.Sort(sort.Reverse(AskOrder(resp.Bids)))

	return resp, nil
}
