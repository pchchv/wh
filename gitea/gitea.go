package gitea

const (
	// Gitea hook types.
	ForkEvent                 Event = "fork"
	PushEvent                 Event = "push"
	CreateEvent               Event = "create"
	DeleteEvent               Event = "delete"
	IssuesEvent               Event = "issues"
	ReleaseEvent              Event = "release"
	RepositoryEvent           Event = "repository"
	IssueLabelEvent           Event = "issue_label"
	IssueAssignEvent          Event = "issue_assign"
	PullRequestEvent          Event = "pull_request"
	IssueCommentEvent         Event = "issue_comment"
	IssueMilestoneEvent       Event = "issue_milestone"
	PullRequestSyncEvent      Event = "pull_request_sync"
	PullRequestLabelEvent     Event = "pull_request_label"
	PullRequestAssignEvent    Event = "pull_request_assign"
	PullRequestReviewEvent    Event = "pull_request_review"
	PullRequestCommentEvent   Event = "pull_request_comment"
	PullRequestMilestoneEvent Event = "pull_request_milestone"
)

// Event defines a GitLab hook event type by the X-Gitlab-Event Header.
type Event string

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	secret string
}
