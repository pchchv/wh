package github

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	// GitHub hook types.
	ForkEvent                                Event = "fork"
	MetaEvent                                Event = "meta"
	TeamEvent                                Event = "team"
	PingEvent                                Event = "ping"
	PushEvent                                Event = "push"
	LabelEvent                               Event = "label"
	WatchEvent                               Event = "watch"
	CreateEvent                              Event = "create"
	DeleteEvent                              Event = "delete"
	GollumEvent                              Event = "gollum"
	PublicEvent                              Event = "public"
	IssuesEvent                              Event = "issues"
	MemberEvent                              Event = "member"
	StatusEvent                              Event = "status"
	ProjectEvent                             Event = "project"
	ReleaseEvent                             Event = "release"
	TeamAddEvent                             Event = "team_add"
	OrgBlockEvent                            Event = "org_block"
	CheckRunEvent                            Event = "check_run"
	DeployKeyEvent                           Event = "deploy_key"
	MilestoneEvent                           Event = "milestone"
	PageBuildEvent                           Event = "page_build"
	CheckSuiteEvent                          Event = "check_suite"
	DeploymentEvent                          Event = "deployment"
	RepositoryEvent                          Event = "repository"
	MembershipEvent                          Event = "membership"
	PullRequestEvent                         Event = "pull_request"
	ProjectCardEvent                         Event = "project_card"
	WorkflowJobEvent                         Event = "workflow_job"
	WorkflowRunEvent                         Event = "workflow_run"
	OrganizationEvent                        Event = "organization"
	IssueCommentEvent                        Event = "issue_comment"
	InstallationEvent                        Event = "installation"
	CommitCommentEvent                       Event = "commit_comment"
	ProjectColumnEvent                       Event = "project_column"
	DependabotAlertEvent                     Event = "dependabot_alert"
	DeploymentStatusEvent                    Event = "deployment_status"
	SecurityAdvisoryEvent                    Event = "security_advisory"
	WorkflowDispatchEvent                    Event = "workflow_dispatch"
	CodeScanningAlertEvent                   Event = "code_scanning_alert"
	PullRequestReviewEvent                   Event = "pull_request_review"
	GitHubAppAuthorizationEvent              Event = "github_app_authorization"
	IntegrationInstallationEvent             Event = "integration_installation"
	PullRequestReviewCommentEvent            Event = "pull_request_review_comment"
	InstallationRepositoriesEvent            Event = "installation_repositories"
	RepositoryVulnerabilityAlertEvent        Event = "repository_vulnerability_alert"
	IntegrationInstallationRepositoriesEvent Event = "integration_installation_repositories"
	// GitHub hook event subtypes.
	NoSubtype     EventSubtype = ""
	TagSubtype    EventSubtype = "tag"
	PullSubtype   EventSubtype = "pull"
	IssueSubtype  EventSubtype = "issues"
	BranchSubtype EventSubtype = "branch"
)

// Options is a namespace var for configuration options.
var Options = WebhookOptions{}

// Event defines a GitHub hook event type.
type Event string

// EventSubtype defines a GitHub Hook Event subtype.
type EventSubtype string

// Option is a configuration option for the webhook.
type Option func(*Webhook) error

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	secret string
}

// New creates and returns a WebHook instance denoted by the Provider type.
func New(options ...Option) (*Webhook, error) {
	hook := new(Webhook)
	for _, opt := range options {
		if err := opt(hook); err != nil {
			return nil, errors.New("Error applying Option")
		}
	}
	return hook, nil
}

// Parse verifies and parses the events specified and returns the payload object or an error.
func (hook Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
	defer func() {
		_, _ = io.Copy(io.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if len(events) == 0 {
		return nil, errors.New("no Event specified to parse")
	}

	if r.Method != http.MethodPost {
		return nil, errors.New("invalid HTTP Method")
	}

	event := r.Header.Get("X-GitHub-Event")
	if event == "" {
		return nil, errors.New("missing X-GitHub-Event Header")
	}
	gitHubEvent := Event(event)

	var found bool
	for _, evt := range events {
		if evt == gitHubEvent {
			found = true
			break
		}
	}

	// event not defined to be parsed
	if !found {
		return nil, errors.New("event not defined to be parsed")
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, errors.New("error parsing payload")
	}

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		signature := r.Header.Get("X-Hub-Signature-256")
		if len(signature) == 0 {
			return nil, errors.New("missing X-Hub-Signature-256 Header")
		}

		signature = strings.TrimPrefix(signature, "sha256=")
		mac := hmac.New(sha256.New, []byte(hook.secret))
		_, _ = mac.Write(payload)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))
		if !hmac.Equal([]byte(signature), []byte(expectedMAC)) {
			return nil, errors.New("HMAC verification failed")
		}
	}

	switch gitHubEvent {
	case CheckRunEvent:
		var pl CheckRunPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case CheckSuiteEvent:
		var pl CheckSuitePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case CommitCommentEvent:
		var pl CommitCommentPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case CreateEvent:
		var pl CreatePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case DeployKeyEvent:
		var pl DeployKeyPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case DeleteEvent:
		var pl DeletePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case DependabotAlertEvent:
		var pl DependabotAlertPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case DeploymentEvent:
		var pl DeploymentPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case DeploymentStatusEvent:
		var pl DeploymentStatusPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case ForkEvent:
		var pl ForkPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case GollumEvent:
		var pl GollumPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case InstallationEvent, IntegrationInstallationEvent:
		var pl InstallationPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case InstallationRepositoriesEvent, IntegrationInstallationRepositoriesEvent:
		var pl InstallationRepositoriesPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssueCommentEvent:
		var pl IssueCommentPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssuesEvent:
		var pl IssuesPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case LabelEvent:
		var pl LabelPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case MemberEvent:
		var pl MemberPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case MembershipEvent:
		var pl MembershipPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case MetaEvent:
		var pl MetaPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case MilestoneEvent:
		var pl MilestonePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case OrganizationEvent:
		var pl OrganizationPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case OrgBlockEvent:
		var pl OrgBlockPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PageBuildEvent:
		var pl PageBuildPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PingEvent:
		var pl PingPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case ProjectCardEvent:
		var pl ProjectCardPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case ProjectColumnEvent:
		var pl ProjectColumnPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case ProjectEvent:
		var pl ProjectPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PublicEvent:
		var pl PublicPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestEvent:
		var pl PullRequestPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestReviewEvent:
		var pl PullRequestReviewPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestReviewCommentEvent:
		var pl PullRequestReviewCommentPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PushEvent:
		var pl PushPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case ReleaseEvent:
		var pl ReleasePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryEvent:
		var pl RepositoryPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryVulnerabilityAlertEvent:
		var pl RepositoryVulnerabilityAlertPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case SecurityAdvisoryEvent:
		var pl SecurityAdvisoryPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case StatusEvent:
		var pl StatusPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case TeamEvent:
		var pl TeamPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case TeamAddEvent:
		var pl TeamAddPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case WatchEvent:
		var pl WatchPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case WorkflowDispatchEvent:
		var pl WorkflowDispatchPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case WorkflowJobEvent:
		var pl WorkflowJobPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case WorkflowRunEvent:
		var pl WorkflowRunPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case GitHubAppAuthorizationEvent:
		var pl GitHubAppAuthorizationPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case CodeScanningAlertEvent:
		var pl CodeScanningAlertPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	default:
		return nil, fmt.Errorf("unknown event %s", gitHubEvent)
	}
}

// WebhookOptions is a namespace for configuration option methods.
type WebhookOptions struct{}

// Secret registers the GitHub secret.
func (WebhookOptions) Secret(secret string) Option {
	return func(hook *Webhook) error {
		hook.secret = secret
		return nil
	}
}
