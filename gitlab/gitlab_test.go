package gitlab

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const path = "/webhooks"

var hook *Webhook

func TestMain(m *testing.M) {
	// setup
	var err error
	if hook, err = New(Options.Secret("sampleToken!")); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
	// teardown
}

func newServer(handler http.HandlerFunc) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(path, handler)
	return httptest.NewServer(mux)
}
