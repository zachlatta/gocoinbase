package gocoinbase

import "encoding/json"

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
  panicOnError(err)
	return b.Amount
}

func (c *Coinbase) ReceiveAddress() string {
  var r receiveAddress
  err := json.Unmarshal(c.Get("/account/receive_address"), &r)
  panicOnError(err)
  if r.Success == false {
    panic("Getting the recieve address failed.")
  }
  return r.Address
}
