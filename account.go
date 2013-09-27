package gocoinbase

import "encoding/json"
import "fmt"

type balance struct {
	Amount float64 `json:",string"`
	Name   string
}

type receiveAddress struct {
  Success bool
  Address string
  CallbackUrl string
}

func (c *Coinbase) Balance() float64 {
	var b balance
	err := json.Unmarshal(c.Get("/account/balance"), &b)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return b.Amount
}

func (c *Coinbase) ReceiveAddress() string {
  var r receiveAddress
  err := json.Unmarshal(c.Get("/account/receive_address"), &r)
  if err != nil {
    fmt.Println(err)
    return ""
  } else if r.Success == false {
    fmt.Println("Getting the recieve address failed.")
    return ""
  }
  return r.Address
}
