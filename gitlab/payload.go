package gitlab

// Reviewers contains all of the GitLab reviewers information.
type Reviewers struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}
