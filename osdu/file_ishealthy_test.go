package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestFileIsHealthy(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/liveness_check", fileServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, "File service is alive.")
	})

	ctx := context.Background()
	isHealthy, err := client.File.IsHealthy(ctx)

	assertNilError(t, err)

	assertNoDiff(t, true, isHealthy)
}
