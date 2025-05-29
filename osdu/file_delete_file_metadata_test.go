package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestDeleteFileMetadata(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	fileID := "file123"
	mux.HandleFunc(fmt.Sprintf("/%s/files/%s/metadata", fileServicePath, fileID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()

	err := client.File.DeleteFileMetadata(ctx, fileID)

	assertNilError(t, err)
}
