package osdu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestPostSchema(t *testing.T) {
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/schema", schemaServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, SchemaTestResult)
	})

	var payload WriteSchemaBody
	var expected SchemaInfo

	assertNilError(t, json.Unmarshal([]byte(SchemaTestPayload), &payload))
	assertNilError(t, json.Unmarshal([]byte(SchemaTestResult), &expected))

	ctx := context.Background()
	res, err := client.Schema.PostSchema(ctx, payload)

	assertNilError(t, err)

	assertNoDiff(t, res.SchemaInfo, expected)
}

const (
	SchemaTestPayload = `{
  "schemaInfo": {
    "schemaIdentity": {
      "authority": "osdu",
      "source": "wks",
      "entityType": "wellbore",
      "schemaVersionMajor": 1,
      "schemaVersionMinor": 1,
      "schemaVersionPatch": 0,
      "id": "osdu:wks:wellbore:1.0.0"
    },
    "createdBy": "user@opendes.com",
    "dateCreated": "2019-05-23T11:16:03Z",
    "status": "PUBLISHED",
    "scope": "INTERNAL",
    "supersededBy": {
      "authority": "osdu",
      "source": "wks",
      "entityType": "wellbore",
      "schemaVersionMajor": 1,
      "schemaVersionMinor": 1,
      "schemaVersionPatch": 0,
      "id": "osdu:wks:wellbore:1.0.0"
    }
  },
  "schema": {}
}`
	SchemaTestResult = `{
  "schemaIdentity": {
    "authority": "osdu",
    "source": "wks",
    "entityType": "wellbore",
    "schemaVersionMajor": 1,
    "schemaVersionMinor": 1,
    "schemaVersionPatch": 0,
    "id": "osdu:wks:wellbore:1.0.0"
  },
  "createdBy": "user@opendes.com",
  "dateCreated": "2019-05-23T11:16:03Z",
  "status": "PUBLISHED",
  "scope": "INTERNAL",
  "supersededBy": {
    "authority": "osdu",
    "source": "wks",
    "entityType": "wellbore",
    "schemaVersionMajor": 1,
    "schemaVersionMinor": 1,
    "schemaVersionPatch": 0,
    "id": "osdu:wks:wellbore:1.0.0"
  }
}`
)
