/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package entities

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (t timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t.Time).Unix(), 10)), nil
}

func (t *timestamp) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseInt(string(data[:]), 10, 64)
	if err != nil {
		return err
	}

	t.Time = time.Unix(i, 0)

	return nil
}
