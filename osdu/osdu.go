package osdu

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	SearchServicePath  = "search/v2"
	StorageServicePath = "storage/v2"
	SchemaServicePath  = "schema-service/v1"
)

type service struct {
	client   *Client
	endpoint string
}

type Client struct {
	client    *http.Client
	Token     *string
	BaseURL   *url.URL
	Partition *string

	Storage *StorageService
	Search  *SearchService
	Schema  *SchemaService
}

func (c *Client) Initialize() {

	if c.client == nil {
		c.client = &http.Client{}
	}

	c.Storage = &StorageService{
		client:   c,
		endpoint: StorageServicePath,
	}
	c.Search = &SearchService{
		client:   c,
		endpoint: SearchServicePath,
	}
	c.Schema = &SchemaService{
		client:   c,
		endpoint: SearchServicePath,
	}
}

func (c *Client) NewRequest(method, urlStr string, body interface{}, urlParams *map[string]string) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("baseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	for k, v := range *urlParams {
		params.Add(k, v)
	}
	u.RawQuery = params.Encode()

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")

	}

	if c.Token != nil {
		req.Header.Set("Authorization", "Bearer "+*c.Token)
	}
	req.Header.Set("Data-Partition-Id", *c.Partition)

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

func (c *Client) BareDo(ctx context.Context, req *http.Request) (*http.Response, error) {
	return c.bareDo(ctx, c.client, req)
}

func (c *Client) bareDo(ctx context.Context, caller *http.Client, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")

	}

	req = withContext(ctx, req)

	resp, err := caller.Do(req)

	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return resp, ctx.Err()
		default:
		}

		return resp, err
	}

	return resp, err
}

func withContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}
