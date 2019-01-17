/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package entities

import "encoding/json"

type TickerResp struct {
	Status     string      `json:"status" description:"Status"`
	Timestamp  timestamp   `json:"timestamp" description:"Timestamp"`
	Bid        json.Number `json:"bid,string" description:"Current best bid."`
	Ask        json.Number `json:"ask,string" description:"Current best bid."`
	Low        json.Number `json:"low,string" description:"Lowest value of requested interval."`
	High       json.Number `json:"high,string" description:"Highest value of requested interval."`
	Vwap       json.Number `json:"vwap,string" description:"Volume weighted average."`
	Volume     json.Number `json:"volume,string" description:"Volume in the requested interval"`
	VolumeUsdt json.Number `json:"volume_usdt,string" description:"Volume in the requested interval (in USDT)"`
}
