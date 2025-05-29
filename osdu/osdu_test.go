package osdu

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const baseURLPath = "/api"

func setup(t *testing.T) (client *Client, mux *http.ServeMux, serverURL string) {
	t.Helper()

	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr, "FAIL: Client.BaseURL path prefix is not preserved in the request URL:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\tDid you accidentally use an absolute endpoint URL rather than relative?")
		http.Error(w, "Client.BaseURL path prefix is not preserved in the request URL.", http.StatusInternalServerError)
	})

	server := httptest.NewServer(apiHandler)

	// Create a custom transport with isolated connection pool
	transport := &http.Transport{
		// Controls connection reuse - false allows reuse, true forces new connections for each request
		DisableKeepAlives: false,
		// Maximum concurrent connections per host (active + idle)
		MaxConnsPerHost: 10,
		// Maximum idle connections maintained per host for reuse
		MaxIdleConnsPerHost: 5,
		// Maximum total idle connections across all hosts
		MaxIdleConns: 20,
		// How long an idle connection remains in the pool before being closed
		IdleConnTimeout: 20 * time.Second,
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	client = NewClient(httpClient)

	url, _ := url.Parse(server.URL + baseURLPath + "/")
	client.BaseURL = url
	client.Partition = "test"

	t.Cleanup(server.Close)

	return client, mux, server.URL

}

func TestNewClient(t *testing.T) {
	client := NewClient(nil)
	if client == nil {
		t.Fatal("NewClient returned nil")
	}
	if client.client == nil {
		t.Fatal("NewClient returned a client with nil http.Client")
	}

	assertNoDiff(t, client.client, client.Dataset.client.client)

}

func TestInitializeWithToken(t *testing.T) {
	token := "test-token"
	client := NewClient(nil).InitializeWithToken(&token)
	assertNoDiff(t, client.token, &token)

	if client.Storage == nil {
		t.Fatal("NewClient returned a client with nil StorageService")
	}
	if client.Search == nil {
		t.Fatal("NewClient returned a client with nil SearchService")
	}
	if client.Schema == nil {
		t.Fatal("NewClient returned a client with nil SchemaService")
	}
	if client.Dataset == nil {
		t.Fatal("NewClient returned a client with nil DatasetService")
	}
	if client.Entitlement == nil {
		t.Fatal("NewClient returned a client with nil EntitlementService")
	}
	if client.File == nil {
		t.Fatal("NewClient returned a client with nil FileService")
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testQuery(t *testing.T, r *http.Request, key, value string) {
	if r.URL.Query().Get(key) != value {
		t.Errorf("Expected query parameter %s to be %s, got %s", key, value, r.URL.Query().Get(key))
	}
}

func assertNoDiff(t *testing.T, want, got interface{}) {
	t.Helper()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("diff mismatch (-want +got):\n%v", diff)
	}
}

func assertNilError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
