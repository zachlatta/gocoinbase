package gocoinbase

import (
	"io/ioutil"
	"net/http"
)

type Coinbase struct {
  apiKey string
  baseUrl string
}

func Init(key string) *Coinbase {
  return &Coinbase{apiKey: key, baseUrl: "https://coinbase.com/api/v1"}
}

func (c *Coinbase) Get(path string) []byte {
	res, err := http.Get(c.makeUrl(path))
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

func (c *Coinbase) makeUrl(relativePath string) string {
	return c.baseUrl + relativePath + "?api_key=" + c.apiKey
}
