package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestDatasetIsHealthy(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/liveness_check", datasetServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, "Dataset service is alive.")
	})

	ctx := context.Background()
	isHealthy, err := client.Dataset.IsHealthy(ctx)

	if err != nil {
		t.Errorf(("Received error"))
	}
	if !isHealthy {
		t.Errorf("isHealhty returned false")
	}
}
