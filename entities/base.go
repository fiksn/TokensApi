/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package entities

type Statuser interface {
	GetStatus() string
	SetStatus(string)
}

type Base struct {
	Status    string    `json:"status" description:"Status"`
	Timestamp timestamp `json:"timestamp" description:"Timestamp"`
}

func (me *Base) GetStatus() string {
	return me.Status
}

func (me *Base) SetStatus(status string) {
	me.Status = status
}
