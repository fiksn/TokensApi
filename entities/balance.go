/*
 * Copyright (C) 2019-2020 Gregor Pogačnik
 */

package entities

import (
	"encoding/json"
	"fmt"
)

type BalanceResp struct {
	Base
	Currency string `json:"currency" description:"Currency"`
	Balance
}

type Balance struct {
	Total     json.Number `json:"total,string" description:"Total amount."`
	Available json.Number `json:"available,string" description:"Available amount."`
}

type Balances map[string]*Balance

type AllBalanceResp struct {
	Base
	Balances `json:"balances" description:"Balances"`
}

func (me BalanceResp) String() string {
	return fmt.Sprintf("Balance %v %v/%v", me.Currency, me.Available, me.Total)
}
