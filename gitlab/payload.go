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
