package gocoinbase

import (
	"github.com/tleyden/fakehttp"
)

var testServer *fakehttp.HTTPServer
var c *Coinbase

func init() {
  c = Init("42")
	testServer = makeServer()
	c.baseUrl = testServer.URL
}

func makeServer() *fakehttp.HTTPServer {
	testServer := fakehttp.NewHTTPServer()
	testServer.Start()
	return testServer
}

func setTestServerResponseBody(body string) {
	headers := make(map[string]string)
	headers["Connection"] = "keep-alive"
	headers["Content-Type"] = "application/json; charset=utf-8"
	headers["Server"] = "cloudflare-nginx"

	testServer.Response(200, headers, body)
}
