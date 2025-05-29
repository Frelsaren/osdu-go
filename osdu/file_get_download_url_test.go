package osdu

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetDownloadURL(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	fileID := "file123"
	expiryTime := "M5"
	mux.HandleFunc(fmt.Sprintf("/%s/files/%s/downloadURL", fileServicePath, fileID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if r.URL.Query().Get("expiryTime") != expiryTime {
			t.Errorf("Expected expiryTime %s, got %s", expiryTime, r.URL.Query().Get("expiryTime"))
		}
		fmt.Fprint(w, `{"SignedUrl": "https://test.com/download/file123"}`)
	})

	ctx := context.Background()

	url, err := client.File.GetDownloadURL(ctx, fileID, expiryTime)

	if err != nil {
		t.Fatalf("GetDownloadURL returned error: %v", err)
	}

	expectedURL := "https://test.com/download/file123"
	if url != expectedURL {
		t.Errorf("GetDownloadURL returned %s, expected %s", url, expectedURL)
	}
}
