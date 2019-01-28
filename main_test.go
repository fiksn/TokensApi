/*
 * Copyright (C) 2019 Gregor Pogačnik
 */
package TokensApi

import (
	"TokensApi/entities"
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/golang/glog"
	"github.com/shopspring/decimal"
)

type Currencies map[string]*CurrencyVertex

type CurrencyVertex struct {
	Symbol              string
	ConnectedCurrencies map[string]*CurrencyEdge
}

type CurrencyEdge struct {
	Data *CurrencyData
	Src  *CurrencyVertex
	Dst  *CurrencyVertex
	Fwd  bool
}

type CurrencyData struct {
	AdditionalData interface{}
	MinAmount      float64
}

func (me *CurrencyEdge) GetTradingPairName() string {
	if me.Fwd {
		return me.Src.Symbol + me.Dst.Symbol
	} else {
		return me.Dst.Symbol + me.Src.Symbol
	}
}

func NewCurrencies() Currencies {
	pairs, err := GetTradingPairs()
	resp := make(map[string]*CurrencyVertex)

	if err != nil {
		glog.Fatalf("Unable to get trading pairs %v", err)
	}

	for _, pair := range pairs {
		one := pair.CounterCurrency

		if resp[one] == nil {
			emptyOne := make(map[string]*CurrencyEdge)
			resp[one] = &CurrencyVertex{
				Symbol:              one,
				ConnectedCurrencies: emptyOne,
			}
		}

		other := pair.BaseCurrency

		if resp[other] == nil {
			emptyTwo := make(map[string]*CurrencyEdge)
			resp[other] = &CurrencyVertex{
				Symbol:              other,
				ConnectedCurrencies: emptyTwo,
			}
		}

		f, err := pair.MinAmount.Float64()
		if err != nil {
			glog.Warningf("Could not convert minAmount %v (%v)", pair.MinAmount, err)
			continue
		}
		if resp[one].ConnectedCurrencies[other] == nil {
			resp[one].ConnectedCurrencies[other] = &CurrencyEdge{
				Data: &CurrencyData{AdditionalData: pair, MinAmount: f},
				Src:  resp[one],
				Dst:  resp[other],
				Fwd:  true,
			}
		}

		if resp[other].ConnectedCurrencies[one] == nil {
			resp[other].ConnectedCurrencies[one] = &CurrencyEdge{
				Data: &CurrencyData{AdditionalData: pair, MinAmount: f},
				Src:  resp[other],
				Dst:  resp[one],
				Fwd:  false,
			}
		}
	}

	return resp
}

type DftResult struct {
	Results [][]string
}

func NewDftResult() *DftResult {
	me := &DftResult{Results: make([][]string, 0)}
	return me
}

func (me *DftResult) AddResult(result []string) {
	me.Results = append(me.Results, result)
}

func dft(currencies Currencies, path []string, limit int, result *DftResult) {
	if len(path) > limit || len(path) <= 0 {
		return
	}

	last := path[len(path)-1]
	currency := currencies[last]

	if currency.Symbol == path[0] && len(path) == limit {

		fmt.Println("Found path")
		for i := 0; i < len(path); i++ {
			fmt.Printf(" -> %v", path[i])
		}
		fmt.Println()

		result.AddResult(path)
	}

	// Explore in depth
OUTER:
	for _, con := range currency.ConnectedCurrencies {
		dst := con.Dst.Symbol

		newPath := make([]string, len(path)+1)
		for i := 0; i < len(path); i++ {
			if i > 0 && path[i] == dst {
				// There would be a (shorter) cycle
				continue OUTER
			}
			newPath[i] = path[i]
		}
		newPath[len(path)] = dst

		dft(currencies, newPath, limit, result)
	}
}

func (currencies Currencies) GetCycles(src string, length int) (resp [][]string) {
	one := currencies[src]

	path := make([]string, 1)
	path[0] = one.Symbol
	result := NewDftResult()

	dft(currencies, path, length, result)

	fmt.Println(len(result.Results))
	return result.Results
}

func (currencies Currencies) GetInterestingPairs(pairs [][]string) []string {
	temp := make(map[string]bool)

OUTER:
	for _, one := range pairs {
		for i := 0; i < len(one)-1; i++ {
			currency := currencies[one[i]]

			if currency == nil {
				glog.Warningf("Invalid currency %v", one[i])
				continue OUTER
			}

			edge := currency.ConnectedCurrencies[one[i+1]]

			if edge == nil {
				glog.Warningf("Invalid currency %v", one[i+1])
				continue OUTER
			}

			if _, ok := temp[edge.GetTradingPairName()]; !ok {
				temp[edge.GetTradingPairName()] = true
			}
		}
	}

	resp := make([]string, 0, len(temp))
	for k := range temp {
		resp = append(resp, k)
	}

	return resp
}

func TestStuff(t *testing.T) {

	// Be able to pass -myV to go test
	var myV = flag.Int("myV", 0, "test")
	flag.Lookup("v").Value.Set(fmt.Sprint(*myV))

	resp := NewCurrencies()
	for _, v := range resp {
		fmt.Printf("Currency %v\n", v.Symbol)
		for s, x := range v.ConnectedCurrencies {
			fmt.Printf("-> %v (%v)\n", s, x.GetTradingPairName())

		}
	}

	foo := resp.GetCycles("eurs", 4)

	fmt.Println(foo)

	fmt.Println(resp.GetInterestingPairs(foo))

	//resp.GetCycles("eurs", 5)

	Init("./credentials")

	balance := GetBalances(true)
	for _, b := range balance {
		fmt.Println(b)
	}

	resp2, err := GetOrderBook("btcusdt")
	fmt.Printf("FOO: %v %v LIQ %v %v\n", resp2.Status, resp2.Asks, resp2.Asks.GetLiquidity(), resp2.Bids.GetLiquidity())
	price, limit := resp2.Bids.GetPriceFor(decimal.NewFromFloat(1))
	fmt.Printf("price %v, limit %v\n", price, limit)

	resp4, err := GetAllCurrencies()
	fmt.Println(resp4)

	fmt.Println("???")

	return

	expiryTime := time.Now().Add(time.Second * 5)
	respY, err := PlaceOrder("eursusdt", entities.BUY,
		2, 2,
		1, 5,
		3,
		&expiryTime)

	fmt.Println(respY)
	fmt.Println("Error for place ")
	fmt.Println(err)

	balance = GetBalances(true)
	for _, b := range balance {
		fmt.Println(b)
	}

	respX, err := GetAllOrders()
	fmt.Println(respX)

	time.Sleep(10 * time.Second)

	err = CancelAllOrders()
	fmt.Println(err)

	respX, err = GetAllOrders()
	fmt.Println(respX)

	fmt.Println("???")

}
