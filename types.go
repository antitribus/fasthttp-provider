package fasthttp_provider

import (
	"context"
	"github.com/valyala/fasthttp"
	"time"
)

// HyperHTTP interface
type HyperHTTP interface {
	Do(ctx context.Context, request *fasthttp.Request) (*fasthttp.Response, error)
	DoTimeout(ctx context.Context, request *fasthttp.Request, duration time.Duration) (*fasthttp.Response, error)
	JSON(ctx context.Context, request *fasthttp.Request, response interface{}, duration time.Duration) (*fasthttp.Response, error)
}
