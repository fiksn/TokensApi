/*
 * Copyright (C) 2019-2020 Gregor Pogačnik
 */

package entities

import "encoding/json"

type TradingPair struct {
	PriceDecimals   int         `json:"priceDecimals" description:"Decimals for price"`
	AmountDecimals  int         `json:"amountDecimals" description:"Decimals for amount"`
	MinAmount       json.Number `json:"minAmount" description:"Minimum amount of base currency."`
	BaseCurrency    string      `json:"baseCurrency" description:"Base currency."`
	CounterCurrency string      `json:"counterCurrency" description:"Counter currency."`
	Title           string      `json:"title" description:"Title."`
}

type TradingPairResp map[string]TradingPair

func (me TradingPair) String() string {
	return "Trading pair " + me.Title
}
