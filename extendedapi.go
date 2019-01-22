/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package TokensApi

import (
	"TokensApi/entities"
	"errors"
	"time"

	"github.com/golang/glog"
)

/**
 * Cancel all outstanding orders.
 */
func CancelAllOrders() error {
	orders, err := GetAllOrders()
	if err != nil {
		return err
	}

	for _, order := range orders.OpenOrders {
		fl, err := order.RemainingAmount.Float64()

		if err != nil || fl <= 0 {
			glog.Warningf("Order %v has a strange remaining amount %v", order.Id, order.RemainingAmount)
		}

		glog.V(2).Infof("Canceling order %v", order.Id)
		CancelOrder(order.Id)
	}

	return nil
}

/**
* Get all supported currency codes.
 */
func GetAllCurrencies() ([]string, error) {
	resp, err := GetTradingPairs()
	if err != nil {
		return nil, err
	}

	set := make(map[string]bool, len(resp))

	for _, pair := range resp {
		if !set[pair.BaseCurrency] {
			set[pair.BaseCurrency] = true
		}
		if !set[pair.CounterCurrency] {
			set[pair.CounterCurrency] = true
		}
	}

	ret := make([]string, len(set))
	idx := 0
	for key := range set {
		ret[idx] = key
		idx++
	}

	return ret, nil
}

type Amount float64
type Price float64

/**
 * Place an order in a type-safe manner to avoid (costly) mistakes.
 */
func PlaceOrderTyped(
	pair *entities.TradingPair,
	side entities.OrderType,
	amount Amount,
	price Price,
	takeProfitPrice *Price,
	expireDate *time.Time) (entities.PlaceOrderResp, error) {

	var resp entities.PlaceOrderResp

	if pair == nil {
		return resp, errors.New("Pair must not be nil")
	}

	minAmount, err := pair.MinAmount.Float64()
	if err != nil || float64(amount) < minAmount {
		return resp, errors.New("Amount is too low")
	}

	if takeProfitPrice != nil {
		return PlaceOrder(
			pair.BaseCurrency+pair.CounterCurrency,
			side,
			float64(amount),
			pair.AmountDecimals,
			float64(price),
			pair.PriceDecimals,
			float64(*takeProfitPrice),
			expireDate)
	} else {
		return PlaceOrder(
			pair.BaseCurrency+pair.CounterCurrency,
			side,
			float64(amount),
			pair.AmountDecimals,
			float64(price),
			pair.PriceDecimals,
			-1,
			expireDate)
	}
}

/**
* Get balances. Warning: this method is quite heavy in terms of how many calls it does to the server.
 */
func GetBalances(hideZero bool) map[string]*entities.BalanceResp {
	resp := make(map[string]*entities.BalanceResp)
	currencies, err := GetAllCurrencies()
	if err != nil {
		return nil
	}

	for _, currency := range currencies {
		result, err := GetBalance(currency)
		if err == nil {
			total, err := result.Total.Float64()
			if err != nil {
				continue
			}
			if (hideZero && total > 0) || !hideZero {
				resp[currency] = &result
			}
		}
	}

	return resp
}
