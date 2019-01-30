/*
 * Copyright (C) 2019 Gregor Pogačnik
 */
package TokensApi

import (
	"flag"
	"fmt"
	"testing"

	"github.com/fiksn/TokensApi/entities"
)

func TestThatNoPublicFunctionErrors(t *testing.T) {
	// Be able to pass -myV to go test
	var myV = flag.Int("myV", 0, "test")
	flag.Lookup("v").Value.Set(fmt.Sprint(*myV))

	var statuser entities.Statuser

	const pair string = "btcusdt"

	_, err := GetTradingPairs()
	if err != nil {
		t.Error("GetTradingPairs failed", err)
	}
	statuser, err = GetOrderBook(pair)
	if statuser.GetStatus() != "ok" {
		t.Error("GetOrderBook failed status", statuser.GetStatus())
	}
	if err != nil {
		t.Error("GetOrderBook failed", err)
	}
	statuser, err = GetTicker(pair, DAY)
	if err != nil {
		t.Error("GetTicker for DAY failed", err)
	}
	if statuser.GetStatus() != "ok" {
		t.Error("GetTicker for DAY failed status", statuser.GetStatus())
	}
	statuser, err = GetTicker(pair, HOUR)
	if err != nil {
		t.Error("GetTicker for HOUR failed", err)
	}
	if statuser.GetStatus() != "ok" {
		t.Error("GetTicker for HOUR failed status", statuser.GetStatus())
	}
	statuser, err = GetTrades(pair, HOUR)
	if err != nil {
		t.Error("GetTrades for HOUR failed", err)
	}
	if statuser.GetStatus() != "ok" {
		t.Error("GetTrades for HOUR failed status", statuser.GetStatus())
	}
	statuser, err = GetTrades(pair, DAY)
	if err != nil {
		t.Error("GetTrades for DAY failed", err)
	}
	if statuser.GetStatus() != "ok" {
		t.Error("GetTrades for DAY failed status", statuser.GetStatus())
	}
	/*
		statuser, err = GetTrades(pair, MINUTE)
		if err != nil {
			t.Error("GetTrades for MINUTE failed", err)
		}
		if statuser.GetStatus() != "ok" {
			t.Error("GetTrades for MINUTE failed status", statuser.GetStatus())
		}
	*/
	statuser, err = GetVotes()
	if err != nil {
		t.Error("GetVotes failed", err)
	}
	if statuser.GetStatus() != "ok" {
		t.Error("GetVotes failed status", statuser.GetStatus())
	}
}
