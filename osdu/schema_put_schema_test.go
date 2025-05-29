package osdu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestPutSchema(t *testing.T) {
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/schema", schemaServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, SchemaTestResult)
	})

	var payload WriteSchemaBody
	var expected SchemaInfo

	assertNilError(t, json.Unmarshal([]byte(SchemaTestPayload), &payload))
	assertNilError(t, json.Unmarshal([]byte(SchemaTestResult), &expected))

	ctx := context.Background()
	res, err := client.Schema.PutSchema(ctx, payload)

	assertNilError(t, err)

	assertNoDiff(t, res.SchemaInfo, expected)
}
