/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package TokensApi

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
)

const (
	TokensBaseUrl = "https://api.tokens.net"
	TakerFee      = 0.02
	MakerFee      = 0
)

/**
* Obtain trading pairs.
 */
func GetTradingPairs() (TradingPairResp, error) {
	var resp TradingPairResp

	jsonBlob := request(TokensBaseUrl + "/public/trading-pairs/get/all/")
	glog.V(2).Infof("GetTradingPairs resp %v", string(jsonBlob))

	fmt.Println(string(jsonBlob))

	err := json.Unmarshal(jsonBlob, &resp)
	if err != nil {
		glog.Warningf("GetTradingPairs unable to unmarshal json blob: %v (%v)", string(jsonBlob), err)
		return resp, err
	}

	return resp, nil
}
