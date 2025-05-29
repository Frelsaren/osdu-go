package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestMetadataUnDelete(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	testID := "test1"
	expectedPath := fmt.Sprintf("/%s/metadataRecord/%s/unDelete", datasetServicePath, testID)

	mux.HandleFunc(expectedPath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

	})

	ctx := context.Background()
	err := client.Dataset.MetadataUnDelete(ctx, "test1")

	assertNilError(t, err)
}
