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
	partition := "default"

	client := osdu.Client{
		BaseURL:   BaseURL,
		Token:     &token,
		Partition: &partition,
	}

	client.Initialize()

	var record osdu.Record
	client.Storage.GetRecord(ctx, "1234567890", &record, nil)

	fmt.Println(record.ID)

}

func acquireToken() string {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOaWNlIjoidHJ5In0.1GqgqivdoFiFjXPJvs8suQthSfmC_B_uAnQMDsrXWA0"
}
