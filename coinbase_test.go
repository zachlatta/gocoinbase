package gocoinbase

import (
	"github.com/tleyden/fakehttp"
	"testing"
)

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

func makeServerWithResponseBody(body string) *fakehttp.HTTPServer {
	testServer := fakehttp.NewHTTPServer()
	testServer.Start()

  headers := make(map[string]string)
  headers["Connection"] = "keep-alive"
  headers["Content-Type"] = "application/json; charset=utf-8"
  headers["Server"] = "cloudflare-nginx"

	testServer.Response(200, headers, body)

  return testServer
}
