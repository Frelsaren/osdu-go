package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestDeleteRecord(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	recordID := "record123"
	mux.HandleFunc(fmt.Sprintf("/%s/records/%s:delete", storageServicePath, recordID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()

	err := client.Storage.DeleteRecord(ctx, recordID)

	assertNilError(t, err)
}
