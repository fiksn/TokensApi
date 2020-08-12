/*
 * Copyright (C) 2019-2020 Gregor Pogačnik
 */

package entities

import (
	"strconv"
	"strings"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t.Time).Unix(), 10)), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	str := string(data[:])

	// Hack to support null response
	if str == "null" {
		//fmt.Printf("Found strange date %v", str)
		t.Time = time.Unix(0, 0)
		return nil
	}

	// Hack to handle decimal point
	split := strings.Split(str, ".")
	if len(split) > 1 {
		//fmt.Printf("Found strange date %v", str)
	}
	str = split[0]

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	t.Time = time.Unix(i, 0)

	return nil
}
