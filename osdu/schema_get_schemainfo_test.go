package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetSchemaInfo(t *testing.T) {
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/schema", schemaServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, testresponse)
	})

	ctx := context.Background()
	resp, err := client.Schema.GetSchemaInfo(ctx, GetSchemaInfoParams{})

	assertNilError(t, err)
	assertNoDiff(t, resp.SchemaInfos[0].SchemaIdentity.ID, "osdu:wks:wellbore:1.0.0")
}

const testresponse = `{
  "schemaInfos": [
    {
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
    }
  ],
  "offset": 0,
  "count": 0,
  "totalCount": 0
}`
