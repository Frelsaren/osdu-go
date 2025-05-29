package osdu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestPutSystemSchema(t *testing.T) {
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/schema/system", schemaServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, SchemaTestResult)
	})

	var payload WriteSchemaBody
	var expected SchemaInfo

	assertNilError(t, json.Unmarshal([]byte(SchemaTestPayload), &payload))
	assertNilError(t, json.Unmarshal([]byte(SchemaTestResult), &expected))

	ctx := context.Background()
	res, err := client.Schema.PutSystemSchema(ctx, payload)

	assertNilError(t, err)

	assertNoDiff(t, res.SchemaInfo, expected)
}
