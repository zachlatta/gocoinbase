package gocoinbase

import "testing"

func TestGet(t *testing.T) {
  expected := "{\"amount\":\"0.00000000\",\"currency\":\"BTC\"}"

  testServer := makeServerWithResponseBody(expected)
  defer testServer.Flush()

  baseUrl = testServer.URL

  actual := Get("/balance")

  if string(actual) != expected {
    t.Error("Got:", actual, "Expected:", expected)
  }
}

func TestMakeUrl(t *testing.T) {
	Init("42")
	actual, expected := makeUrl("/balance"),
		"https://coinbase.com/api/v1/balance?api_key=42"
	if actual != expected {
		t.Error("Got:", actual, "Expected:", expected)
	}
}
