package gitlab

// Event defines a GitLab hook event type by the X-Gitlab-Event Header.
type Event string

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	secretHash []byte
}
