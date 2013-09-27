package gocoinbase

import "encoding/json"
import "fmt"

type balance struct {
	Amount float64 `json:",string"`
	Name   string
}

func Balance() float64 {
	var b balance
	err := json.Unmarshal(Get("/balance"), &b)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return b.Amount
}
