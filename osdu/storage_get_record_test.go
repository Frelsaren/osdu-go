package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetRecord(t *testing.T) {
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/records/%s", storageServicePath, "12345"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id": "12345", "data": {"key": "value"}}`)
	})

	ctx := context.Background()
	var record Record
	err := client.Storage.GetRecord(ctx, "12345", &record, nil)
	if err != nil {
		t.Errorf("GetRecord returned error: %v", err)
	}

	expected := Record{ID: "12345", Data: map[string]interface{}{"key": "value"}}
	assertNoDiff(t, record, expected)
}
