package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateFileMetadata(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	payload := &Record{
		Kind: "osdu:wks:dataset--file",
		Data: map[string]interface{}{
			"fileName": "test.txt",
			"fileSize": 12345,
		},
	}

	mux.HandleFunc(fmt.Sprintf("/%s/files/metadata", fileServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		jsonResponse := `{"id":"12345"}`
		fmt.Fprint(w, jsonResponse)
	})

	ctx := context.Background()

	id, err := client.File.CreateFileMetadata(ctx, "12345", payload)

	assertNilError(t, err)
	assertNoDiff(t, id, "12345")
}
