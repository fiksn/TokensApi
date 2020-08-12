// +build go1.12

/*
 * Copyright (C) 2019-2020 Gregor Pogačnik
 */
package TokensApi

import "flag"

// Golang 1.13+ has some different logic
func init() {
	flag.Parse()
}
