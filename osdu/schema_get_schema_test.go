package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetSchema(t *testing.T) {
	client, mux, _ := setup(t)

	id := "test:master-data--Schema:123123"
	mux.HandleFunc(fmt.Sprintf("/%s/schema/%s", schemaServicePath, id), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id": "test:master-data--Schema:123123", "kind": "osdu:wks:Schema:1.0.0", "version": "1.0.0"}`)
	})

	ctx := context.Background()
	var schema any
	err := client.Schema.GetSchema(ctx, id, &schema)

	assertNilError(t, err)
	expectedSchema := map[string]any{
		"id":      "test:master-data--Schema:123123",
		"kind":    "osdu:wks:Schema:1.0.0",
		"version": "1.0.0",
	}

	assertNoDiff(t, schema, expectedSchema)
}
