package gocoinbase

import "testing"

func TestBalance(t *testing.T) {
	response := "{\"amount\":\"2.32400000\",\"currency\":\"BTC\"}"

	setTestServerResponseBody(response)

	expected := 2.324
	actual := c.Balance()

	if actual != expected {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}
