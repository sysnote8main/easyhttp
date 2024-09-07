package httpclient

import (
	"context"
	"net/http"

	"golang.org/x/time/rate"
)

type HTTPClient struct {
	client      *http.Client
	RateLimiter *rate.Limiter
}

func NewClient(limiter *rate.Limiter) *HTTPClient {
	return &HTTPClient{
		client:      http.DefaultClient,
		RateLimiter: limiter,
	}
}

func (c *HTTPClient) CheckRate() rate.Limit {
	return c.RateLimiter.Limit()
}

func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	ctx := context.Background()
	err := c.RateLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
