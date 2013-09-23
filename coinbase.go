package gocoinbase

import (
	"io/ioutil"
	"net/http"
)

var baseUrl, apiKey string

func Init(key string) {
	apiKey = key
	baseUrl = "https://coinbase.com/api/v1"
}

func Get(path string) []byte {
	res, err := http.Get(makeUrl(path))
	panicOnError(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	panicOnError(err)
	return body
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func makeUrl(relativePath string) string {
	return baseUrl + relativePath + "?api_key=" + apiKey
}
