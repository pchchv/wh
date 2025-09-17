package gitea

import "time"

// PusherType define the type to push.
type PusherType string

// StateType issue state type.
type StateType string

// HookRepoAction an action that happens to a repo.
type HookRepoAction string

// HookIssueAction FIXME.
type HookIssueAction string

// HookReleaseAction defines hook release action type.
type HookReleaseAction string

// HookIssueCommentAction defines hook issue comment action.
type HookIssueCommentAction string

// Label a label to an issue or a pr.
type Label struct {
	ID          int64  `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

// Permission represents a set of permissions.
type Permission struct {
	Pull  bool `json:"pull"`
	Push  bool `json:"push"`
	Admin bool `json:"admin"`
}

// InternalTracker represents settings for internal tracker.
type InternalTracker struct {
	EnableTimeTracker                bool `json:"enable_time_tracker"`
	EnableIssueDependencies          bool `json:"enable_issue_dependencies"`
	AllowOnlyContributorsToTrackTime bool `json:"allow_only_contributors_to_track_time"`
}

// ExternalTracker represents settings for external tracker.
type ExternalTracker struct {
	ExternalTrackerURL    string `json:"external_tracker_url"`
	ExternalTrackerStyle  string `json:"external_tracker_style"`
	ExternalTrackerFormat string `json:"external_tracker_format"`
}

// ExternalWiki represents setting for external wiki.
type ExternalWiki struct {
	ExternalWikiURL string `json:"external_wiki_url"`
}

// RepositoryMeta basic repository information.
type RepositoryMeta struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	FullName string `json:"full_name"`
}

// PullRequestMeta PR info if an issue is a PR.
type PullRequestMeta struct {
	HasMerged bool       `json:"merged"`
	Merged    *time.Time `json:"merged_at"`
}

// PayloadUser represents the author or committer of a commit.
type PayloadUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

// PayloadCommitVerification represents the GPG verification of a commit.
type PayloadCommitVerification struct {
	Verified  bool         `json:"verified"`
	Reason    string       `json:"reason"`
	Payload   string       `json:"payload"`
	Signature string       `json:"signature"`
	Signer    *PayloadUser `json:"signer"`
}

// ReviewPayload FIXME.
type ReviewPayload struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// ChangesPayload represents the payload information of issue change.
type ChangesPayload struct {
	Ref   *ChangesFromPayload `json:"ref,omitempty"`
	Body  *ChangesFromPayload `json:"body,omitempty"`
	Title *ChangesFromPayload `json:"title,omitempty"`
}

// ChangesFromPayload FIXME.
type ChangesFromPayload struct {
	From string `json:"from"`
}
