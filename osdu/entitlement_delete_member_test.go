package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestDeleteMember(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	memberEmail := "test.member@test.com"
	expectedUrl := fmt.Sprintf("/%s/members/%s", entitlementServicePath, memberEmail)
	mux.HandleFunc(expectedUrl, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		// Return an empty response
		w.WriteHeader(http.StatusNoContent)
	})
	ctx := context.Background()
	err := client.Entitlement.DeleteMember(ctx, memberEmail)
	assertNilError(t, err)

}
