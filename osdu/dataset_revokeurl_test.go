package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestRevokeURL(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/revokeURL", datasetServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	err := client.Dataset.RevokeURL(ctx, "testtype", map[string]string{})

	assertNilError(t, err)
}
