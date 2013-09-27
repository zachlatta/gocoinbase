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

func TestReceiveAddress(t *testing.T) {
  response := "{\"success\":true,\"address\":\"1P8idYiG7AwwLbXNrMbNcfZSY4DG" +              "92CcWw\",\"callback_url\":null}"

  setTestServerResponseBody(response)

  expected := "1P8idYiG7AwwLbXNrMbNcfZSY4DG92CcWw"
  actual := c.ReceiveAddress()

	if actual != expected {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}
