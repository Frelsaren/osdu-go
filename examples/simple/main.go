package main

import (
	"context"
	"fmt"
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
	partition := "partition"

	client := osdu.Client{
		BaseURL:   BaseURL,
		Token:     &token,
		Partition: &partition,
	}

	client.Initialize()

	searchResults, err := client.Search.Query(ctx, osdu.QueryParams{
		Kind: []string{"osdu:wks:master-data--Field:1.1.0"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(len(searchResults.Results))

	var sampleInterface interface{}
	err = client.Storage.GetRecord(ctx, "id", &sampleInterface, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(sampleInterface.(map[string]interface{})["id"])

}

func acquireToken() string {
	return "token"
}
