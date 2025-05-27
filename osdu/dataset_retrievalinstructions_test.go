package osdu

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetRetrievalInstructions(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/retrievalInstructions", datasetServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		if !strings.Contains(r.URL.RawQuery, "expiryTime") {
			t.Errorf("expiryTime not preset")
		}

		fmt.Fprint(w, "{\"datasets\": [{\"id\":\"test1\"},{\"id\":\"test2\"}]}")

	})

	ctx := context.Background()
	input := []string{"test1", "test2"}
	res, err := client.Dataset.GetRetrievalInstructions(ctx, input, "5M")

	assertNilError(t, err)
	assertNoDiff(t, len(input), len(res.Datasets))
}
