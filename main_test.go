package TokensApi

import (
	"fmt"
	"testing"

	"github.com/golang/glog"
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

	fmt.Println("burek")
	fmt.Println(i)
}
