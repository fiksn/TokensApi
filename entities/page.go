/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package entities

type Page struct {
	Base
	CurrentPage int `json:"page" description:"Current page"`
	TotalPages  int `json:"pages" description:"Total pages"`
}
