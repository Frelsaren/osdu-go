package osdu

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestRegisterDataset(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/registerDataset", datasetServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			assertNilError(t, err)

		}

		fmt.Fprint(w, string(body))
	})

	ctx := context.TODO()

	input := RegisterDatasetRequest{
		DatasetRegistries: []Record{
			{
				ID: "test1",
			},
		}}
	res, err := client.Dataset.RegisterDataset(ctx, input)

	assertNilError(t, err)

	assertNoDiff(t, input, *res)
}
