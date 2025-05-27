package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestDeleteGroup(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	expectedUrl := fmt.Sprintf("/%s/groups/%s", entitlementServicePath, "test@test.com")
	mux.HandleFunc(expectedUrl, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		// Return an empty response
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	err := client.Entitlement.DeleteGroup(ctx, "test@test.com")
	assertNilError(t, err)

}
