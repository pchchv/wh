package bitbucket

const (
	// Bitbucket hook types
	RepoPushEvent                  Event = "repo:push"
	RepoForkEvent                  Event = "repo:fork"
	RepoUpdatedEvent               Event = "repo:updated"
	IssueCreatedEvent              Event = "issue:created"
	IssueUpdatedEvent              Event = "issue:updated"
	PullRequestMergedEvent         Event = "pullrequest:fulfilled"
	PullRequestCreatedEvent        Event = "pullrequest:created"
	PullRequestUpdatedEvent        Event = "pullrequest:updated"
	PullRequestApprovedEvent       Event = "pullrequest:approved"
	IssueCommentCreatedEvent       Event = "issue:comment_created"
	PullRequestDeclinedEvent       Event = "pullrequest:rejected"
	PullRequestUnapprovedEvent     Event = "pullrequest:unapproved"
	RepoCommitStatusCreatedEvent   Event = "repo:commit_status_created"
	RepoCommitStatusUpdatedEvent   Event = "repo:commit_status_updated"
	RepoCommitCommentCreatedEvent  Event = "repo:commit_comment_created"
	PullRequestCommentCreatedEvent Event = "pullrequest:comment_created"
	PullRequestCommentUpdatedEvent Event = "pullrequest:comment_updated"
	PullRequestCommentDeletedEvent Event = "pullrequest:comment_deleted"
)

// Event defines a Bitbucket hook event type.
type Event string

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	uuid string
}
