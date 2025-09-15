// The `azure` package accepts Azure DevOps Server webhooks.
package azure

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	// Azure DevOps Server hook types.
	GitPushEventType               Event = "git.push"
	BuildCompleteEventType         Event = "build.complete"
	GitPullRequestMergedEventType  Event = "git.pullrequest.merged"
	GitPullRequestCreatedEventType Event = "git.pullrequest.created"
	GitPullRequestUpdatedEventType Event = "git.pullrequest.updated"
)

var (
	// Parse error.
	ErrParsingPayload = errors.New("error parsing payload")
	// Options is a namespace var for configuration options.
	Options = WebhookOptions{}
)

// Event defines an Azure DevOps server hook event type.
type Event string

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	username string
	password string
}

// Parse verifies and parses the events specified and returns the payload object or an error.
func (hook Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
	defer func() {
		_, _ = io.Copy(io.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if !hook.verifyBasicAuth(r) {
		return nil, errors.New("basic auth verification failed")
	}

	if r.Method != http.MethodPost {
		return nil, errors.New("invalid HTTP Method")
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	var pl BasicEvent
	err = json.Unmarshal([]byte(payload), &pl)
	if err != nil {
		return nil, ErrParsingPayload
	}

	switch pl.EventType {
	case GitPushEventType:
		var fpl GitPushEvent
		err = json.Unmarshal([]byte(payload), &fpl)
		return fpl, err
	case GitPullRequestCreatedEventType, GitPullRequestMergedEventType, GitPullRequestUpdatedEventType:
		var fpl GitPullRequestEvent
		err = json.Unmarshal([]byte(payload), &fpl)
		return fpl, err
	case BuildCompleteEventType:
		var fpl BuildCompleteEvent
		err = json.Unmarshal([]byte(payload), &fpl)
		return fpl, err
	default:
		return nil, fmt.Errorf("unknown event %s", pl.EventType)
	}
}

func (hook Webhook) verifyBasicAuth(r *http.Request) bool {
	if hook.username == "" && hook.password == "" {
		// skip validation if username or password was not provided
		return true
	}

	username, password, ok := r.BasicAuth()
	return ok && username == hook.username && password == hook.password
}

// Option is a configuration option for the webhook.
type Option func(*Webhook) error

// New creates and returns a WebHook instance.
func New(options ...Option) (*Webhook, error) {
	hook := new(Webhook)
	for _, opt := range options {
		if err := opt(hook); err != nil {
			return nil, errors.New("Error applying Option")
		}
	}
	return hook, nil
}

// WebhookOptions is a namespace for configuration option methods.
type WebhookOptions struct{}

// BasicAuth verifies payload using basic auth
func (WebhookOptions) BasicAuth(username, password string) Option {
	return func(hook *Webhook) error {
		hook.username = username
		hook.password = password
		return nil
	}
}
