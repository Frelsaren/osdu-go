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
	partition := "default"

	client := osdu.Client{
		BaseURL:   BaseURL,
		Token:     &token,
		Partition: &partition,
	}

	client.Initialize()

	params := osdu.GetRecordOfKindParams{
		Kind: "osdu:wks:master-data--Field:1.0.0",
	}

	client.Storage.GetRecordsOfKind(ctx, params)

}

func acquireToken() string {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOaWNlIjoidHJ5In0.1GqgqivdoFiFjXPJvs8suQthSfmC_B_uAnQMDsrXWA0"
}
