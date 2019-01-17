/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package entities

type Base struct {
	Status    string    `json:"status" description:"Status"`
	Timestamp timestamp `json:"timestamp" description:"Timestamp"`
}
