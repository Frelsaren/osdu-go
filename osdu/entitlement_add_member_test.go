package osdu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestAddMember(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	payload := AddMemberResponse{
		Email: "test@test.com",
		Role:  "member",
	}

	mux.HandleFunc(fmt.Sprintf("/%s/groups/%s/members", entitlementServicePath, payload.Email), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		// jsonify the payload and return it
		jsonResponse, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("Failed to marshal payload: %v", err)
		}
		fmt.Fprint(w, string(jsonResponse))
	})

	ctx := context.Background()
	res, err := client.Entitlement.AddMember(ctx, payload.Email, payload.Email, payload.Role)
	assertNilError(t, err)
	assertNoDiff(t, res, payload)
}
