package TokensApi

import (
	"fmt"
	"testing"

	"github.com/golang/glog"
	"github.com/shopspring/decimal"
)

func TestStuff(t *testing.T) {
	resp, err := GetTradingPairs()

	if err != nil {
		glog.Fatalf("Unable to get trading pairs %v", err)
	}

	i := 0
	for _, pair := range resp {
		fmt.Println(pair)
		i++
	}

	resp2, err := GetOrderBook("btcusdt")
	fmt.Printf("FOO: %v %v LIQ %v %v\n", resp2.Status, resp2.Asks, resp2.Asks.GetLiquidity(), resp2.Bids.GetLiquidity())
	price, limit := resp2.Asks.GetPriceFor(decimal.NewFromFloat(12))
	fmt.Printf("price %v, limit %v\n", price, limit)
}
