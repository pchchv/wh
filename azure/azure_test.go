package azure

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestBasicAuth(t *testing.T) {
	const user = "user"
	const pass = "pass123"
	opt := Options.BasicAuth(user, pass)
	h := &Webhook{}
	err := opt(h)

	assert.NoError(t, err)
	assert.Equal(t, h.username, user)
	assert.Equal(t, h.password, pass)
}

func TestParseBasicAuth(t *testing.T) {
	const validUser = "validUser"
	const validPass = "pass123"
	tests := []struct {
		name        string
		webhookUser string
		webhookPass string
		reqUser     string
		reqPass     string
		expectedErr error
	}{
		{
			name:        "valid basic auth",
			webhookUser: validUser,
			webhookPass: validPass,
			reqUser:     validUser,
			reqPass:     validPass,
			expectedErr: fmt.Errorf("unknown event "), // no event passed, so this is expected
		},
		{
			name:        "no basic auth provided",
			expectedErr: fmt.Errorf("unknown event "), // no event passed, so this is expected
		},
		{
			name:        "invalid basic auth",
			webhookUser: validUser,
			webhookPass: validPass,
			reqUser:     "fakeUser",
			reqPass:     "fakePass",
			expectedErr: ErrBasicAuthVerificationFailed,
		},
	}

	for _, tt := range tests {
		h := Webhook{
			username: tt.webhookUser,
			password: tt.webhookPass,
		}
		body := []byte(`{}`)
		r, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(body))
		assert.NoError(t, err)
		r.SetBasicAuth(tt.reqUser, tt.reqPass)

		p, err := h.Parse(r)
		assert.Equal(t, err, tt.expectedErr)
		assert.Nil(t, p)
	}
}

func newServer(handler http.HandlerFunc) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(virtualDir, handler)
	return httptest.NewServer(mux)
}
