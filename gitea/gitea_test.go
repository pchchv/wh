package gitea

import (
	"net/http"
	"net/http/httptest"
)

const path = "/webhooks"

func newServer(handler http.HandlerFunc) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(path, handler)
	return httptest.NewServer(mux)
}
