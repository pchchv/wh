package bitbucket

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

const path = "/webhooks"

var hook *Webhook

func TestMain(m *testing.M) {
	// setup
	var err error
	hook, err = New(Options.UUID("MY_UUID"))
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
	// teardown
}

func TestBadRequests(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name    string
		event   Event
		payload io.Reader
		headers http.Header
	}{
		{
			name:    "UUIDMissingEvent",
			event:   RepoPushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Event-Key": []string{"noneexistant_event"},
			},
		},
		{
			name:    "UUIDDoesNotMatchEvent",
			event:   RepoPushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Hook-UUID": []string{"THIS_DOES_NOT_MATCH"},
				"X-Event-Key": []string{"repo:push"},
			},
		},
		{
			name:    "BadNoEventHeader",
			event:   RepoPushEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
			},
		},
		{
			name:    "BadBody",
			event:   RepoPushEvent,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:push"},
			},
		},
		{
			name:    "UnsubscribedEvent",
			event:   RepoPushEvent,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"noneexistant_event"},
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		client := &http.Client{}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var parseError error
			server := newServer(func(w http.ResponseWriter, r *http.Request) {
				_, parseError = hook.Parse(r, tc.event)
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+path, tc.payload)
			assert.NoError(err)
			req.Header = tc.headers
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.Error(parseError)
		})
	}
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
			name:     "RepoPush",
			event:    RepoPushEvent,
			typ:      RepoPushPayload{},
			filename: "./testdata/repo-push.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:push"},
			},
		},
		{
			name:     "RepoFork",
			event:    RepoForkEvent,
			typ:      RepoForkPayload{},
			filename: "./testdata/repo-fork.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:fork"},
			},
		},
		{
			name:     "RepoUpdated",
			event:    RepoUpdatedEvent,
			typ:      RepoUpdatedPayload{},
			filename: "./testdata/repo-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:updated"},
			},
		},
		{
			name:     "RepoCommitCommentCreated",
			event:    RepoCommitCommentCreatedEvent,
			typ:      RepoCommitCommentCreatedPayload{},
			filename: "./testdata/commit-comment-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:commit_comment_created"},
			},
		},
		{
			name:     "RepoCommitStatusCreated",
			event:    RepoCommitStatusCreatedEvent,
			typ:      RepoCommitStatusCreatedPayload{},
			filename: "./testdata/repo-commit-status-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:commit_status_created"},
			},
		},
		{
			name:     "RepoCommitStatusUpdated",
			event:    RepoCommitStatusUpdatedEvent,
			typ:      RepoCommitStatusUpdatedPayload{},
			filename: "./testdata/repo-commit-status-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"repo:commit_status_updated"},
			},
		},
		{
			name:     "IssueCreated",
			event:    IssueCreatedEvent,
			typ:      IssueCreatedPayload{},
			filename: "./testdata/issue-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"issue:created"},
			},
		},
		{
			name:     "IssueUpdated",
			event:    IssueUpdatedEvent,
			typ:      IssueUpdatedPayload{},
			filename: "./testdata/issue-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"issue:updated"},
			},
		},
		{
			name:     "IssueUpdated",
			event:    IssueUpdatedEvent,
			typ:      IssueUpdatedPayload{},
			filename: "./testdata/issue-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"issue:updated"},
			},
		},
		{
			name:     "IssueCommentCreated",
			event:    IssueCommentCreatedEvent,
			typ:      IssueCommentCreatedPayload{},
			filename: "./testdata/issue-comment-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"issue:comment_created"},
			},
		},
		{
			name:     "PullRequestCreated",
			event:    PullRequestCreatedEvent,
			typ:      PullRequestCreatedPayload{},
			filename: "./testdata/pull-request-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:created"},
			},
		},
		{
			name:     "PullRequestUpdated",
			event:    PullRequestUpdatedEvent,
			typ:      PullRequestUpdatedPayload{},
			filename: "./testdata/pull-request-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:updated"},
			},
		},
		{
			name:     "PullRequestApproved",
			event:    PullRequestApprovedEvent,
			typ:      PullRequestApprovedPayload{},
			filename: "./testdata/pull-request-approved.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:approved"},
			},
		},
		{
			name:     "PullRequestApprovalRemoved",
			event:    PullRequestUnapprovedEvent,
			typ:      PullRequestUnapprovedPayload{},
			filename: "./testdata/pull-request-approval-removed.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:unapproved"},
			},
		},
		{
			name:     "PullRequestMerged",
			event:    PullRequestMergedEvent,
			typ:      PullRequestMergedPayload{},
			filename: "./testdata/pull-request-merged.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:fulfilled"},
			},
		},
		{
			name:     "PullRequestDeclined",
			event:    PullRequestDeclinedEvent,
			typ:      PullRequestDeclinedPayload{},
			filename: "./testdata/pull-request-declined.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:rejected"},
			},
		},
		{
			name:     "PullRequestCommentCreated",
			event:    PullRequestCommentCreatedEvent,
			typ:      PullRequestCommentCreatedPayload{},
			filename: "./testdata/pull-request-comment-created.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:comment_created"},
			},
		},
		{
			name:     "PullRequestCommentUpdated",
			event:    PullRequestCommentUpdatedEvent,
			typ:      PullRequestCommentUpdatedPayload{},
			filename: "./testdata/pull-request-comment-updated.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:comment_updated"},
			},
		},
		{
			name:     "PullRequestCommentDeleted",
			event:    PullRequestCommentDeletedEvent,
			typ:      PullRequestCommentDeletedPayload{},
			filename: "./testdata/pull-request-comment-deleted.json",
			headers: http.Header{
				"X-Hook-UUID": []string{"MY_UUID"},
				"X-Event-Key": []string{"pullrequest:comment_deleted"},
			},
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
			req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
			assert.NoError(err)
			req.Header = tc.headers
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
}

func newServer(handler http.HandlerFunc) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(path, handler)
	return httptest.NewServer(mux)
}
