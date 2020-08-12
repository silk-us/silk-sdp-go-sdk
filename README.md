# Silk SDP Go SDK

# Installation

```go get github.com/silk-us/silk-sdp-go-sdk/silksdp```

# Example

```go
package main

import (
	"fmt"
        "log"
	
	"github.com/silk-us/silk-sdp-go-sdk/silksdp"
)

func main() {

	silk, err := silksdp.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}
	
	// GET hosts on the Silk server
	getHosts, err := silk.Get("/hosts")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(getHosts.(map[string]interface{})["hits"])

	// Simplified Function to get hosts on the Silk server
	getHosts, err := silk.GetHosts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(get)

}
```

# Documentation

* [SDK for Go Documentation](https://godoc.org/github.com/silk-us/silk-sdp-go-sdk/silksdp)

