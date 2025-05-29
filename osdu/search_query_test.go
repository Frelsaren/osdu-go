package osdu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestQuery(t *testing.T) {
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/query", searchServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, SearchTestResult)
	})

	var payload QueryParams
	var expected SearchResponse

	assertNilError(t, json.Unmarshal([]byte(SearchTestPayload), &payload))
	assertNilError(t, json.Unmarshal([]byte(SearchTestResult), &expected))

	ctx := context.Background()
	res, err := client.Search.Query(ctx, payload)

	assertNilError(t, err)
	assertNoDiff(t, res, expected)
}

const (
	SearchTestPayload = `{
		"kind": ["osdu:wks:wellbore:1.0.0"]
	}`
	SearchTestResult = `{
		"results": [],
		"totalCount": 0
	}`
)
