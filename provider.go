package fasthttp_provider

import (
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"time"
)

type FastHTTPProvider struct {
	Client        *fasthttp.Client
	BeforeRequest *[]func(ctx context.Context) context.Context
	AfterRequest  *[]func(ctx context.Context) context.Context
}

func (fhp *FastHTTPProvider) Request(ctx context.Context, request *fasthttp.Request, duration *time.Duration) (*fasthttp.Response, error) {
	if fhp.BeforeRequest != nil {
		for _, br := range *fhp.BeforeRequest {
			br(ctx)
		}
	}

	var err error
	response := fasthttp.AcquireResponse()

	if duration != nil && *duration > 0 {
		err = fhp.Client.DoTimeout(request, response, *duration)
	} else {
		err = fhp.Client.Do(request, response)
	}

	defer fasthttp.ReleaseRequest(request)

	if fhp.AfterRequest != nil {
		for _, ar := range *fhp.AfterRequest {
			ar(ctx)
		}
	}

	return response, err
}

func (fhp *FastHTTPProvider) Do(ctx context.Context, request *fasthttp.Request) (*fasthttp.Response, error) {
	return fhp.Request(ctx, request, nil)
}

func (fhp *FastHTTPProvider) DoTimeout(ctx context.Context, request *fasthttp.Request, duration *time.Duration) (*fasthttp.Response, error) {
	return fhp.Request(ctx, request, duration)
}

func (fhp *FastHTTPProvider) JSON(ctx context.Context, request *fasthttp.Request, response interface{}, duration *time.Duration) (*fasthttp.Response, error) {
	result, err := fhp.Request(ctx, request, duration)

	if err == nil {
		_ = json.Unmarshal(result.Body(), &response)
	}

	return result, err
}
