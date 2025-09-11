package bitbucket

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	// Bitbucket hook types
	RepoPushEvent                  Event = "repo:push"
	RepoForkEvent                  Event = "repo:fork"
	RepoUpdatedEvent               Event = "repo:updated"
	IssueCreatedEvent              Event = "issue:created"
	IssueUpdatedEvent              Event = "issue:updated"
	PullRequestMergedEvent         Event = "pullrequest:fulfilled"
	PullRequestCreatedEvent        Event = "pullrequest:created"
	PullRequestUpdatedEvent        Event = "pullrequest:updated"
	PullRequestApprovedEvent       Event = "pullrequest:approved"
	IssueCommentCreatedEvent       Event = "issue:comment_created"
	PullRequestDeclinedEvent       Event = "pullrequest:rejected"
	PullRequestUnapprovedEvent     Event = "pullrequest:unapproved"
	RepoCommitStatusCreatedEvent   Event = "repo:commit_status_created"
	RepoCommitStatusUpdatedEvent   Event = "repo:commit_status_updated"
	RepoCommitCommentCreatedEvent  Event = "repo:commit_comment_created"
	PullRequestCommentCreatedEvent Event = "pullrequest:comment_created"
	PullRequestCommentUpdatedEvent Event = "pullrequest:comment_updated"
	PullRequestCommentDeletedEvent Event = "pullrequest:comment_deleted"
)

// Event defines a Bitbucket hook event type.
type Event string

// Option is a configuration option for the webhook.
type Option func(*Webhook) error

// New creates and returns a WebHook instance denoted by the Provider type.
func New(options ...Option) (*Webhook, error) {
	hook := new(Webhook)
	for _, opt := range options {
		if err := opt(hook); err != nil {
			return nil, errors.New("Error applying Option")
		}
	}
	return hook, nil
}

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	uuid string
}

// Parse verifies and parses the events specified and returns the payload object or an error.
func (hook Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
	defer func() {
		_, _ = io.Copy(io.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if len(events) == 0 {
		return nil, errors.New("no Event specified to parse")
	}

	if r.Method != http.MethodPost {
		return nil, errors.New("invalid HTTP Method")
	}

	uuid := r.Header.Get("X-Hook-UUID")
	if hook.uuid != "" && uuid == "" {
		return nil, errors.New("missing X-Hook-UUID Header")
	}

	event := r.Header.Get("X-Event-Key")
	if event == "" {
		return nil, errors.New("missing X-Event-Key Header")
	}

	if len(hook.uuid) > 0 && uuid != hook.uuid {
		return nil, errors.New("UUID verification failed")
	}

	var found bool
	bitbucketEvent := Event(event)
	for _, evt := range events {
		if evt == bitbucketEvent {
			found = true
			break
		}
	}

	// event not defined to be parsed
	if !found {
		return nil, errors.New("event not defined to be parsed")
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, errors.New("error parsing payload")
	}

	switch bitbucketEvent {
	case RepoPushEvent:
		var pl RepoPushPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepoForkEvent:
		var pl RepoForkPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepoUpdatedEvent:
		var pl RepoUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepoCommitCommentCreatedEvent:
		var pl RepoCommitCommentCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepoCommitStatusCreatedEvent:
		var pl RepoCommitStatusCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepoCommitStatusUpdatedEvent:
		var pl RepoCommitStatusUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssueCreatedEvent:
		var pl IssueCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssueUpdatedEvent:
		var pl IssueUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssueCommentCreatedEvent:
		var pl IssueCommentCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCreatedEvent:
		var pl PullRequestCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestUpdatedEvent:
		var pl PullRequestUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestApprovedEvent:
		var pl PullRequestApprovedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestUnapprovedEvent:
		var pl PullRequestUnapprovedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestMergedEvent:
		var pl PullRequestMergedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestDeclinedEvent:
		var pl PullRequestDeclinedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCommentCreatedEvent:
		var pl PullRequestCommentCreatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCommentUpdatedEvent:
		var pl PullRequestCommentUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCommentDeletedEvent:
		var pl PullRequestCommentDeletedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	default:
		return nil, fmt.Errorf("unknown event %s", bitbucketEvent)
	}
}
