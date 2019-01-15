/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package TokensApi

import (
	"encoding/json"

	"github.com/shopspring/decimal"

	"github.com/golang/glog"
)

// volume, price
type Quotation [][2]json.Number

type OrderBook struct {
	Bids Quotation `json:"bids" description:"Bids (buy requests)"`
	Asks Quotation `json:"asks" description:"Asks (sell requests)"`
}

type OrderBookResp struct {
	Timestamp timestamp `json:"timestamp" description:"Timestamp"`
	Status    string    `json:"status" description:"Status"`
	Bids      Quotation `json:"bids" description:"Bids (buy requests)"`
	Asks      Quotation `json:"asks" description:"Asks (sell requests)"`
}

type AskOrder Quotation // also ascending

func (v AskOrder) Len() int      { return len(v) }
func (v AskOrder) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (a AskOrder) Less(i, j int) bool {
	first, err := decimal.NewFromString(a[i][1].String())
	second, err := decimal.NewFromString(a[j][1].String())

	if err != nil {
		return false
	}

	return first.Cmp(second) < 0
}

/**
* Get liquidity in counter currency (e.g., USDT).
 */
func (me *Quotation) GetLiquidity() decimal.Decimal {
	sum := decimal.New(0, 0)

	for _, val := range *me {

		onePrice, err := decimal.NewFromString(val[1].String())
		if err != nil {
			glog.Warningf("GetLiquidity fatal error %v", err)
			return sum
		}

		oneVolume, err := decimal.NewFromString(val[0].String())
		if err != nil {
			glog.Warningf("GetLiquidity fatal error %v", err)
			return sum
		}

		// sum += (oneVolume * onePrice)
		sum = sum.Add(oneVolume.Mul(onePrice))
	}

	return sum
}

func (me *Quotation) Copy() *Quotation {
	created := make(Quotation, len(*me))
	copy(created, *me)
	return &created
}

func (me *Quotation) VolumeInOtherUnit() *Quotation {
	created := make(Quotation, len(*me))
	copy(created, *me)

	for i := range created {

		onePrice, err := decimal.NewFromString(created[i][1].String())
		if err != nil {
			glog.Warningf("VolumeInOtherUnit fatal error %v", err)
			continue
		}

		oneVolume, err := decimal.NewFromString(created[i][0].String())
		if err != nil {
			glog.Warningf("VolumeInOtherUnit fatal error %v", err)
			continue
		}

		if onePrice.Cmp(decimal.New(0, 0)) < 0 {
			continue
		}

		created[i][0] = json.Number(onePrice.Mul(oneVolume).StringFixedBank(8))
	}

	return &created
}

/**
* What price would I get if I offered given amount of units of base currency (e.g., BTC)?
* Price is in counter currency (e.g., USDT), Limit is what I need to provide for order to go through
* (limit >= price in the ask scenario and limit <= price in the bid scenario).
 */
func (me *Quotation) GetPriceFor(units decimal.Decimal) (price decimal.Decimal, limit decimal.Decimal) {
	amountSum, priceAmountSum := decimal.New(0, 0), decimal.New(0, 0)
	price, limit = decimal.New(0, 0), decimal.New(0, 0)

	if units.Cmp(decimal.New(0, 0)) < 0 {
		return
	}

	for _, one := range *me {
		onePrice, err := decimal.NewFromString(one[1].String())
		oneAmount, err := decimal.NewFromString(one[0].String())
		if err != nil {
			glog.Warningf("GetPriceFor fatal error %v", err)
			continue
		}
		//priceAmountSum += (onePrice * oneAmount)

		priceAmountSum = priceAmountSum.Add(onePrice.Mul(oneAmount))
		//amountSum += oneAmount
		amountSum = amountSum.Add(oneAmount)

		// if amountSum >= units {
		if amountSum.Cmp(units) >= 0 {
			limit = onePrice
			price = priceAmountSum.Div(amountSum)
			return
		}
	}

	// Not enough liquidity

	return
}
