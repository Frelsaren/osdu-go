package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Frelsaren/osdu-go/osdu"
)

func main() {
	ctx := context.Background()
	BaseURL, e := url.Parse("http://localhost:8010/")
	if e != nil {
		panic(e)
	}
	token := acquireToken()
	partition := "default"

	client := osdu.Client{
		BaseURL:   BaseURL,
		Token:     &token,
		Partition: &partition,
	}

	client.Initialize()

	storageHealth, err := client.Storage.IsHealthy(ctx)
	if err != nil {
		panic(err)
	}
	searchHealth, err := client.Search.IsHealthy(ctx)
	if err != nil {
		panic(err)
	}
	schemaHealth, err := client.Schema.IsHealthy(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Storage is up: %t\n", storageHealth)
	fmt.Printf("Search is up: %t\n", searchHealth)
	fmt.Printf("Schema is up: %t\n", schemaHealth)

}

func acquireToken() string {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOaWNlIjoidHJ5In0.1GqgqivdoFiFjXPJvs8suQthSfmC_B_uAnQMDsrXWA0"
}
