package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetUploadURL(t *testing.T) {
	client, mux, _ := setup(t)

	mux.HandleFunc(fmt.Sprintf("/%s/files/uploadURL", fileServicePath), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"FileID": "test.txt", "Location": {}}`)
	})

	resp, err := client.File.GetUploadURL(context.Background(), "5M")
	assertNilError(t, err)

	assertNoDiff(t, resp.FileID, "test.txt")

}
