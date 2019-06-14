```go
package main

import (
	"gitlab.com/nwehr/logger-go"
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