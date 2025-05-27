package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestEntitlementIsHealthy(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/_ah/liveness_check", entitlementServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, "Entitlement service is alive.")
	})

	ctx := context.Background()
	isHealthy, err := client.Entitlement.IsHealthy(ctx)

	if err != nil {
		t.Errorf(("Received error"))
	}
	if !isHealthy {
		t.Errorf("isHealhty returned false")
	}
}
