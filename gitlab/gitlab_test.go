package gitlab

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
	if hook, err = New(Options.Secret("sampleToken!")); err != nil {
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
			name:    "BadNoEventHeader",
			event:   PushEvents,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{},
		},
		{
			name:    "UnsubscribedEvent",
			event:   PushEvents,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Gitlab-Event": []string{"noneexistant_event"},
			},
		},
		{
			name:    "BadBody",
			event:   PushEvents,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Gitlab-Event": []string{"Push Hook"},
				"X-Gitlab-Token": []string{"sampleToken!"},
			},
		},
		{
			name:    "TokenMismatch",
			event:   PushEvents,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Gitlab-Event": []string{"Push Hook"},
				"X-Gitlab-Token": []string{"badsampleToken!!"},
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
			name:     "PushEvent",
			event:    PushEvents,
			typ:      PushEventPayload{},
			filename: "./testdata/push-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Push Hook"},
			},
		},
		{
			name:     "TagEvent",
			event:    TagEvents,
			typ:      TagEventPayload{},
			filename: "./testdata/tag-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Tag Push Hook"},
			},
		},
		{
			name:     "IssueEvent",
			event:    IssuesEvents,
			typ:      IssueEventPayload{},
			filename: "./testdata/issue-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Issue Hook"},
			},
		},
		{
			name:     "ConfidentialIssueEvent",
			event:    ConfidentialIssuesEvents,
			typ:      ConfidentialIssueEventPayload{},
			filename: "./testdata/confidential-issue-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Confidential Issue Hook"},
			},
		},
		{
			name:     "CommentCommitEvent",
			event:    CommentEvents,
			typ:      CommentEventPayload{},
			filename: "./testdata/comment-commit-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Note Hook"},
			},
		},
		{
			name:     "ConfidentialCommentCommitEvent",
			event:    ConfidentialCommentEvents,
			typ:      ConfidentialCommentEventPayload{},
			filename: "./testdata/confidential-comment-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Confidential Note Hook"},
			},
		},
		{
			name:     "CommentMergeRequestEvent",
			event:    CommentEvents,
			typ:      CommentEventPayload{},
			filename: "./testdata/comment-merge-request-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Note Hook"},
			},
		},
		{
			name:     "CommentIssueEvent",
			event:    CommentEvents,
			typ:      CommentEventPayload{},
			filename: "./testdata/comment-issue-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Note Hook"},
			},
		},
		{
			name:     "CommentSnippetEvent",
			event:    CommentEvents,
			typ:      CommentEventPayload{},
			filename: "./testdata/comment-snippet-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Note Hook"},
			},
		},
		{
			name:     "MergeRequestEvent",
			event:    MergeRequestEvents,
			typ:      MergeRequestEventPayload{},
			filename: "./testdata/merge-request-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Merge Request Hook"},
			},
		},
		{
			name:     "WikipageEvent",
			event:    WikiPageEvents,
			typ:      WikiPageEventPayload{},
			filename: "./testdata/wikipage-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Wiki Page Hook"},
			},
		},
		{
			name:     "PipelineEvent",
			event:    PipelineEvents,
			typ:      PipelineEventPayload{},
			filename: "./testdata/pipeline-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Pipeline Hook"},
			},
		},
		{
			name:     "BuildEvent",
			event:    BuildEvents,
			typ:      BuildEventPayload{},
			filename: "./testdata/build-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Build Hook"},
			},
		},
		{
			name:     "DeploymentEvent",
			event:    DeploymentEvents,
			typ:      DeploymentEventPayload{},
			filename: "./testdata/deployment-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Deployment Hook"},
			},
		},
		{
			name:     "ReleaseEvent",
			event:    ReleaseEvents,
			typ:      ReleaseEventPayload{},
			filename: "./testdata/release-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Release Hook"},
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
			req.Header.Set("X-Gitlab-Token", "sampleToken!")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
}

func TestJobHooks(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name     string
		events   []Event
		typ      interface{}
		filename string
		headers  http.Header
	}{
		{
			name:     "JobEvent",
			events:   []Event{JobEvents, BuildEvents},
			typ:      BuildEventPayload{},
			filename: "./testdata/build-event.json",
			headers: http.Header{
				"X-Gitlab-Event": []string{"Job Hook"},
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
				results, parseError = hook.Parse(r, tc.events...)
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
			assert.NoError(err)
			req.Header = tc.headers
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-Gitlab-Token", "sampleToken!")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
}

func TestSystemHooks(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		name     string
		event    Event
		typ      interface{}
		filename string
	}{
		{
			name:     "PushEvent",
			event:    PushEvents,
			typ:      PushEventPayload{},
			filename: "./testdata/system-push-event.json",
		},
		{
			name:     "TagEvent",
			event:    TagEvents,
			typ:      TagEventPayload{},
			filename: "./testdata/system-tag-event.json",
		},
		{
			name:     "MergeRequestEvent",
			event:    MergeRequestEvents,
			typ:      MergeRequestEventPayload{},
			filename: "./testdata/system-merge-request-event.json",
		},
		{
			name:     "ProjectCreatedEvent",
			event:    SystemHookEvents,
			typ:      ProjectCreatedEventPayload{},
			filename: "./testdata/system-project-created.json",
		},
		{
			name:     "ProjectDestroyedEvent",
			event:    SystemHookEvents,
			typ:      ProjectDestroyedEventPayload{},
			filename: "./testdata/system-project-destroyed.json",
		},
		{
			name:     "ProjectRenamedEvent",
			event:    SystemHookEvents,
			typ:      ProjectRenamedEventPayload{},
			filename: "./testdata/system-project-renamed.json",
		},
		{
			name:     "ProjectTransferredEvent",
			event:    SystemHookEvents,
			typ:      ProjectTransferredEventPayload{},
			filename: "./testdata/system-project-transferred.json",
		},
		{
			name:     "ProjectUpdatedEvent",
			event:    SystemHookEvents,
			typ:      ProjectUpdatedEventPayload{},
			filename: "./testdata/system-project-updated.json",
		},
		{
			name:     "TeamMemberAddedEvent",
			event:    SystemHookEvents,
			typ:      TeamMemberAddedEventPayload{},
			filename: "./testdata/system-team-member-added.json",
		},
		{
			name:     "TeamMemberRemovedEvent",
			event:    SystemHookEvents,
			typ:      TeamMemberRemovedEventPayload{},
			filename: "./testdata/system-team-member-removed.json",
		},
		{
			name:     "TeamMemberUpdatedEvent",
			event:    SystemHookEvents,
			typ:      TeamMemberUpdatedEventPayload{},
			filename: "./testdata/system-team-member-updated.json",
		},
		{
			name:     "UserCreatedEvent",
			event:    SystemHookEvents,
			typ:      UserCreatedEventPayload{},
			filename: "./testdata/system-user-created.json",
		},
		{
			name:     "UserRemovedEvent",
			event:    SystemHookEvents,
			typ:      UserRemovedEventPayload{},
			filename: "./testdata/system-user-removed.json",
		},
		{
			name:     "UserFailedLoginEvent",
			event:    SystemHookEvents,
			typ:      UserFailedLoginEventPayload{},
			filename: "./testdata/system-user-failed-login.json",
		},
		{
			name:     "UserRenamedEvent",
			event:    SystemHookEvents,
			typ:      UserRenamedEventPayload{},
			filename: "./testdata/system-user-renamed.json",
		},
		{
			name:     "KeyAddedEvent",
			event:    SystemHookEvents,
			typ:      KeyAddedEventPayload{},
			filename: "./testdata/system-key-added.json",
		},
		{
			name:     "KeyRemovedEvent",
			event:    SystemHookEvents,
			typ:      KeyRemovedEventPayload{},
			filename: "./testdata/system-key-removed.json",
		},
		{
			name:     "GroupCreatedEvent",
			event:    SystemHookEvents,
			typ:      GroupCreatedEventPayload{},
			filename: "./testdata/system-group-created.json",
		},
		{
			name:     "GroupRemovedEvent",
			event:    SystemHookEvents,
			typ:      GroupRemovedEventPayload{},
			filename: "./testdata/system-group-removed.json",
		},
		{
			name:     "GroupRenamedEvent",
			event:    SystemHookEvents,
			typ:      GroupRenamedEventPayload{},
			filename: "./testdata/system-group-renamed.json",
		},
		{
			name:     "GroupMemberAddedEvent",
			event:    SystemHookEvents,
			typ:      GroupMemberAddedEventPayload{},
			filename: "./testdata/system-group-member-added.json",
		},
		{
			name:     "GroupMemberRemovedEvent",
			event:    SystemHookEvents,
			typ:      GroupMemberRemovedEventPayload{},
			filename: "./testdata/system-group-member-removed.json",
		},
		{
			name:     "GroupMemberUpdatedEvent",
			event:    SystemHookEvents,
			typ:      GroupMemberUpdatedEventPayload{},
			filename: "./testdata/system-group-member-updated.json",
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
				results, parseError = hook.Parse(r, SystemHookEvents, tc.event)
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
			assert.NoError(err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-Gitlab-Token", "sampleToken!")
			req.Header.Set("X-Gitlab-Event", "System Hook")

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
