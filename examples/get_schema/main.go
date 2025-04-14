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

	var schema interface{}
	client.Schema.GetSchema(ctx, "osdu:wks:master-data--Field:1.1.0", &schema)

}

func acquireToken() string {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOaWNlIjoidHJ5In0.1GqgqivdoFiFjXPJvs8suQthSfmC_B_uAnQMDsrXWA0"
}
