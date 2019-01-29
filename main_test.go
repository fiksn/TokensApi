/*
 * Copyright (C) 2019 Gregor Pogačnik
 */
package TokensApi

import (
	"flag"
	"fmt"
	"github.com/fiksn/TokensApi/entities"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestStuff(t *testing.T) {

	// Be able to pass -myV to go test
	var myV = flag.Int("myV", 0, "test")
	flag.Lookup("v").Value.Set(fmt.Sprint(*myV))

}
