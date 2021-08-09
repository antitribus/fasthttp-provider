package fasthttp_provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

type FastHTTPProvider struct {
	client        *fasthttp.Client
	beforeRequest *[]func(ctx context.Context) context.Context
	afterRequest  *[]func(ctx context.Context) context.Context
}

func (hhc *FastHTTPProvider) Request(ctx context.Context, request *fasthttp.Request, duration *time.Duration) (*fasthttp.Response, error) {
	if hhc.beforeRequest != nil {
		for _, br := range *hhc.beforeRequest {
			br(ctx)
		}
	}

	var err error
	response := fasthttp.AcquireResponse()

	if duration != nil && *duration > 0 {
		err = hhc.client.DoTimeout(request, response, *duration)
	} else {
		err = hhc.client.Do(request, response)
	}

	defer fasthttp.ReleaseRequest(request)

	if hhc.afterRequest != nil {
		for _, ar := range *hhc.afterRequest {
			defer ar(ctx)
		}
	}

	return response, err
}

func (hhc *FastHTTPProvider) Do(ctx context.Context, request *fasthttp.Request) (*fasthttp.Response, error) {
	return hhc.Request(ctx, request, nil)
}

func (hhc *FastHTTPProvider) DoTimeout(ctx context.Context, request *fasthttp.Request, duration *time.Duration) (*fasthttp.Response, error) {
	return hhc.Request(ctx, request, duration)
}

func (hhc *FastHTTPProvider) MarshalResponse(ctx context.Context, response *fasthttp.Response, resp interface{}) error {
	err := json.Unmarshal(response.Body(), &resp)
	return err
}

func (hhc *FastHTTPProvider) OnRequestError(ctx context.Context, err error) error {
	fmt.Println("on-request-error", err.Error())
	return err
}

func (hhc *FastHTTPProvider) OnMarshalResponseError(ctx context.Context, err error) error {
	fmt.Println("on-marshal-error", err.Error())
	return err
}
