package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetRecordVersion(t *testing.T) {
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/records/%s/%d", storageServicePath, "1234", 123), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id": "12345","version":123, "data": {"key": "value"}}`)
	})

	ctx := context.Background()
	var record Record
	err := client.Storage.GetRecordVersion(ctx, "1234", 123, &record, nil)
	if err != nil {
		t.Errorf("GetRecord returned error: %v", err)
	}

	expected := Record{ID: "12345", Version: 123, Data: map[string]interface{}{"key": "value"}}
	assertNoDiff(t, record, expected)
}
