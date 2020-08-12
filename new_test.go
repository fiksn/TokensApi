// +build go1.13

/*
 * Copyright (C) 2019-2020 Gregor Pogačnik
 */
package TokensApi

import "testing"

// Golang 1.13+ has some different logic
var _ = func() bool {
	testing.Init()
	return true
}()
