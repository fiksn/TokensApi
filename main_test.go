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

func TestStuff(t *testing.T) {

	// Be able to pass -myV to go test
	var myV = flag.Int("myV", 0, "test")
	flag.Lookup("v").Value.Set(fmt.Sprint(*myV))

	resp, err := GetTradingPairs()

	if err != nil {
		glog.Fatalf("Unable to get trading pairs %v", err)
	}

	i := 0
	for key, pair := range resp {
		fmt.Println(key)
		fmt.Println(pair)
		i++
	}

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
