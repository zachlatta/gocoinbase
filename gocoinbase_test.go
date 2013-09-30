package gocoinbase

import "testing"

func TestGet(t *testing.T) {
	expected := "{\"amount\":\"0.00000000\",\"currency\":\"BTC\"}"

	setTestServerResponseBody(expected)

  actual := c.Get("/balance")

	if string(actual) != expected {
		t.Error("Got:", actual, "Expected:", expected)
	}
}

func TestMakeUrl(t *testing.T) {
  c.baseUrl = "https://coinbase.com/api/v1"
	actual, expected := c.makeUrl("/balance"),
		"https://coinbase.com/api/v1/balance?api_key=42"
	if actual != expected {
		t.Error("Got:", actual, "Expected:", expected)
	}
	c.baseUrl = testServer.URL
}
