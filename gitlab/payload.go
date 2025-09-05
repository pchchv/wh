package gitlab

import (
	"strings"
	"time"
)

// Assets represent artefacts and links associated to a release.
type Assets struct {
	Count   int           `json:"count"`
	Links   []Link        `json:"links"`
	Sources []AssetSource `json:"sources"`
}

// AssetSource represent the download url for an asset.
type AssetSource struct {
	URL    string `json:"url"`
	Format string `json:"format"`
}

// Reviewers contains all of the GitLab reviewers information.
type Reviewers struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

// SystemHookPayload contains the ObjectKind to match with real hook events.
type SystemHookPayload struct {
	EventName  string `json:"event_name"`
	ObjectKind string `json:"object_kind"`
}

// DeploymentEventPayload contains the information for GitLab's triggered when a deployment.
type DeploymentEventPayload struct {
	User           User    `json:"user"`
	Project        Project `json:"project"`
	Status         string  `json:"status"`
	UserUrl        string  `json:"user_url"`
	ShortSha       string  `json:"short_sha"`
	CommitUrl      string  `json:"commit_url"`
	ObjectKind     string  `json:"object_kind"`
	Environment    string  `json:"environment"`
	CommitTitle    string  `json:"commit_title"`
	DeployableUrl  string  `json:"deployable_url"`
	StatusChangeAt string  `json:"status_changed_at"`
	DeployableId   int64   `json:"deployable_id"`
	DeploymentId   int64   `json:"deployment_id"`
}

// Link represent a generic html link.
type Link struct {
	ID       int    `json:"id"`
	External bool   `json:"external"`
	LinkType string `json:"link_type"`
	Name     string `json:"name"`
	URL      string `json:"url"`
}

// Runner represents a runner agent.
type Runner struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	IsShared    bool   `json:"is_shared"`
}

// User contains all of the GitLab user information.
type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	UserName  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
}

// Project contains all of the GitLab project information.
type Project struct {
	ID                int64  `json:"id"`
	VisibilityLevel   int64  `json:"visibility_level"`
	URL               string `json:"url"`
	Name              string `json:"name"`
	WebURL            string `json:"web_url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
	Homepage          string `json:"homepage"`
	Namespace         string `json:"namespace"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Description       string `json:"description"`
	DefaultBranch     string `json:"default_branch"`
	CiConfigPath      string `json:"ci_config_path"`
	PathWithNamespace string `json:"path_with_namespace"`
}

// Wiki contains all of the GitLab wiki information.
type Wiki struct {
	WebURL            string `json:"web_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	DefaultBranch     string `json:"default_branch"`
	PathWithNamespace string `json:"path_with_namespace"`
}

// Position defines a specific location,
// identified by paths line numbers and image coordinates,
// within a specific diff, identified by start,
// head and base commit ids.
//
// Text position will have: new_line and old_line.
// Image position will have: width, height, x, y.
type Position struct {
	BaseSHA      string `json:"base_sha"`
	HeadSHA      string `json:"head_sha"`
	OldPath      string `json:"old_path"`
	NewPath      string `json:"new_path"`
	StartSHA     string `json:"start_sha"`
	PositionType string `json:"position_type"`
	OldLine      int64  `json:"old_line"`
	NewLine      int64  `json:"new_line"`
	Height       int64  `json:"height"`
	Width        int64  `json:"width"`
	X            int64  `json:"x"`
	Y            int64  `json:"y"`
}

// ArtifactsFile contains all of the GitLab artifact information.
type ArtifactsFile struct {
	Size     string `json:"size"`
	Filename string `json:"filename"`
}

// Variable contains pipeline variables.
type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Author contains all of the GitLab author information.
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Repository contains all of the GitLab repository information.
type Repository struct {
	URL             string `json:"url"`
	Name            string `json:"name"`
	Homepage        string `json:"homepage"`
	GitSSHURL       string `json:"git_ssh_url"`
	GitHTTPURL      string `json:"git_http_url"`
	Description     string `json:"description"`
	VisibilityLevel int64  `json:"visibility_level"`
}

// DraftChanges contains the current and previous value of the draft property,
// tells us if draft was toggles.
type DraftChanges struct {
	Current  bool `json:"current"`
	Previous bool `json:"previous"`
}

// Assignee contains all of the GitLab assignee information.
type Assignee struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

// Source contains all of the GitLab source information.
type Source struct {
	URL               string `json:"url"`
	Name              string `json:"name"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	Homepage          string `json:"homepage"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
	Description       string `json:"description"`
	DefaultBranch     string `json:"default_branch"`
	PathWithNamespace string `json:"path_with_namespace"`
	VisibilityLevel   int64  `json:"visibility_level"`
}

// Target contains all of the GitLab target information.
type Target struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
	VisibilityLevel   int64  `json:"visibility_level"`
}

// StDiff contains all of the GitLab diff information.
type StDiff struct {
	Diff        string `json:"diff"`
	NewPath     string `json:"new_path"`
	OldPath     string `json:"old_path"`
	AMode       string `json:"a_mode"`
	BMode       string `json:"b_mode"`
	NewFile     bool   `json:"new_file"`
	RenamedFile bool   `json:"renamed_file"`
	DeletedFile bool   `json:"deleted_file"`
}

// Label contains all of the GitLab label information.
type Label struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Color       string     `json:"color"`
	ProjectID   int64      `json:"project_id"`
	CreatedAt   customTime `json:"created_at"`
	UpdatedAt   customTime `json:"updated_at"`
	Template    bool       `json:"template"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	GroupID     int64      `json:"group_id"`
}

// Commit contains all of the GitLab commit information.
type Commit struct {
	ID        string     `json:"id"`
	URL       string     `json:"url"`
	Title     string     `json:"title"`
	Message   string     `json:"message"`
	Added     []string   `json:"added"`
	Removed   []string   `json:"removed"`
	Modified  []string   `json:"modified"`
	Author    Author     `json:"author"`
	Timestamp customTime `json:"timestamp"`
}

// BuildCommit contains all of the GitLab build commit information.
type BuildCommit struct {
	ID          int64      `json:"id"`
	Duration    float64    `json:"duration"`
	SHA         string     `json:"sha"`
	Status      string     `json:"status"`
	Message     string     `json:"message"`
	AuthorName  string     `json:"author_name"`
	AuthorEmail string     `json:"author_email"`
	StartedAt   customTime `json:"started_at"`
	FinishedAt  customTime `json:"finished_at"`
}

// BuildEventPayload contains the information for GitLab's build status change event.
type BuildEventPayload struct {
	Tag                 bool        `json:"tag"`
	BuildAllowFailure   bool        `json:"build_allow_failure"`
	BuildID             int64       `json:"build_id"`
	ProjectID           int64       `json:"project_id"`
	BuildDuration       float64     `json:"build_duration"`
	BuildQueuedDuration float64     `json:"build_queued_duration"`
	Ref                 string      `json:"ref"`
	SHA                 string      `json:"sha"`
	BuildName           string      `json:"build_name"`
	BeforeSHA           string      `json:"before_sha"`
	ObjectKind          string      `json:"object_kind"`
	BuildStage          string      `json:"build_stage"`
	BuildStatus         string      `json:"build_status"`
	ProjectName         string      `json:"project_name"`
	User                User        `json:"user"`
	Runner              Runner      `json:"runner"`
	Commit              BuildCommit `json:"commit"`
	BuildStartedAt      customTime  `json:"build_started_at"`
	BuildFinishedAt     customTime  `json:"build_finished_at"`
	Repository          Repository  `json:"repository"`
}

// JobEventPayload contains the information for GitLab's Job status change.
type JobEventPayload struct {
	Tag                 bool        `json:"tag"`
	BuildAllowFailure   bool        `json:"build_allow_failure"`
	BuildID             int64       `json:"build_id"`
	ProjectID           int64       `json:"project_id"`
	PipelineID          int64       `json:"pipeline_id"`
	BuildDuration       float64     `json:"build_duration"`
	BuildQueuedDuration float64     `json:"build_queued_duration"`
	Ref                 string      `json:"ref"`
	SHA                 string      `json:"sha"`
	BuildName           string      `json:"build_name"`
	BeforeSHA           string      `json:"before_sha"`
	ObjectKind          string      `json:"object_kind"`
	BuildStage          string      `json:"build_stage"`
	ProjectName         string      `json:"project_name"`
	BuildStatus         string      `json:"build_status"`
	BuildFailureReason  string      `json:"build_failure_reason"`
	User                User        `json:"user"`
	Runner              Runner      `json:"runner"`
	BuildStartedAt      customTime  `json:"build_started_at"`
	BuildFinishedAt     customTime  `json:"build_finished_at"`
	Commit              BuildCommit `json:"commit"`
	Repository          Repository  `json:"repository"`
}

// PushEventPayload contains the information for GitLab's push event.
type PushEventPayload struct {
	UserID            int64      `json:"user_id"`
	ProjectID         int64      `json:"project_id"`
	TotalCommitsCount int64      `json:"total_commits_count"`
	Ref               string     `json:"ref"`
	ObjectKind        string     `json:"object_kind"`
	Before            string     `json:"before"`
	After             string     `json:"after"`
	CheckoutSHA       string     `json:"checkout_sha"`
	UserName          string     `json:"user_name"`
	UserUsername      string     `json:"user_username"`
	UserEmail         string     `json:"user_email"`
	UserAvatar        string     `json:"user_avatar"`
	Project           Project    `json:"project"`
	Repository        Repository `json:"repository"`
	Commits           []Commit   `json:"commits"`
}

// TagEventPayload contains the information for GitLab's tag push event.
type TagEventPayload struct {
	UserID            int64      `json:"user_id"`
	ProjectID         int64      `json:"project_id"`
	TotalCommitsCount int64      `json:"total_commits_count"`
	ObjectKind        string     `json:"object_kind"`
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	CheckoutSHA       string     `json:"checkout_sha"`
	UserName          string     `json:"user_name"`
	UserUsername      string     `json:"user_username"`
	UserAvatar        string     `json:"user_avatar"`
	Project           Project    `json:"project"`
	Repository        Repository `json:"repository"`
	Commits           []Commit   `json:"commits"`
}

// ProjectCreatedEventPayload contains the information about GitLab's project created event.
type ProjectCreatedEventPayload struct {
	ProjectID         int64      `json:"project_id"`
	CreatedAt         customTime `json:"created_at"`
	UpdatedAt         customTime `json:"updated_at"`
	Name              string     `json:"name"`
	Path              string     `json:"path"`
	EventName         string     `json:"event_name"`
	OwnerName         string     `json:"owner_name"`
	OwnerEmail        string     `json:"owner_email"`
	PathWithNamespace string     `json:"path_with_namespace"`
	ProjectVisibility string     `json:"project_visibility"`
	Owners            []Author   `json:"owners"`
}

// ProjectDestroyedEventPayload contains the information about GitLab's project destroyed event.
type ProjectDestroyedEventPayload struct {
	ProjectID         int64      `json:"project_id"`
	CreatedAt         customTime `json:"created_at"`
	UpdatedAt         customTime `json:"updated_at"`
	Name              string     `json:"name"`
	Path              string     `json:"path"`
	EventName         string     `json:"event_name"`
	OwnerName         string     `json:"owner_name"`
	OwnerEmail        string     `json:"owner_email"`
	PathWithNamespace string     `json:"path_with_namespace"`
	ProjectVisibility string     `json:"project_visibility"`
	Owners            []Author   `json:"owners"`
}

// ProjectRenamedEventPayload contains the information about GitLab's project renamed event.
type ProjectRenamedEventPayload struct {
	ProjectID            int64      `json:"project_id"`
	CreatedAt            customTime `json:"created_at"`
	UpdatedAt            customTime `json:"updated_at"`
	EventName            string     `json:"event_name"`
	Name                 string     `json:"name"`
	Path                 string     `json:"path"`
	OwnerName            string     `json:"owner_name"`
	OwnerEmail           string     `json:"owner_email"`
	ProjectVisibility    string     `json:"project_visibility"`
	PathWithNamespace    string     `json:"path_with_namespace"`
	OldPathWithNamespace string     `json:"old_path_with_namespace"`
	Owners               []Author   `json:"owners"`
}

// ProjectTransferredEventPayload contains the information about GitLab's project transferred event.
type ProjectTransferredEventPayload struct {
	ProjectID            int64      `json:"project_id"`
	CreatedAt            customTime `json:"created_at"`
	UpdatedAt            customTime `json:"updated_at"`
	Name                 string     `json:"name"`
	Path                 string     `json:"path"`
	EventName            string     `json:"event_name"`
	OwnerName            string     `json:"owner_name"`
	OwnerEmail           string     `json:"owner_email"`
	PathWithNamespace    string     `json:"path_with_namespace"`
	ProjectVisibility    string     `json:"project_visibility"`
	OldPathWithNamespace string     `json:"old_path_with_namespace"`
	Owners               []Author   `json:"owners"`
}

// ProjectUpdatedEventPayload contains the information about GitLab's project updated event.
type ProjectUpdatedEventPayload struct {
	ProjectID         int64      `json:"project_id"`
	CreatedAt         customTime `json:"created_at"`
	UpdatedAt         customTime `json:"updated_at"`
	Name              string     `json:"name"`
	Path              string     `json:"path"`
	EventName         string     `json:"event_name"`
	OwnerName         string     `json:"owner_name"`
	OwnerEmail        string     `json:"owner_email"`
	PathWithNamespace string     `json:"path_with_namespace"`
	ProjectVisibility string     `json:"project_visibility"`
	Owners            []Author   `json:"owners"`
}

type customTime struct {
	time.Time
}

func (t *customTime) UnmarshalJSON(b []byte) (err error) {
	layout := []string{
		"2006-01-02 15:04:05 MST",
		"2006-01-02 15:04:05 Z07:00",
		"2006-01-02 15:04:05 Z0700",
		time.RFC3339,
	}
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}

	for _, l := range layout {
		t.Time, err = time.Parse(l, s)
		if err == nil {
			break
		}
	}

	return
}
