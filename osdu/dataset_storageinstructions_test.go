package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGenerateStorageInstructions(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/storageInstructions", datasetServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQuery(t, r, "expiryTime", "5M")
		testQuery(t, r, "kindSubtype", "osdu:example:dataset:1.0.0")
		fmt.Fprint(w, `{"storageLocation": {"type": "s3", "bucket": "example-bucket"}, "providerKey": "test"}`)
	})

	ctx := context.Background()
	storageInstructions, err := client.Dataset.GenerateStorageInstructions(ctx, "osdu:example:dataset:1.0.0", "5M")
	assertNilError(t, err)
	assertNoDiff(t, *storageInstructions, StorageInstructions{
		StorageLocation: map[string]any{
			"type":   "s3",
			"bucket": "example-bucket",
		},
		ProviderKey: "test",
	})
}
