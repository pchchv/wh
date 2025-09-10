package gogs

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	client "github.com/gogits/go-gogs-client"
)

const (
	// Gogs hook types
	PushEvent         Event = "push"
	ForkEvent         Event = "fork"
	CreateEvent       Event = "create"
	DeleteEvent       Event = "delete"
	IssuesEvent       Event = "issues"
	ReleaseEvent      Event = "release"
	PullRequestEvent  Event = "pull_request"
	IssueCommentEvent Event = "issue_comment"
)

// Event defines a Gogs hook event type.
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
	secret string
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

	event := r.Header.Get("X-Gogs-Event")
	if len(event) == 0 {
		return nil, errors.New("missing X-Gogs-Event Header")
	}

	var found bool
	gogsEvent := Event(event)
	for _, evt := range events {
		if evt == gogsEvent {
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

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		signature := r.Header.Get("X-Gogs-Signature")
		if len(signature) == 0 {
			return nil, errors.New("missing X-Gogs-Signature Header")
		}

		mac := hmac.New(sha256.New, []byte(hook.secret))
		_, _ = mac.Write(payload)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))
		if !hmac.Equal([]byte(signature), []byte(expectedMAC)) {
			return nil, errors.New("HMAC verification failed")
		}
	}

	switch gogsEvent {
	case CreateEvent:
		var pl client.CreatePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case ReleaseEvent:
		var pl client.ReleasePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PushEvent:
		var pl client.PushPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case DeleteEvent:
		var pl client.DeletePayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case ForkEvent:
		var pl client.ForkPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssuesEvent:
		var pl client.IssuesPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssueCommentEvent:
		var pl client.IssueCommentPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestEvent:
		var pl client.PullRequestPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	default:
		return nil, fmt.Errorf("unknown event %s", gogsEvent)
	}
}
