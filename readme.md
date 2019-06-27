[![Build Status](https://travis-ci.org/nwehr/logger-go.svg?branch=master)](https://travis-ci.org/nwehr/logger-go)

## Install ##

```
$ go get github.com/nwehr/logger-go
```

## Example Usage ##

```go
package main

import (
	"github.com/nwehr/logger-go"
)

func main() {
	logger.Debug("Hello, World!")
	logger.Warning("The temperature seems a little high!")
	logger.Error("Houston, we have a problem!")

	defer func() {
		if err := recover(); err != nil {
			// Increase the caller depth so we know where panic() was called from
			logger.CallerDepth = logger.DefaultCallerDepth + 1
			logger.Fatal(err)
		}
	}()

	panic("Boom")
}
```
