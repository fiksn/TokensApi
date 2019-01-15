/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package TokensApi

import "encoding/json"

type BalanceResp struct {
	Status    string      `json:"status" description:"Status"`
	Total     json.Number `json:"total,string" description:"Total amount."`
	Currency  string      `json:"currency" description:"Currency"`
	Available json.Number `json:"available,string" description:"Available amount."`
	Timestamp timestamp   `json:"timestamp" description:"Timestamp"`
}
