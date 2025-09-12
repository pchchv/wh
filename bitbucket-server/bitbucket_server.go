package bitbucket_server

const (
	DiagnosticsPingEvent                 Event = "diagnostics:ping"
	RepositoryForkedEvent                Event = "repo:forked"
	RepositoryModifiedEvent              Event = "repo:modified"
	RepositoryCommentAddedEvent          Event = "repo:comment:added"
	RepositoryCommentEditedEvent         Event = "repo:comment:edited"
	RepositoryCommentDeletedEvent        Event = "repo:comment:deleted"
	RepositoryReferenceChangedEvent      Event = "repo:refs_changed"
	PullRequestOpenedEvent               Event = "pr:opened"
	PullRequestMergedEvent               Event = "pr:merged"
	PullRequestDeletedEvent              Event = "pr:deleted"
	PullRequestModifiedEvent             Event = "pr:modified"
	PullRequestDeclinedEvent             Event = "pr:declined"
	PullRequestCommentAddedEvent         Event = "pr:comment:added"
	PullRequestCommentEditedEvent        Event = "pr:comment:edited"
	PullRequestCommentDeletedEvent       Event = "pr:comment:deleted"
	PullRequestReviewerUpdatedEvent      Event = "pr:reviewer:updated"
	PullRequestReviewerApprovedEvent     Event = "pr:reviewer:approved"
	PullRequestReviewerNeedsWorkEvent    Event = "pr:reviewer:needs_work"
	PullRequestReviewerUnapprovedEvent   Event = "pr:reviewer:unapproved"
	PullRequestFromReferenceUpdatedEvent Event = "pr:from_ref_updated"
)

type Event string
