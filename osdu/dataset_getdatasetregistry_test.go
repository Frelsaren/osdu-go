package osdu

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestGetDatasetRegistries(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/getDatasetRegistry", datasetServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		jsonData, _ := json.Marshal(DatasetRegistry{
			DatasetRegistries: []Record{
				{
					ID:   "Test1",
					Kind: "wks:osdu:master-data--dataset:1.0.0",
				},
			},
		})
		fmt.Fprint(w, string(jsonData))
	})

	ctx := context.Background()
	registries, err := client.Dataset.GetDatasetRegistries(ctx, []string{"Test1"})
	if err != nil {
		t.Errorf("an error was thrown")
	}

	if len(registries.DatasetRegistries) != 1 {
		t.Errorf("didnt return expected amount of results")
	}
}
