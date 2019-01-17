/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package entities

import "encoding/json"

type BalanceResp struct {
	Base
	Total     json.Number `json:"total,string" description:"Total amount."`
	Currency  string      `json:"currency" description:"Currency"`
	Available json.Number `json:"available,string" description:"Available amount."`
}
