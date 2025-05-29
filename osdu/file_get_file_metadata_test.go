package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetFileMetadata(t *testing.T) {
	client, mux, _ := setup(t)

	id := "test:master-data--File:123123"
	mux.HandleFunc(fmt.Sprintf("/%s/files/%s/metadata", fileServicePath, id), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id": "test:master-data--File:123123", "kind": "osdu:wks:File:1.0.0", "version": "1.0.0"}`)
	})
	ctx := context.Background()
	metadata := &Record{}
	err := client.File.GetFileMetadata(ctx, id, metadata)

	assertNilError(t, err)
	expectedMetadata := &Record{
		ID:      "test:master-data--File:123123",
		Kind:    "osdu:wks:File:1.0.0",
		Version: 123,
	}

	assertNoDiff(t, metadata, expectedMetadata)
}
