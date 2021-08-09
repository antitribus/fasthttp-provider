package fasthttp_provider

import (
	"context"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"testing"
	"time"
)

func TestFastHTTPProvider_Do(t *testing.T) {
	before := new([]func(ctx context.Context) context.Context)
	after := new([]func(ctx context.Context) context.Context)

	*before = append(*before, func(ctx context.Context) context.Context {
		fmt.Println("Antes do request")
		return ctx
	})

	*after = append(*after, func(ctx context.Context) context.Context {
		fmt.Println("Depois do request")
		return ctx
	})

	hcc := FastHTTPProvider{
		client: &fasthttp.Client{MaxConnsPerHost: 300,
			MaxConnDuration:     60 * time.Second,
			MaxIdleConnDuration: 1 * time.Second,
			WriteTimeout:        1 * time.Second,
			ReadTimeout:         1 * time.Second,
		},
		beforeRequest: before,
		afterRequest:  after,
	}

	req := &fasthttp.Request{}
	req.SetRequestURI("http://product-v3-submarino-npf.internal.b2w.io/product/44172011")

	ctx := context.Background()

	resp, err := hcc.Do(ctx, req)

	if err != nil {
		log.Fatal("Erro na chamada", err.Error())
	}

	m := make(map[string]interface{})
	err = hcc.MarshalResponse(ctx, resp, &m)

	if err != nil {
		log.Println("Erro interface", err.Error())
	}

	log.Println(m["name"])
}
