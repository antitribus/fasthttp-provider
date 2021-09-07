package fasthttp_provider

import (
	"context"
	"github.com/valyala/fasthttp"
	"time"
)

type HyperHTTP interface {
	Do(ctx context.Context, request *fasthttp.Request) (response *fasthttp.Response, err error)
	DoTimeout(ctx context.Context, request *fasthttp.Request, duration time.Duration) (response *fasthttp.Response, err error)
	JSON(ctx context.Context, request *fasthttp.Request, response interface{}, duration time.Duration)
}
