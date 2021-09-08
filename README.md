# fasthttp-provider

## Install

```shell
go get github.com/antitribus/fasthttp-provider
```

## Example:

```go
package main

import (
	"context"
	pvd "fasthttp-provider"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

var before = new([]func(ctx context.Context))
var after = new([]func(ctx context.Context))

var showMessage = func(s string) func(c context.Context) {
	return func(ctx context.Context) {
		println(s)
	}
}

var fhp = pvd.FastHTTPProvider{
	Client: &fasthttp.Client{MaxConnsPerHost: 300,
		MaxConnDuration:     60 * time.Second,
		MaxIdleConnDuration: 1 * time.Second,
		WriteTimeout:        1 * time.Second,
		ReadTimeout:         1 * time.Second,
	},
	BeforeRequest: before,
	AfterRequest:  after,
}

var duration = time.Millisecond * 500

type Pokemon struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}

func main() {
	ctx := context.Background()

	*before = append(*before, showMessage("before message 1"))
	*after = append(*after, showMessage("after message 1"))
	*before = append(*before, showMessage("before message 2"))

	req := &fasthttp.Request{}
	req.SetRequestURI("https://pokeapi.co/api/v2/pokemon/ditto")

	pokemon := Pokemon{}

	completeResponse, _ := fhp.JSON(ctx, req, &pokemon, &duration)

	fmt.Println(pokemon.Name) // ditto
	fmt.Println(completeResponse.StatusCode()) // 200
}
```
