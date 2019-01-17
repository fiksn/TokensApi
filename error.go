/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package TokensApi

type ErrorResp struct {
	Status    string    `json:"status" description:"Status"`
	Reason    string    `json:"reason" description:"Reason for the failure."`
	ErrorCode int       `json:"errorCode" description:"Code for the error"`
	Timestamp timestamp `json:"timestamp" description:"Timestamp"`
}
