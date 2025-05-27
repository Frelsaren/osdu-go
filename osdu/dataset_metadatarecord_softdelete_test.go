package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestMetadataSoftDelete(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	testID := "test1"
	expectedPath := fmt.Sprintf("/%s/metadataRecord/%s/softdelete", datasetServicePath, testID)

	mux.HandleFunc(expectedPath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

	})

	ctx := context.Background()
	err := client.Dataset.MetadataSoftDelete(ctx, "test1")

	if err != nil {
		t.Errorf("error was thrown")
	}
}
