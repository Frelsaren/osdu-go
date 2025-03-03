package main

import (
	"context"
	"net/url"

	"github.com/Frelsaren/osdu-go/osdu"
)

func main() {
	ctx := context.Background()
	BaseURL, e := url.Parse("http://localhost:8080/")
	if e != nil {
		panic(e)
	}
	token := acquireToken()

	client := osdu.Client{
		BaseURL: BaseURL,
		Token:   &token,
	}

	client.Initialize()

	client.Search.Query(ctx, osdu.QueryParams{
		Kind: []string{"osdu:wks:master-data--Field:1.1.0"},
	})

}

func acquireToken() string {
	return "bearer token"
}
