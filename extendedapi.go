/*
 * Copyright (C) 2019-2020 Gregor Pogačnik
 */

package TokensApi

import (
	"errors"
	"strings"
	"time"

	"github.com/fiksn/TokensApi/entities"

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

		glog.V(6).Infof("Canceling order %v", order.Id)
		CancelOrder(order.Id)
	}

	return nil
}

/**
* Get all supported currency codes.
* See GetCurrencies()
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
		return resp, errors.New("150 Amount is too low")
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
* Get balances.
 */
func GetBalances(hideZero bool) map[string]*entities.BalanceResp {
	all, err := GetAllBalances()
	resp := make(map[string]*entities.BalanceResp)
	if err != nil {
		return nil
	}

	for currency, balance := range all.Balances {
		total, err := balance.Total.Float64()
		if err != nil {
			continue
		}
		if (hideZero && total > 0) || !hideZero {
			resp[strings.ToLower(currency)] = &entities.BalanceResp{
				Base:     all.Base,
				Currency: strings.ToLower(currency),
				Balance:  *balance,
			}
		}
	}

	return resp
}

/**
 * Get transactions. WARNING might take a while.
 */
func GetAllTransactions() (entities.TransactionResp, error) {

	// starts with 1 not 0
	page := 1
	resp := entities.TransactionResp{}
	resp.TotalPages = page
	resp.Transactions = make([]entities.Transaction, 0)

	for page < resp.TotalPages {
		temp, err := GetTransactions(page)

		resp.CurrentPage = temp.CurrentPage
		resp.TotalPages = temp.TotalPages
		resp.Status = temp.Status
		resp.Timestamp = temp.Timestamp

		resp.Transactions = append(resp.Transactions, temp.Transactions...)

		if err != nil {
			return resp, err
		}
		page++
	}

	return resp, nil
}
