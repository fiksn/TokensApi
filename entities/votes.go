/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package entities

type VotesResp struct {
	Base
	VotingEndDate Timestamp       `json:"votingEndDate" description:"When voting ends"`
	WinnerList    []WinnerElement `json:"winnerList" description:"Winner list"`
	ActiveList    []ActiveElement `json:"activeList" description:"Active voting list"`
}

type PublishStatus int

const (
	// currency will be added to exchange at unknown date
	AddUnkown = 0
	// currency will be added to exchange at known date (AddToMarketDate)
	AddKnown = 1
	// currency has been added and trading is enabled
	AddDone = 1
)

type CurrencyInfo struct {
	Currency     string `json:"currency" description:"Currency code like XLM"`
	CurrencyName string `json:"currencyName" description:"Currency name like Stellar"`
}

type WinnerElement struct {
	CurrencyInfo
	TotalVotes      int           `json:"totalVotes" description:"Number of votes"`
	WinDate         Timestamp     `json:"winDate" description:"When the currency has won"`
	AddToMarketDate Timestamp     `json:"addToMarketDate" description:"When the currency will be added to market"`
	PublishStatus   PublishStatus `json:"publishStatus" description:"Is adding currency to market date known?"`
}

type ActiveElement struct {
	CurrencyInfo
	Votes int `json:"votes" description:"Number of votes so far"`
}
