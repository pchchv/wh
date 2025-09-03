package github

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
	hook, err = New(Options.Secret("IsWishesWereHorsesWedAllBeEatingSteak!"))
	if err != nil {
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
