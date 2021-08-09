package fasthttp_provider

import (
	"context"
	"github.com/valyala/fasthttp"
	"time"
)

type HyperHTTP interface {
	Do(ctx context.Context, request *fasthttp.Request) (response *fasthttp.Response, err error)
	DoTimeout(ctx context.Context, request *fasthttp.Request, timeout time.Duration) (response *fasthttp.Response, err error)
	MarshalResponse(ctx context.Context, response *fasthttp.Response, respInt *interface{}) error
	OnRequestError(ctx context.Context, err error) error
	OnMarshalResponseError(ctx context.Context, err error) error
}
