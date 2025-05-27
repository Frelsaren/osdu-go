package osdu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateGRoup(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	payload := CreateGroupBody{
		Name:        "Test Group",
		Description: "This is a test group",
	}

	mux.HandleFunc(fmt.Sprintf("/%s/groups", entitlementServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		// jsonify the payload and return it
		jsonResponse, err := json.Marshal(CreateGroupResponse{
			Email:       "test@test.com",
			Name:        payload.Name,
			Description: payload.Description,
		})
		assertNilError(t, err)
		fmt.Fprint(w, string(jsonResponse))
	})

	ctx := context.Background()

	group, err := client.Entitlement.CreateGroup(ctx, payload)

	assertNilError(t, err)

	assertNoDiff(t, group, CreateGroupResponse{
		Email:       "test@test.com",
		Name:        payload.Name,
		Description: payload.Description,
	})
}
