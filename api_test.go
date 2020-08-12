/*
 * Copyright (C) 2019-2020 Gregor Pogačnik
 */
package TokensApi

import (
	"flag"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/fiksn/TokensApi/entities"
)

const pair string = "btcusdt"

func init() {
	flag.Parse()
}

func TestThatNoPublicFunctionErrors(t *testing.T) {
	if *e2e {
		return
	}

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

	statuser, err = GetCurrencies()
	if err != nil {
		t.Error("GetCurrencies failed", err)
	}
	if statuser.GetStatus() != "ok" {
		t.Error("GetCurrencies failed status", statuser.GetStatus())
	}

}

func TestThatOrderBookIsSane(t *testing.T) {

	if *e2e {
		return
	}

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

	if *e2e {
		return
	}

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
	if *e2e {
		return
	}

	resp, err := GetVotes()

	if err != nil {
		t.Error("GetVotes failed", err)
	}

	now := time.Now()

	if resp.VotingEndDate.Before(now) {
		t.Error("Voting should end in the future")
	}
}

func TestGetTransactions(t *testing.T) {

	if !initCredentials() {
		return
	}

	resp, err := GetTransactions(1)
	if err != nil {
		t.Error("GetTransactions failed", err)
	}

	if resp.CurrentPage != 1 {
		t.Error("First page should be returned instead of ", resp.CurrentPage)
	}

	if resp.TotalPages < resp.CurrentPage {
		t.Error("Current page should be within total pages range ", resp.CurrentPage, resp.TotalPages)
	}

	// fun fact page zero or 4000 should return last page
	resp, err = GetTransactions(0)
	if err != nil {
		t.Error("GetTransactions failed", err)
	}

	if resp.TotalPages != resp.CurrentPage {
		t.Error("Zero page did not return last page")
	}

	resp, err = GetTransactions(resp.TotalPages + 1)
	if err != nil {
		t.Error("GetTransactions failed", err)
	}

	if resp.TotalPages != resp.CurrentPage {
		t.Error("Last page plus one did not return last page")
	}
}

func TestThatCurrenciesAndTradingPairsAreInSync(t *testing.T) {
	if *e2e {
		return
	}

	currencies, err := GetCurrencies()
	if err != nil {
		t.Error("GetCurrencies failed", err)
	}

	currencySet := make(map[string]bool)

	for name, _ := range currencies.Currencies {
		currencySet[strings.ToLower(name)] = true
	}

	fromPairs, err := GetAllCurrencies()
	if err != nil {
		t.Error("GetAllCurrencies failed", err)
	}

	fromPairSet := make(map[string]bool)
	for _, name := range fromPairs {
		fromPairSet[strings.ToLower(name)] = true
	}

	/* Check */

	if len(currencySet) != len(fromPairSet) {
		t.Errorf("Number of currencies %d is not the same as number of currencies from trading pairs %d\n", len(currencySet), len(fromPairSet))
	}

	for name, _ := range currencySet {
		if _, ok := fromPairSet[name]; !ok {
			t.Errorf("Currency %s in GetCurrencies() but not found in trading pairs", name)
		}
	}

	for name, _ := range fromPairSet {
		if _, ok := currencySet[name]; !ok {
			t.Errorf("Currency %s from trading pair - GetAllCurrencies() but not found in GetCurrencies()", name)
		}
	}
}
