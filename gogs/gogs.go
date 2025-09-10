package gogs

const (
	// Gogs hook types
	PushEvent         Event = "push"
	ForkEvent         Event = "fork"
	CreateEvent       Event = "create"
	DeleteEvent       Event = "delete"
	IssuesEvent       Event = "issues"
	ReleaseEvent      Event = "release"
	PullRequestEvent  Event = "pull_request"
	IssueCommentEvent Event = "issue_comment"
)

// Event defines a Gogs hook event type.
type Event string

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	secret string
}
