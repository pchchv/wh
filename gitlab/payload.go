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

// LastCommit contains all of the GitLab last commit information.
type LastCommit struct {
	ID        string     `json:"id"`
	URL       string     `json:"url"`
	Title     string     `json:"title"`
	Message   string     `json:"message"`
	Author    Author     `json:"author"`
	Timestamp customTime `json:"timestamp"`
}

// Changes contains all changes associated with a GitLab issue or MR.
type Changes struct {
	DraftChanges DraftChanges `json:"draft"`
	LabelChanges LabelChanges `json:"labels"`
}

// LabelChanges contains changes in labels associated with a GitLab issue or MR.
type LabelChanges struct {
	Current  []Label `json:"current"`
	Previous []Label `json:"previous"`
}

// ObjectAttributes contains all of the GitLab object attributes information.
type ObjectAttributes struct {
	ID               int64      `json:"id"`
	IID              int64      `json:"iid"`
	StateID          int64      `json:"state_id"`
	Duration         int64      `json:"duration"`
	AuthorID         int64      `json:"author_id"`
	ProjectID        int64      `json:"project_id"`
	AssigneeID       int64      `json:"assignee_id"`
	NoteableID       int64      `json:"noteable_id"`
	MilestoneID      int64      `json:"milestone_id"`
	UpdatedByID      int64      `json:"updated_by_id"`
	TimeEstimate     int64      `json:"time_estimate"`
	TotalTimeSpent   int64      `json:"total_time_spent"`
	LastEditedByID   int64      `json:"last_edited_by_id"`
	SourceProjectID  int64      `json:"source_project_id"`
	TargetProjectID  int64      `json:"target_project_id"`
	RelativePosition int64      `json:"relative_position"`
	AssigneeIDS      []int64    `json:"assignee_ids"`
	ReviewerIDs      []int64    `json:"reviewer_ids"`
	URL              string     `json:"url"`
	SHA              string     `json:"sha"`
	Ref              string     `json:"ref"`
	Note             string     `json:"note"`
	Slug             string     `json:"slug"`
	Title            string     `json:"title"`
	State            string     `json:"state"`
	Action           string     `json:"action"`
	Format           string     `json:"format"`
	Status           string     `json:"status"`
	Content          string     `json:"content"`
	Message          string     `json:"message"`
	LineCode         string     `json:"line_code"`
	CommitID         string     `json:"commit_id"`
	StCommits        string     `json:"st_commits"`
	BeforeSHA        string     `json:"before_sha"`
	BranchName       string     `json:"branch_name"`
	Description      string     `json:"description"`
	MergeStatus      string     `json:"merge_status"`
	TargetBranch     string     `json:"target_branch"`
	SourceBranch     string     `json:"source_branch"`
	NotebookType     string     `json:"noteable_type"`
	Stages           []string   `json:"stages"`
	StDiffs          []StDiff   `json:"st_diffs"`
	Tag              bool       `json:"tag"`
	Draft            bool       `json:"draft"`
	System           bool       `json:"system"`
	Confidential     bool       `json:"confidential"`
	WorkInProgress   bool       `json:"work_in_progress"`
	DiscussionLocked bool       `json:"discussion_locked"`
	At               customTime `json:"attachment"`
	DueDate          customTime `json:"due_date"`
	CreatedAt        customTime `json:"created_at"`
	UpdatedAt        customTime `json:"updated_at"`
	LastEditedAt     customTime `json:"last_edited_at"`
	Source           Source     `json:"source"`
	Target           Target     `json:"target"`
	Position         Position   `json:"position"`
	Assignee         Assignee   `json:"assignee"`
	LastCommit       LastCommit `json:"last_commit"`
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

// UserCreatedEventPayload contains the information about GitLab's user created event.
type UserCreatedEventPayload struct {
	UserID    int64      `json:"user_id"`
	CreatedAt customTime `json:"created_at"`
	UpdatedAt customTime `json:"updated_at"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	EventName string     `json:"event_name"`
}

// UserRemovedEventPayload contains the information about GitLab's user removed event.
type UserRemovedEventPayload struct {
	UserID    int64      `json:"user_id"`
	CreatedAt customTime `json:"created_at"`
	UpdatedAt customTime `json:"updated_at"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	EventName string     `json:"event_name"`
}

// UserFailedLoginEventPayload contains the information about GitLab's user login failed event.
type UserFailedLoginEventPayload struct {
	UserID    int64      `json:"user_id"`
	CreatedAt customTime `json:"created_at"`
	UpdatedAt customTime `json:"updated_at"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	State     string     `json:"state"`
	Username  string     `json:"username"`
	EventName string     `json:"event_name"`
}

// UserRenamedEventPayload contains the information about GitLab's user renamed event.
type UserRenamedEventPayload struct {
	UserID      int64      `json:"user_id"`
	CreatedAt   customTime `json:"created_at"`
	UpdatedAt   customTime `json:"updated_at"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Username    string     `json:"username"`
	EventName   string     `json:"event_name"`
	OldUsername string     `json:"old_username"`
}

// TeamMemberAddedEventPayload contains the information about GitLab's new team member event.
type TeamMemberAddedEventPayload struct {
	UserID                   int64      `json:"user_id"`
	ProjectID                int64      `json:"project_id"`
	CreatedAt                customTime `json:"created_at"`
	UpdatedAt                customTime `json:"updated_at"`
	UserName                 string     `json:"user_name"`
	UserEmail                string     `json:"user_email"`
	EventName                string     `json:"event_name"`
	AccessLevel              string     `json:"access_level"`
	ProjectName              string     `json:"project_name"`
	ProjectPath              string     `json:"project_path"`
	UserUsername             string     `json:"user_username"`
	ProjectVisibility        string     `json:"project_visibility"`
	ProjectPathWithNamespace string     `json:"project_path_with_namespace"`
}

// TeamMemberRemovedEventPayload contains the information about GitLab's team member removed event.
type TeamMemberRemovedEventPayload struct {
	UserID                   int64      `json:"user_id"`
	ProjectID                int64      `json:"project_id"`
	CreatedAt                customTime `json:"created_at"`
	UpdatedAt                customTime `json:"updated_at"`
	UserName                 string     `json:"user_name"`
	UserEmail                string     `json:"user_email"`
	EventName                string     `json:"event_name"`
	AccessLevel              string     `json:"access_level"`
	ProjectName              string     `json:"project_name"`
	ProjectPath              string     `json:"project_path"`
	UserUsername             string     `json:"user_username"`
	ProjectVisibility        string     `json:"project_visibility"`
	ProjectPathWithNamespace string     `json:"project_path_with_namespace"`
}

// TeamMemberUpdatedEventPayload contains the information about GitLab's team member updated event.
type TeamMemberUpdatedEventPayload struct {
	UserID                   int64      `json:"user_id"`
	ProjectID                int64      `json:"project_id"`
	CreatedAt                customTime `json:"created_at"`
	UpdatedAt                customTime `json:"updated_at"`
	UserName                 string     `json:"user_name"`
	UserEmail                string     `json:"user_email"`
	EventName                string     `json:"event_name"`
	AccessLevel              string     `json:"access_level"`
	ProjectName              string     `json:"project_name"`
	ProjectPath              string     `json:"project_path"`
	UserUsername             string     `json:"user_username"`
	ProjectVisibility        string     `json:"project_visibility"`
	ProjectPathWithNamespace string     `json:"project_path_with_namespace"`
}

// ReleaseEventPayload contains the information about GitLab's release event.
type ReleaseEventPayload struct {
	ID          int        `json:"id"`
	CreatedAt   customTime `json:"created_at"`
	ReleasedAt  customTime `json:"released_at"`
	Tag         string     `json:"tag"`
	URL         string     `json:"url"`
	Name        string     `json:"name"`
	Action      string     `json:"action"`
	ObjectKind  string     `json:"object_kind"`
	Description string     `json:"description"`
	Assets      Assets     `json:"assets"`
	Project     Project    `json:"project"`
}

// GroupCreatedEventPayload contains the information about GitLab's group created event.
type GroupCreatedEventPayload struct {
	CreatedAt customTime `json:"created_at"`
	UpdatedAt customTime `json:"updated_at"`
	EventName string     `json:"event_name"`
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	GroupID   int64      `json:"group_id"`
}

// GroupRemovedEventPayload contains the information about GitLab's group removed event.
type GroupRemovedEventPayload struct {
	CreatedAt customTime `json:"created_at"`
	UpdatedAt customTime `json:"updated_at"`
	EventName string     `json:"event_name"`
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	GroupID   int64      `json:"group_id"`
}

// GroupRenamedEventPayload contains the information about GitLab's group renamed event.
type GroupRenamedEventPayload struct {
	CreatedAt   customTime `json:"created_at"`
	UpdatedAt   customTime `json:"updated_at"`
	OldFullPath string     `json:"old_full_path"`
	EventName   string     `json:"event_name"`
	FullPath    string     `json:"full_path"`
	Name        string     `json:"name"`
	Path        string     `json:"path"`
	OldPath     string     `json:"old_path"`
	GroupID     int64      `json:"group_id"`
}

// GroupMemberAddedEventPayload contains the information about GitLab's new group member event.
type GroupMemberAddedEventPayload struct {
	CreatedAt    customTime `json:"created_at"`
	UpdatedAt    customTime `json:"updated_at"`
	UserName     string     `json:"user_name"`
	EventName    string     `json:"event_name"`
	GroupName    string     `json:"group_name"`
	GroupPath    string     `json:"group_path"`
	UserEmail    string     `json:"user_email"`
	GroupAccess  string     `json:"group_access"`
	UserUsername string     `json:"user_username"`
	GroupID      int64      `json:"group_id"`
	UserID       int64      `json:"user_id"`
}

// GroupMemberRemovedEventPayload contains the information about GitLab's group member removed event.
type GroupMemberRemovedEventPayload struct {
	CreatedAt    customTime `json:"created_at"`
	UpdatedAt    customTime `json:"updated_at"`
	EventName    string     `json:"event_name"`
	GroupAccess  string     `json:"group_access"`
	GroupName    string     `json:"group_name"`
	GroupPath    string     `json:"group_path"`
	UserEmail    string     `json:"user_email"`
	UserName     string     `json:"user_name"`
	UserUsername string     `json:"user_username"`
	UserID       int64      `json:"user_id"`
	GroupID      int64      `json:"group_id"`
}

// GroupMemberUpdatedEventPayload contains the information about GitLab's group member updated event.
type GroupMemberUpdatedEventPayload struct {
	CreatedAt    customTime `json:"created_at"`
	UpdatedAt    customTime `json:"updated_at"`
	EventName    string     `json:"event_name"`
	GroupAccess  string     `json:"group_access"`
	GroupName    string     `json:"group_name"`
	GroupPath    string     `json:"group_path"`
	UserEmail    string     `json:"user_email"`
	UserName     string     `json:"user_name"`
	UserUsername string     `json:"user_username"`
	UserID       int64      `json:"user_id"`
	GroupID      int64      `json:"group_id"`
}

// KeyAddedEventPayload contains the information about GitLab's key added event.
type KeyAddedEventPayload struct {
	UpdatedAt customTime `json:"updated_at"`
	EventName string     `json:"event_name"`
	CreatedAt string     `json:"created_at"`
	Username  string     `json:"username"`
	Key       string     `json:"key"`
	Id        int64      `json:"id"`
}

// KeyRemovedEventPayload contains the information about GitLab's key removed event.
type KeyRemovedEventPayload struct {
	UpdatedAt customTime `json:"updated_at"`
	EventName string     `json:"event_name"`
	CreatedAt string     `json:"created_at"`
	Username  string     `json:"username"`
	Key       string     `json:"key"`
	Id        int64      `json:"id"`
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
