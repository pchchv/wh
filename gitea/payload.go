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

// Issue represents an issue in a repository.
type Issue struct {
	ID               int64            `json:"id"`
	Index            int64            `json:"number"`
	OriginalAuthorID int64            `json:"original_author_id"`
	Comments         int              `json:"comments"`
	Poster           *User            `json:"user"`
	URL              string           `json:"url"`
	Ref              string           `json:"ref"`
	Body             string           `json:"body"`
	Title            string           `json:"title"`
	HTMLURL          string           `json:"html_url"`
	OriginalAuthor   string           `json:"original_author"`
	IsLocked         bool             `json:"is_locked"`
	Labels           []*Label         `json:"labels"`
	Assignees        []*User          `json:"assignees"`
	Milestone        *Milestone       `json:"milestone"`
	Assignee         *User            `json:"assignee"`
	State            StateType        `json:"state"`
	Created          time.Time        `json:"created_at"`
	Updated          time.Time        `json:"updated_at"`
	Closed           *time.Time       `json:"closed_at"`
	Deadline         *time.Time       `json:"due_date"`
	PullRequest      *PullRequestMeta `json:"pull_request"`
	Repo             *RepositoryMeta  `json:"repository"`
}

// Repository represents a repository.
type Repository struct {
	ID                        int64            `json:"id"`
	Size                      int              `json:"size"`
	Stars                     int              `json:"stars_count"`
	Forks                     int              `json:"forks_count"`
	Watchers                  int              `json:"watchers_count"`
	Releases                  int              `json:"release_counter"`
	OpenPulls                 int              `json:"open_pr_counter"`
	OpenIssues                int              `json:"open_issues_count"`
	Owner                     *User            `json:"owner"`
	Name                      string           `json:"name"`
	FullName                  string           `json:"full_name"`
	Description               string           `json:"description"`
	IgnoreWhitespaceConflicts bool             `json:"ignore_whitespace_conflicts"`
	AllowRebaseMerge          bool             `json:"allow_rebase_explicit"`
	HasPullRequests           bool             `json:"has_pull_requests"`
	AllowRebase               bool             `json:"allow_rebase"`
	AllowSquash               bool             `json:"allow_squash_merge"`
	HasProjects               bool             `json:"has_projects"`
	AllowMerge                bool             `json:"allow_merge_commits"`
	HasIssues                 bool             `json:"has_issues"`
	Template                  bool             `json:"template"`
	Archived                  bool             `json:"archived"`
	Internal                  bool             `json:"internal"`
	Private                   bool             `json:"private"`
	HasWiki                   bool             `json:"has_wiki"`
	Mirror                    bool             `json:"mirror"`
	Empty                     bool             `json:"empty"`
	Fork                      bool             `json:"fork"`
	SSHURL                    string           `json:"ssh_url"`
	Website                   string           `json:"website"`
	HTMLURL                   string           `json:"html_url"`
	CloneURL                  string           `json:"clone_url"`
	AvatarURL                 string           `json:"avatar_url"`
	OriginalURL               string           `json:"original_url"`
	DefaultBranch             string           `json:"default_branch"`
	MirrorInterval            string           `json:"mirror_interval"`
	DefaultMergeStyle         string           `json:"default_merge_style"`
	InternalTracker           *InternalTracker `json:"internal_tracker,omitempty"`
	ExternalTracker           *ExternalTracker `json:"external_tracker,omitempty"`
	ExternalWiki              *ExternalWiki    `json:"external_wiki,omitempty"`
	Permissions               *Permission      `json:"permissions,omitempty"`
	Parent                    *Repository      `json:"parent"`
	Created                   time.Time        `json:"created_at"`
	Updated                   time.Time        `json:"updated_at"`
}

// Comment represents a comment on a commit or issue.
type Comment struct {
	ID               int64     `json:"id"`
	OriginalAuthorID int64     `json:"original_author_id"`
	OriginalAuthor   string    `json:"original_author"`
	IssueURL         string    `json:"issue_url"`
	HTMLURL          string    `json:"html_url"`
	PRURL            string    `json:"pull_request_url"`
	Body             string    `json:"body"`
	Poster           *User     `json:"user"`
	Created          time.Time `json:"created_at"`
	Updated          time.Time `json:"updated_at"`
}

// Release represents a repository release.
type Release struct {
	ID           int64         `json:"id"`
	URL          string        `json:"url"`
	Note         string        `json:"body"`
	Title        string        `json:"name"`
	ZipURL       string        `json:"zipball_url"`
	TarURL       string        `json:"tarball_url"`
	Target       string        `json:"target_commitish"`
	HTMLURL      string        `json:"html_url"`
	TagName      string        `json:"tag_name"`
	IsDraft      bool          `json:"draft"`
	IsPrerelease bool          `json:"prerelease"`
	Attachments  []*Attachment `json:"assets"`
	PublishedAt  time.Time     `json:"published_at"`
	CreatedAt    time.Time     `json:"created_at"`
	Publisher    *User         `json:"author"`
}

// Attachment a generic attachment.
type Attachment struct {
	ID            int64     `json:"id"`
	Size          int64     `json:"size"`
	DownloadCount int64     `json:"download_count"`
	Name          string    `json:"name"`
	UUID          string    `json:"uuid"`
	DownloadURL   string    `json:"browser_download_url"`
	Created       time.Time `json:"created_at"`
}

// PullRequest represents a pull request.
type PullRequest struct {
	ID             int64         `json:"id"`
	Index          int64         `json:"number"`
	Comments       int           `json:"comments"`
	MergedCommitID *string       `json:"merge_commit_sha"`
	MergeBase      string        `json:"merge_base"`
	PatchURL       string        `json:"patch_url"`
	DiffURL        string        `json:"diff_url"`
	HTMLURL        string        `json:"html_url"`
	Title          string        `json:"title"`
	Body           string        `json:"body"`
	URL            string        `json:"url"`
	IsLocked       bool          `json:"is_locked"`
	Mergeable      bool          `json:"mergeable"`
	HasMerged      bool          `json:"merged"`
	Deadline       *time.Time    `json:"due_date"`
	Created        *time.Time    `json:"created_at"`
	Updated        *time.Time    `json:"updated_at"`
	Merged         *time.Time    `json:"merged_at"`
	Closed         *time.Time    `json:"closed_at"`
	Poster         *User         `json:"user"`
	Assignee       *User         `json:"assignee"`
	MergedBy       *User         `json:"merged_by"`
	Assignees      []*User       `json:"assignees"`
	Milestone      *Milestone    `json:"milestone"`
	Labels         []*Label      `json:"labels"`
	State          StateType     `json:"state"`
	Base           *PRBranchInfo `json:"base"`
	Head           *PRBranchInfo `json:"head"`
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

// PRBranchInfo information about a branch.
type PRBranchInfo struct {
	Sha        string      `json:"sha"`
	Ref        string      `json:"ref"`
	Name       string      `json:"label"`
	RepoID     int64       `json:"repo_id"`
	Repository *Repository `json:"repo"`
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
