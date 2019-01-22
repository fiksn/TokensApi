/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package entities

import (
	"encoding/json"
	"fmt"
)

type BalanceResp struct {
	Base
	Total     json.Number `json:"total,string" description:"Total amount."`
	Currency  string      `json:"currency" description:"Currency"`
	Available json.Number `json:"available,string" description:"Available amount."`
}

func (me *BalanceResp) String() string {
	return fmt.Sprintf("Balance %v %v/%v", me.Currency, me.Available, me.Total)
}
