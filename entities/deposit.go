/*
 * Copyright (C) 2019-2020 Gregor Pogačnik
 */

package entities

type DepositAddrResp struct {
	Base
	Address string `json:"address" description:"Deposit address."`
}
