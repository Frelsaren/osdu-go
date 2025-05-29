package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetRecordsOfKind(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	var limit int32 = 10
	payload := GetRecordOfKindParams{
		Kind:  "osdu:wks:dataset--Dataset:1.0.0",
		Limit: &limit,
	}
	mux.HandleFunc(fmt.Sprintf("/%s/records", storageServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"results": []}`)
	})

	ctx := context.Background()

	records, err := client.Storage.GetRecordsOfKind(ctx, payload)

	assertNilError(t, err)
	assertNoDiff(t, &records, &RecordsOfKindResponse{
		Results: []string{},
	})

}
