package main

import (
	"context"
	"net/url"

	"github.com/Frelsaren/osdu-go/osdu"
)

func main() {
	ctx := context.Background()
	BaseURL, err := url.Parse("http://localhost:8080/")
	if err != nil {
		panic(err)
	}
	token := acquireToken()
	partition := "default"

	client := osdu.NewClient(nil).InitializeWithToken(&token)
	client.BaseURL = BaseURL
	client.Partition = partition

	var schema any
	client.Schema.GetSchema(ctx, "osdu:wks:master-data--Field:1.1.0", &schema)

}

func acquireToken() string {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOaWNlIjoidHJ5In0.1GqgqivdoFiFjXPJvs8suQthSfmC_B_uAnQMDsrXWA0"
}
