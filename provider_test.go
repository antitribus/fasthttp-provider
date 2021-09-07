package fasthttp_provider

import (
	"context"
	"github.com/valyala/fasthttp"
	"testing"
	"time"
)

var before = new([]func(ctx context.Context) context.Context)
var after = new([]func(ctx context.Context) context.Context)

var showMessage = func(s string) func(c context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		println(s)
		return ctx
	}
}

var fhp = FastHTTPProvider{
	client: &fasthttp.Client{MaxConnsPerHost: 300,
		MaxConnDuration:     60 * time.Second,
		MaxIdleConnDuration: 1 * time.Second,
		WriteTimeout:        1 * time.Second,
		ReadTimeout:         1 * time.Second,
	},
	beforeRequest: before,
	afterRequest:  after,
}

var duration = time.Millisecond * 1000

type Pokemon struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}

func TestFastHTTPProvider_Do(t *testing.T) {
	ctx := context.Background()

	*before = append(*before, showMessage("before message 1")) //TODO implement verify called function
	*after = append(*after, showMessage("after message 1"))
	*before = append(*before, showMessage("before message 2"))

	req := &fasthttp.Request{}
	req.SetRequestURI("https://pokeapi.co/api/v2/pokemon/ditto")

	pokemon := Pokemon{}

	result, err := fhp.JSON(ctx, req, &pokemon, &duration)

	if err != nil {
		t.Error(err)
	}

	if result.StatusCode() != 200 {
		t.Error("status code dont is 200")
	}

	if len(pokemon.Name) <= 0 {
		t.Error("should return a pokemon with name")
	}

	if pokemon.Order == 0 {
		t.Error("should return a pokemon with order")
	}
}
