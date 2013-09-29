# gocoinbase [![Build Status](https://travis-ci.org/zachlatta/gocoinbase.png?branch=master)](https://travis-ci.org/zachlatta/gocoinbase)

Go wrapper for the Coinbase API.

## Installation

    go get github.com/zachlatta/gocoinbase

Gocoinbase connects with your Coinbase account using your API key. Start out by 
[enabling an API key](https://coinbase.com/account/integrations) on your
account.

Next, create an instance of the gocoinbase client and give it your API key.
Notice that we don't hard code our API key into our application. Instead, we
load it from an environment variable called `COINBASE_API_KEY`. Remember to
import `os` in your application.

```go
coinbase := gocoinbase.Init(os.Getenv("COINBASE_API_KEY"))
```
