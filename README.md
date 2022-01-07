# pokeapi

An API client for [pokeapi](https://pokeapi.co) implemented in Golang.

# Installation

```
go get github.com/atselvan/pokeapi
```

# Go Mod

```
require github.com/atselvan/pokeapi 
```

# Example

Make a new `Client` then use the clients methods to access the resources.

```go
package main

import (
	"github.com/atselvan/pokeapi"
)

func main() {
	client := pokeapi.NewClient()

	berry, restErr := client.Berry.Get("1") // Can also be the name of the berry
	if restErr != nil {
		// do something
	}
}
```