package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/VictoriaMetrics/fastcache"
	"github.com/go-resty/resty/v2"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/httputils"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"net/http"
)

const (
	apiURL           = "https://pokeapi.co/api"
	apiVersion       = "v2"
	defaultCacheSize = 100
)

// ClientOption to configure API client
type ClientOption func(*Client)

// HTTPRequestOption to configure API client
type HTTPRequestOption func(interface{})

// Client represents the pokeapi client.
type Client struct {
	httpClient *resty.Client
	cache      *fastcache.Cache
	CacheSize  int

	Resource ResourcesClient
	Berry    BerriesClient
}

// NewClient configures and returns a new pokeapi client.
// You can override some default configuration using ClientOption.
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		httpClient: resty.New(),
		CacheSize:  defaultCacheSize,
	}

	c.Resource = &resourcesClient{client: c}
	c.Berry = &berriesClient{client: c}

	for _, opt := range opts {
		opt(c)
	}

	c.cache = fastcache.New(c.CacheSize)

	return c
}

// WithHTTPClient overrides the default http.Client
func WithHTTPClient(client *resty.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

// WithCacheSize overrides the default cacheSize
func WithCacheSize(cacheSize int) ClientOption {
	return func(c *Client) {
		c.CacheSize = cacheSize
	}
}

// request makes a request to the pokeapi APIs to get the data.
// Data is also cached based on the endpoint that is called and if the same endpoint is called again
// then the data is returned from cache.
func (c *Client) request(url string, res interface{}) *errors.RestErr {

	if c.cache.Has([]byte(url)) {
		if err := json.Unmarshal(c.cache.Get(nil, []byte(url)), res); err != nil {
			return errors.InternalServerError(err.Error())
		}
		return nil
	}

	resp, err := c.httpClient.SetBaseURL(fmt.Sprintf("%s/%s", apiURL, apiVersion)).
		SetHeader(httputils.ContentTypeHeaderKey, httputils.ApplicationJsonMIMEType).
		SetHeader(httputils.AcceptHeaderKey, httputils.ApplicationJsonMIMEType).
		R().SetResult(res).Get(url)
	logger.RestyDebugLogs(resp)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		c.cache.Set([]byte(url), resp.Body())
		return nil
	case http.StatusNotFound:
		return errors.NotFoundErrorf("")
	default:
		return httputils.CheckHTTPErrorStatusCode(resp.StatusCode(), "")
	}
}
