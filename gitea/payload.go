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

// User represents a user.
type User struct {
	ID            int64     `json:"id"`
	Followers     int       `json:"followers_count"`
	Following     int       `json:"following_count"`
	StarredRepos  int       `json:"starred_repos_count"`
	Description   string    `json:"description"`
	Visibility    string    `json:"visibility"`
	AvatarURL     string    `json:"avatar_url"`
	UserName      string    `json:"login"`
	FullName      string    `json:"full_name"`
	Language      string    `json:"language"`
	Location      string    `json:"location"`
	Website       string    `json:"website"`
	Email         string    `json:"email"`
	IsAdmin       bool      `json:"is_admin"`
	IsActive      bool      `json:"active"`
	Restricted    bool      `json:"restricted"`
	ProhibitLogin bool      `json:"prohibit_login"`
	LastLogin     time.Time `json:"last_login,omitempty"`
	Created       time.Time `json:"created,omitempty"`
}

// Milestone milestone is a collection of issues on one repository.
type Milestone struct {
	ID           int64      `json:"id"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	Description  string     `json:"description"`
	Title        string     `json:"title"`
	Created      time.Time  `json:"created_at"`
	Deadline     *time.Time `json:"due_on"`
	Updated      *time.Time `json:"updated_at"`
	Closed       *time.Time `json:"closed_at"`
	State        StateType  `json:"state"`
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

// PayloadCommit represents a commit.
type PayloadCommit struct {
	ID           string                     `json:"id"`
	URL          string                     `json:"url"`
	Message      string                     `json:"message"`
	Modified     []string                   `json:"modified"`
	Removed      []string                   `json:"removed"`
	Added        []string                   `json:"added"`
	Author       *PayloadUser               `json:"author"`
	Committer    *PayloadUser               `json:"committer"`
	Verification *PayloadCommitVerification `json:"verification"`
	Timestamp    time.Time                  `json:"timestamp"`
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
