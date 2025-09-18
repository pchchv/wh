package github

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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
	hook, err = New(Options.Secret("IsWishesWereHorsesWedAllBeEatingSteak!"))
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
	// teardown
}

func TestWebhooks(t *testing.T) {
	tests := []struct {
		name     string
		event    Event
		typ      interface{}
		filename string
		headers  http.Header
	}{
		{
			name:     "CheckRunEvent",
			event:    CheckRunEvent,
			typ:      CheckRunPayload{},
			filename: "./testdata/check-run.json",
			headers: http.Header{
				"X-Github-Event": []string{"check_run"},
			},
		},
		{
			name:     "CheckSuiteEvent",
			event:    CheckSuiteEvent,
			typ:      CheckSuitePayload{},
			filename: "./testdata/check-suite.json",
			headers: http.Header{
				"X-Github-Event": []string{"check_suite"},
			},
		},
		{
			name:     "CommitCommentEvent",
			event:    CommitCommentEvent,
			typ:      CommitCommentPayload{},
			filename: "./testdata/commit-comment.json",
			headers: http.Header{
				"X-Github-Event": []string{"commit_comment"},
			},
		},
		{
			name:     "CreateEvent",
			event:    CreateEvent,
			typ:      CreatePayload{},
			filename: "./testdata/create.json",
			headers: http.Header{
				"X-Github-Event": []string{"create"},
			},
		},
		{
			name:     "DeleteEvent",
			event:    DeleteEvent,
			typ:      DeletePayload{},
			filename: "./testdata/delete.json",
			headers: http.Header{
				"X-Github-Event": []string{"delete"},
			},
		},
		{
			name:     "DependabotAlertEvent",
			event:    DependabotAlertEvent,
			typ:      DependabotAlertPayload{},
			filename: "./testdata/dependabot_alert.json",
			headers: http.Header{
				"X-Github-Event": []string{"dependabot_alert"},
			},
		},
		{
			name:     "DeployKeyEvent",
			event:    DeployKeyEvent,
			typ:      DeployKeyPayload{},
			filename: "./testdata/deploy_key.json",
			headers: http.Header{
				"X-Github-Event": []string{"deploy_key"},
			},
		},
		{
			name:     "DeploymentEvent",
			event:    DeploymentEvent,
			typ:      DeploymentPayload{},
			filename: "./testdata/deployment.json",
			headers: http.Header{
				"X-Github-Event": []string{"deployment"},
			},
		},
		{
			name:     "DeploymentStatusEvent",
			event:    DeploymentStatusEvent,
			typ:      DeploymentStatusPayload{},
			filename: "./testdata/deployment-status.json",
			headers: http.Header{
				"X-Github-Event": []string{"deployment_status"},
			},
		},
		{
			name:     "ForkEvent",
			event:    ForkEvent,
			typ:      ForkPayload{},
			filename: "./testdata/fork.json",
			headers: http.Header{
				"X-Github-Event": []string{"fork"},
			},
		},
		{
			name:     "GollumEvent",
			event:    GollumEvent,
			typ:      GollumPayload{},
			filename: "./testdata/gollum.json",
			headers: http.Header{
				"X-Github-Event": []string{"gollum"},
			},
		},
		{
			name:     "InstallationEvent",
			event:    InstallationEvent,
			typ:      InstallationPayload{},
			filename: "./testdata/installation.json",
			headers: http.Header{
				"X-Github-Event": []string{"installation"},
			},
		},
		{
			name:     "InstallationRepositoriesEvent",
			event:    InstallationRepositoriesEvent,
			typ:      InstallationRepositoriesPayload{},
			filename: "./testdata/installation-repositories.json",
			headers: http.Header{
				"X-Github-Event": []string{"installation_repositories"},
			},
		},
		{
			name:     "IntegrationInstallationEvent",
			event:    IntegrationInstallationEvent,
			typ:      InstallationPayload{},
			filename: "./testdata/integration-installation.json",
			headers: http.Header{
				"X-Github-Event": []string{"integration_installation"},
			},
		},
		{
			name:     "IntegrationInstallationRepositoriesEvent",
			event:    IntegrationInstallationRepositoriesEvent,
			typ:      InstallationRepositoriesPayload{},
			filename: "./testdata/integration-installation-repositories.json",
			headers: http.Header{
				"X-Github-Event": []string{"integration_installation_repositories"},
			},
		},
		{
			name:     "IssueCommentEvent",
			event:    IssueCommentEvent,
			typ:      IssueCommentPayload{},
			filename: "./testdata/issue-comment.json",
			headers: http.Header{
				"X-Github-Event": []string{"issue_comment"},
			},
		},
		{
			name:     "PullRequestIssueCommentEvent",
			event:    IssueCommentEvent,
			typ:      IssueCommentPayload{},
			filename: "./testdata/pull-request-issue-comment.json",
			headers: http.Header{
				"X-Github-Event": []string{"issue_comment"},
			},
		},
		{
			name:     "IssuesEvent",
			event:    IssuesEvent,
			typ:      IssuesPayload{},
			filename: "./testdata/issues.json",
			headers: http.Header{
				"X-Github-Event": []string{"issues"},
			},
		},
		{
			name:     "LabelEvent",
			event:    LabelEvent,
			typ:      LabelPayload{},
			filename: "./testdata/label.json",
			headers: http.Header{
				"X-Github-Event": []string{"label"},
			},
		},
		{
			name:     "MemberEvent",
			event:    MemberEvent,
			typ:      MemberPayload{},
			filename: "./testdata/member.json",
			headers: http.Header{
				"X-Github-Event": []string{"member"},
			},
		},
		{
			name:     "MembershipEvent",
			event:    MembershipEvent,
			typ:      MembershipPayload{},
			filename: "./testdata/membership.json",
			headers: http.Header{
				"X-Github-Event": []string{"membership"},
			},
		},
		{
			name:     "MilestoneEvent",
			event:    MilestoneEvent,
			typ:      MilestonePayload{},
			filename: "./testdata/milestone.json",
			headers: http.Header{
				"X-Github-Event": []string{"milestone"},
			},
		},
		{
			name:     "OrganizationEvent",
			event:    OrganizationEvent,
			typ:      OrganizationPayload{},
			filename: "./testdata/organization.json",
			headers: http.Header{
				"X-Github-Event": []string{"organization"},
			},
		},
		{
			name:     "OrgBlockEvent",
			event:    OrgBlockEvent,
			typ:      OrgBlockPayload{},
			filename: "./testdata/org-block.json",
			headers: http.Header{
				"X-Github-Event": []string{"org_block"},
			},
		},
		{
			name:     "PageBuildEvent",
			event:    PageBuildEvent,
			typ:      PageBuildPayload{},
			filename: "./testdata/page-build.json",
			headers: http.Header{
				"X-Github-Event": []string{"page_build"},
			},
		},
		{
			name:     "PingEvent",
			event:    PingEvent,
			typ:      PingPayload{},
			filename: "./testdata/ping.json",
			headers: http.Header{
				"X-Github-Event": []string{"ping"},
			},
		},
		{
			name:     "ProjectCardEvent",
			event:    ProjectCardEvent,
			typ:      ProjectCardPayload{},
			filename: "./testdata/project-card.json",
			headers: http.Header{
				"X-Github-Event": []string{"project_card"},
			},
		},
		{
			name:     "ProjectColumnEvent",
			event:    ProjectColumnEvent,
			typ:      ProjectColumnPayload{},
			filename: "./testdata/project-column.json",
			headers: http.Header{
				"X-Github-Event": []string{"project_column"},
			},
		},
		{
			name:     "ProjectEvent",
			event:    ProjectEvent,
			typ:      ProjectPayload{},
			filename: "./testdata/project.json",
			headers: http.Header{
				"X-Github-Event": []string{"project"},
			},
		},
		{
			name:     "PublicEvent",
			event:    PublicEvent,
			typ:      PublicPayload{},
			filename: "./testdata/public.json",
			headers: http.Header{
				"X-Github-Event": []string{"public"},
			},
		},
		{
			name:     "PullRequestEvent",
			event:    PullRequestEvent,
			typ:      PullRequestPayload{},
			filename: "./testdata/pull-request.json",
			headers: http.Header{
				"X-Github-Event": []string{"pull_request"},
			},
		},
		{
			name:     "PullRequestReviewEvent",
			event:    PullRequestReviewEvent,
			typ:      PullRequestReviewPayload{},
			filename: "./testdata/pull-request-review.json",
			headers: http.Header{
				"X-Github-Event": []string{"pull_request_review"},
			},
		},
		{
			name:     "PullRequestReviewCommentEvent",
			event:    PullRequestReviewCommentEvent,
			typ:      PullRequestReviewCommentPayload{},
			filename: "./testdata/pull-request-review-comment.json",
			headers: http.Header{
				"X-Github-Event": []string{"pull_request_review_comment"},
			},
		},
		{
			name:     "PushEvent",
			event:    PushEvent,
			typ:      PushPayload{},
			filename: "./testdata/push.json",
			headers: http.Header{
				"X-Github-Event": []string{"push"},
			},
		},
		{
			name:     "ReleaseEvent",
			event:    ReleaseEvent,
			typ:      ReleasePayload{},
			filename: "./testdata/release.json",
			headers: http.Header{
				"X-Github-Event": []string{"release"},
			},
		},
		{
			name:     "RepositoryEvent",
			event:    RepositoryEvent,
			typ:      RepositoryPayload{},
			filename: "./testdata/repository.json",
			headers: http.Header{
				"X-Github-Event": []string{"repository"},
			},
		},
		{
			name:     "RepositoryEditedEvent",
			event:    RepositoryEvent,
			typ:      RepositoryPayload{},
			filename: "./testdata/repository-edited.json",
			headers: http.Header{
				"X-Github-Event": []string{"repository"},
			},
		},
		{
			name:     "RepositoryVulnerabilityAlertEvent",
			event:    RepositoryVulnerabilityAlertEvent,
			typ:      RepositoryVulnerabilityAlertPayload{},
			filename: "./testdata/repository-vulnerability-alert.json",
			headers: http.Header{
				"X-Github-Event": []string{"repository_vulnerability_alert"},
			},
		},
		{
			name:     "SecurityAdvisoryEvent",
			event:    SecurityAdvisoryEvent,
			typ:      SecurityAdvisoryPayload{},
			filename: "./testdata/security-advisory.json",
			headers: http.Header{
				"X-Github-Event": []string{"security_advisory"},
			},
		},
		{
			name:     "StatusEvent",
			event:    StatusEvent,
			typ:      StatusPayload{},
			filename: "./testdata/status.json",
			headers: http.Header{
				"X-Github-Event": []string{"status"},
			},
		},
		{
			name:     "TeamEvent",
			event:    TeamEvent,
			typ:      TeamPayload{},
			filename: "./testdata/team.json",
			headers: http.Header{
				"X-Github-Event": []string{"team"},
			},
		},
		{
			name:     "TeamAddEvent",
			event:    TeamAddEvent,
			typ:      TeamAddPayload{},
			filename: "./testdata/team-add.json",
			headers: http.Header{
				"X-Github-Event": []string{"team_add"},
			},
		},
		{
			name:     "WatchEvent",
			event:    WatchEvent,
			typ:      WatchPayload{},
			filename: "./testdata/watch.json",
			headers: http.Header{
				"X-Github-Event": []string{"watch"},
			},
		},
		{
			name:     "WorkflowDispatchEvent",
			event:    WorkflowDispatchEvent,
			typ:      WorkflowDispatchPayload{},
			filename: "./testdata/workflow_dispatch.json",
			headers: http.Header{
				"X-Github-Event": []string{"workflow_dispatch"},
			},
		},
		{
			name:     "WorkflowJobEvent",
			event:    WorkflowJobEvent,
			typ:      WorkflowJobPayload{},
			filename: "./testdata/workflow_job.json",
			headers: http.Header{
				"X-Github-Event": []string{"workflow_job"},
			},
		},
		{
			name:     "WorkflowRunEvent",
			event:    WorkflowRunEvent,
			typ:      WorkflowRunPayload{},
			filename: "./testdata/workflow_run.json",
			headers: http.Header{
				"X-Github-Event": []string{"workflow_run"},
			},
		},
		{
			name:     "GitHubAppAuthorizationEvent",
			event:    GitHubAppAuthorizationEvent,
			typ:      GitHubAppAuthorizationPayload{},
			filename: "./testdata/github-app-authorization.json",
			headers: http.Header{
				"X-Github-Event": []string{"github_app_authorization"},
			},
		},
		{
			name:     "CodeScanningAlertEvent",
			event:    CodeScanningAlertEvent,
			typ:      CodeScanningAlertPayload{},
			filename: "./testdata/code_scanning_alert.json",
			headers: http.Header{
				"X-Github-Event": []string{"code_scanning_alert"},
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		client := &http.Client{}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert := require.New(t)
			payload, err := os.ReadFile(tc.filename)
			assert.NoError(err)

			var parseError error
			var results interface{}
			server := newServer(func(w http.ResponseWriter, r *http.Request) {
				results, parseError = hook.Parse(r, tc.event)
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+path, bytes.NewReader(payload))
			assert.NoError(err)
			req.Header = tc.headers
			mac := hmac.New(sha256.New, []byte(hook.secret))
			mac.Write(payload)

			req.Header.Set("X-Hub-Signature-256", "sha256="+hex.EncodeToString(mac.Sum(nil)))
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
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
			event:   CreateEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{},
		},
		{
			name:    "UnsubscribedEvent",
			event:   CreateEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Github-Event": []string{"noneexistant_event"},
			},
		},
		{
			name:    "BadBody",
			event:   CommitCommentEvent,
			payload: bytes.NewBuffer([]byte("")),
			headers: http.Header{
				"X-Github-Event":      []string{"commit_comment"},
				"X-Hub-Signature-256": []string{"sha256=156404ad5f721c53151147f3d3d302329f95a3ab"},
			},
		},
		{
			name:    "BadSignatureLength",
			event:   CommitCommentEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Github-Event":      []string{"commit_comment"},
				"X-Hub-Signature-256": []string{""},
			},
		},
		{
			name:    "BadSignatureMatch",
			event:   CommitCommentEvent,
			payload: bytes.NewBuffer([]byte("{}")),
			headers: http.Header{
				"X-Github-Event":      []string{"commit_comment"},
				"X-Hub-Signature-256": []string{"111"},
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

func newServer(handler http.HandlerFunc) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(path, handler)
	return httptest.NewServer(mux)
}
