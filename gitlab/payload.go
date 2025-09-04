package gitlab

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
