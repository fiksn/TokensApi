/*
 * Copyright (C) 2019-2020 Gregor Pogačnik
 */

package entities

import "encoding/json"

type Currency struct {
	Name          string      `json:"name" description:"Name of the currency."`
	NamePlural    string      `json:"namePlural" description:"Plural name of the currency."`
	Decimals      int         `json:"decimals,int" description:"Number of decimals."`
	WithdrawalFee json.Number `json:"withdrawalFee,string" description:"Fee in this currency to whitdraw."`
	MinimalOrder  json.Number `json:"minimalOrder,string" description:"Minimal order."`
	CanWithdraw   bool        `json:"canWithDraw" description:"Can the currency currently be withdrawn?"`
	UsdtValue     json.Number `json:"usdtValue,string" description:"Value of this currency in USDT."`
	BtcValue      json.Number `json:"btcValue,string" description:"Value of this currency in BTC."`
	EthValue      json.Number `json:"ethValue,string" description:"Value of this currency in ETH."`
}

type CurrencyResp struct {
	Base
	Currencies map[string]Currency `json:"currencies" description:"The actual currencies."`
}
