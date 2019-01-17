/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package entities

import "encoding/json"

type TradesResp struct {
	Status    string    `json:"status" description:"Status"`
	Timestamp timestamp `json:"timestamp" description:"Timestamp"`
	Trades    []Trade   `json:"trades" description:"Timestamp"`
}

type Trade struct {
	Id       int64       `json:"id" description:"ID"`
	Datetime timestamp   `json:"datetime" description:"Timestamp of trade"`
	Price    json.Number `json:"price,string" description:"Price of trade."`
	Amount   json.Number `json:"amount,string" description:"Amount of trade."`
	Type     string      `json:"type" description:"Either buy or sell"`
}
