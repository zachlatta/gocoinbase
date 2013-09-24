package gocoinbase

import "github.com/tleyden/fakehttp"

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
