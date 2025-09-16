package azure

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestWebhooks(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name     string
		event    Event
		typ      interface{}
		filename string
		headers  http.Header
	}{
		{
			name:     "build.complete",
			event:    BuildCompleteEventType,
			typ:      BuildCompleteEvent{},
			filename: "../testdata/azuredevops/build.complete.json",
		},
		{
			name:     "git.pullrequest.created",
			event:    GitPullRequestCreatedEventType,
			typ:      GitPullRequestEvent{},
			filename: "../testdata/azuredevops/git.pullrequest.created.json",
		},
		{
			name:     "git.pullrequest.merged",
			event:    GitPullRequestMergedEventType,
			typ:      GitPullRequestEvent{},
			filename: "../testdata/azuredevops/git.pullrequest.merged.json",
		},
		{
			name:     "git.pullrequest.updated",
			event:    GitPullRequestUpdatedEventType,
			typ:      GitPullRequestEvent{},
			filename: "../testdata/azuredevops/git.pullrequest.updated.json",
		},
		{
			name:     "git.push",
			event:    GitPushEventType,
			typ:      GitPushEvent{},
			filename: "../testdata/azuredevops/git.push.json",
		},
	}

	for _, tt := range tests {
		tc := tt
		client := &http.Client{}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			payload, err := os.Open(tc.filename)
			assert.NoError(err)
			defer func() {
				_ = payload.Close()
			}()

			var parseError error
			var results interface{}
			server := newServer(func(w http.ResponseWriter, r *http.Request) {
				results, parseError = hook.Parse(r, tc.event)
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+virtualDir, payload)
			assert.NoError(err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
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
