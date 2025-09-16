package azure

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const virtualDir = "/webhooks"

var hook *Webhook

func TestMain(m *testing.M) {
	// setup
	var err error
	hook, err = New()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
	// teardown
}

func newServer(handler http.HandlerFunc) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(virtualDir, handler)
	return httptest.NewServer(mux)
}
