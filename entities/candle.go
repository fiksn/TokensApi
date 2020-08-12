/*
 * Copyright (C) 2019-2020 Gregor Pogačnik
 */

package entities

import (
	"encoding/json"
	"fmt"
)

func (n *CandleItem) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.Time, &n.Open, &n.High, &n.Low, &n.Close, &n.Volume}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in CandleItem: %d != %d", g, e)
	}
	return nil
}

type CandlesResp struct {
	Base
	MoreAvailable bool         `json:"more_available" description:"Is more data available?"`
	Data          []CandleItem `json:"data" description:"Actual data"`
}

type CandleItem struct {
	Time   Timestamp
	Open   json.Number
	High   json.Number
	Low    json.Number
	Close  json.Number
	Volume json.Number
}
