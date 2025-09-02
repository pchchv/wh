package github

import "errors"

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

// WebhookOptions is a namespace for configuration option methods.
type WebhookOptions struct{}

// Secret registers the GitHub secret.
func (WebhookOptions) Secret(secret string) Option {
	return func(hook *Webhook) error {
		hook.secret = secret
		return nil
	}
}

// Option is a configuration option for the webhook.
type Option func(*Webhook) error
