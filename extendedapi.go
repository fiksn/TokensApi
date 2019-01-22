/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package TokensApi

import "github.com/golang/glog"

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
