/*
 * Copyright (C) 2019 Gregor Pogačnik
 */
package TokensApi

import (
	"math"
	"testing"
	"time"

	"github.com/fiksn/TokensApi/entities"
)

const pair string = "btcusdt"

func TestThatNoPublicFunctionErrors(t *testing.T) {
	var statuser entities.Statuser

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
	statuser, err = GetTicker(pair, Day)
	if err != nil {
		t.Error("GetTicker for Day failed", err)
	}
	if statuser.GetStatus() != "ok" {
		t.Error("GetTicker for Day failed status", statuser.GetStatus())
	}
	statuser, err = GetTicker(pair, Hour)
	if err != nil {
		t.Error("GetTicker for Hour failed", err)
	}
	if statuser.GetStatus() != "ok" {
		t.Error("GetTicker for Hour failed status", statuser.GetStatus())
	}
	statuser, err = GetTrades(pair, Hour)
	if err != nil {
		t.Error("GetTrades for Hour failed", err)
	}
	if statuser.GetStatus() != "ok" {
		t.Error("GetTrades for Hour failed status", statuser.GetStatus())
	}
	statuser, err = GetTrades(pair, Day)
	if err != nil {
		t.Error("GetTrades for Day failed", err)
	}
	if statuser.GetStatus() != "ok" {
		t.Error("GetTrades for Day failed status", statuser.GetStatus())
	}
	statuser, err = GetVotes()
	if err != nil {
		t.Error("GetVotes failed", err)
	}
	if statuser.GetStatus() != "ok" {
		t.Error("GetVotes failed status", statuser.GetStatus())
	}
}

func TestThatOrderBookIsSane(t *testing.T) {

	resp, err := GetOrderBook(pair)

	if err != nil {
		t.Error("GetOrderBook failed")
		return
	}

	now := time.Now()
	diff := math.Abs(now.Sub(resp.Timestamp.Time).Minutes())

	if diff >= 1 {
		t.Error("Timestamp from order book is not accurate (diff more than 1 minute)", now, resp.Timestamp.Time)
	}

	if len(resp.Bids) < 2 || len(resp.Asks) < 2 {
		t.Error("No liquidity")
		return
	}

	if resp.Bids[0][entities.Price] > resp.Asks[0][entities.Price] {
		t.Error("Ask price should be higher or equal that bid")
	}

	if resp.Asks[0][entities.Price] > resp.Asks[1][entities.Price] {
		t.Error("Ask prices should be ordered ascending")
	}

	if resp.Bids[0][entities.Price] < resp.Bids[1][entities.Price] {
		t.Error("Bid prices should be ordered descending")
	}
}

func TestThatTradingPairsAreSane(t *testing.T) {

	resp, err := GetTradingPairs()

	val, err := resp[pair].MinAmount.Float64()
	if err != nil || val <= 0 || val > 1000000 {
		t.Error("Trading pair should have sane MinAmount", resp[pair].MinAmount)
	}
	if resp[pair].BaseCurrency+resp[pair].CounterCurrency != pair {
		t.Error("Trading pair base and counter currency are not correct", pair, resp[pair].BaseCurrency, resp[pair].CounterCurrency)
	}
}

func TestThatVotingIsSane(t *testing.T) {
	resp, err := GetVotes()

	if err != nil {
		t.Error("GetVotes failed", err)
	}

	now := time.Now()

	if resp.VotingEndDate.Before(now) {
		t.Error("Voting should end in the future")
	}
}
