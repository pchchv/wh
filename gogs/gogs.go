package gogs

// Event defines a Gogs hook event type.
type Event string

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	secret string
}
