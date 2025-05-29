package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestDeleteRecords(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	ids := []string{"record1", "record2"}
	mux.HandleFunc(fmt.Sprintf("/%s/records/delete", storageServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()

	err := client.Storage.DeleteRecords(ctx, ids)

	assertNilError(t, err)
}
