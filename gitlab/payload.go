package gitlab

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
