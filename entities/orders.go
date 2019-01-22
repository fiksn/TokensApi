/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package entities

import (
	"encoding/json"

	"github.com/satori/go.uuid"
)

type OrderType string

const (
	BUY  OrderType = "buy"
	SELL OrderType = "sell"
)

type OrderDetails string

const (
	OPEN     OrderDetails = "Open"
	FILLED   OrderDetails = "Filled"
	CANCELED OrderDetails = "Canceled"
	EXPIRED  OrderDetails = "Expired"
)

type OrdersResp struct {
	Base
	// All orders here are Open
	OpenOrders []OpenOrder `json:"openOrders" description:"Open orders."`
}

type OpenOrder struct {
	Id              uuid.UUID   `json:"id" description:"ID"`
	Created         timestamp   `json:"created" description:"Timestamp of order"`
	Type            OrderType   `json:"type" description:"Either buy or sell"`
	Price           json.Number `json:"price,string" description:"Price of order."`
	TakeProfitPrice json.Number `json:"takeeprofit,string" description:"Price to take profit."`
	Amount          json.Number `json:"amount,string" description:"Amount of order."`
	RemainingAmount json.Number `json:"remainingAmount,string" description:"Remaining amount of order."`
	CurrencyPair    string      `json:"currencyPair" description:"Currency pair of order."`
}

type OrderDetailsResp struct {
	Base
	// This order might have a different status than "Open", check OrderStatus
	OpenOrder
	Trades []ExtendedTrade `json:"trades" description:"Trades."`

	OrderStatus OrderDetails `json:"orderStatus" description:"Either buy or sell"`
}

type ExtendedTrade struct {
	Trade
	Value json.Number `json:"value,string" description:"Value is price * amount"`
}

type PlaceOrderResp struct {
	Base
	OrderId uuid.UUID `json:"orderId" description:"ID"`
}
