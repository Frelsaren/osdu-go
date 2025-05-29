package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetRecordVersions(t *testing.T) {
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/records/versions/%s", storageServicePath, "test:master-data--test-record:1123"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"recordId": "test:master-data--test-record:1123", "versions": [1111, 2222, 3333]}`)
	})

	ctx := context.Background()
	resp, err := client.Storage.GetRecordVersions(ctx, "test:master-data--test-record:1123")
	if err != nil {
		t.Errorf("Storage.GetRecordVersions returned error: %v", err)
	}

	want := RecordVersionsResponse{
		RecordID: "test:master-data--test-record:1123",
		Versions: []int64{1111, 2222, 3333},
	}

	assertNoDiff(t, resp, want)

}
