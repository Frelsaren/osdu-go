package osdu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestQueryWithCursor(t *testing.T) {
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/query_with_cursor", searchServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, SearchTestResult)
	})

	var payload QueryWithCursorParams
	var expected SearchResponse

	assertNilError(t, json.Unmarshal([]byte(SearchTestPayload), &payload))
	assertNilError(t, json.Unmarshal([]byte(SearchTestResult), &expected))

	ctx := context.Background()
	res, err := client.Search.QueryWithCursor(ctx, payload)

	assertNilError(t, err)
	assertNoDiff(t, res.Cursor, expected.Cursor)
}

const (
	SearchWithCursorTestPayload = `{
		"kind": ["osdu:wks:wellbore:1.0.0"]
	}`
	SearchWithCursorTestResult = `{
		"results": [],
		"totalCount": 0,
		"cursor": "nextCursorValue"
	}`
)
