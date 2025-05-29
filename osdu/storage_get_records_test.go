package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetRecords(t *testing.T) {
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/query/records", storageServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"records": [], "invalidRecords": ["record1", "record2"]}`)
	})

	ctx := context.Background()
	records, err := client.Storage.GetRecords(ctx, GetRecordsParams{
		Records: []string{"record1", "record2"},
	})
	assertNilError(t, err)
	assertNoDiff(t, &records, &RecordsResponse{
		Records:        []Record{},
		InvalidRecords: []string{"record1", "record2"},
	})
}
