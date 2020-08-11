/*
 * Copyright (C) 2019-2020 Gregor Pogačnik
 */

package entities

type ErrorResp struct {
	Base
	Reason    string `json:"reason" description:"Reason for the failure."`
	ErrorCode int    `json:"errorCode" description:"Code for the error"`
}
