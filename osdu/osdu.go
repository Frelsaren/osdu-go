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
	searchServicePath      = "search/v2"
	storageServicePath     = "storage/v2"
	schemaServicePath      = "schema-service/v1"
	datasetServicePath     = "dataset/v1"
	entitlementServicePath = "entitlements/v2"
	fileServicePath        = "file/v2"
)

type service struct {
	client *Client
}

type Client struct {
	client *http.Client
	token  *string

	BaseURL   *url.URL
	Partition string

	common service

	Storage     *StorageService
	Search      *SearchService
	Schema      *SchemaService
	Dataset     *DatasetService
	Entitlement *EntitlementService
	File        *FileService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	httpClient2 := *httpClient
	c := &Client{client: &httpClient2}
	c.initialize()
	return c
}

func (c *Client) InitializeWithToken(token *string) *Client {
	c.token = token
	c.initialize()
	return c
}

func (c *Client) initialize() {

	if c.client == nil {
		c.client = &http.Client{}
	}

	c.common.client = c
	c.Storage = (*StorageService)(&c.common)
	c.Search = (*SearchService)(&c.common)
	c.Schema = (*SchemaService)(&c.common)
	c.Dataset = (*DatasetService)(&c.common)
	c.Entitlement = (*EntitlementService)(&c.common)
	c.File = (*FileService)(&c.common)
}

func (c *Client) NewRequest(method, urlStr string, body interface{}, urlParams *map[string]string) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("baseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if urlParams != nil {
		params := url.Values{}

		for k, v := range *urlParams {
			params.Add(k, v)
		}

		u.RawQuery = params.Encode()
	}

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

	if c.token != nil {
		req.Header.Set("Authorization", "Bearer "+*c.token)
	}
	req.Header.Set("Data-Partition-Id", c.Partition)

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return resp, errors.New(string(bodyBytes))
	}

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
