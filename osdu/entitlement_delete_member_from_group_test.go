package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestDeleteMemberFromGroup(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	memberEmail := "test.member@test.com"
	groupEmail := "test.group@test.com"
	expectedUrl := fmt.Sprintf("/%s/groups/%s/members/%s", entitlementServicePath, groupEmail, memberEmail)
	mux.HandleFunc(expectedUrl, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		// Return an empty response
		w.WriteHeader(http.StatusNoContent)
	})
	ctx := context.Background()
	err := client.Entitlement.DeleteMemberFromGroup(ctx, memberEmail, groupEmail)
	assertNilError(t, err)
}
