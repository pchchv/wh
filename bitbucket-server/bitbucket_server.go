package bitbucket_server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	DiagnosticsPingEvent                 Event = "diagnostics:ping"
	RepositoryForkedEvent                Event = "repo:forked"
	RepositoryModifiedEvent              Event = "repo:modified"
	RepositoryCommentAddedEvent          Event = "repo:comment:added"
	RepositoryCommentEditedEvent         Event = "repo:comment:edited"
	RepositoryCommentDeletedEvent        Event = "repo:comment:deleted"
	RepositoryReferenceChangedEvent      Event = "repo:refs_changed"
	PullRequestOpenedEvent               Event = "pr:opened"
	PullRequestMergedEvent               Event = "pr:merged"
	PullRequestDeletedEvent              Event = "pr:deleted"
	PullRequestModifiedEvent             Event = "pr:modified"
	PullRequestDeclinedEvent             Event = "pr:declined"
	PullRequestCommentAddedEvent         Event = "pr:comment:added"
	PullRequestCommentEditedEvent        Event = "pr:comment:edited"
	PullRequestCommentDeletedEvent       Event = "pr:comment:deleted"
	PullRequestReviewerUpdatedEvent      Event = "pr:reviewer:updated"
	PullRequestReviewerApprovedEvent     Event = "pr:reviewer:approved"
	PullRequestReviewerNeedsWorkEvent    Event = "pr:reviewer:needs_work"
	PullRequestReviewerUnapprovedEvent   Event = "pr:reviewer:unapproved"
	PullRequestFromReferenceUpdatedEvent Event = "pr:from_ref_updated"
)

type Event string

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	secret string
}

func (hook *Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
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

	event := r.Header.Get("X-Event-Key")
	if event == "" {
		return nil, errors.New("missing X-Event-Key Header")
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

	if bitbucketEvent == DiagnosticsPingEvent {
		return DiagnosticsPingPayload{}, nil
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, errors.New("error parsing payload")
	}

	if len(hook.secret) > 0 {
		signature := r.Header.Get("X-Hub-Signature")
		if len(signature) == 0 {
			return nil, errors.New("missing X-Hub-Signature Header")
		}

		mac := hmac.New(sha256.New, []byte(hook.secret))
		_, _ = mac.Write(payload)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))
		if !hmac.Equal([]byte(signature[7:]), []byte(expectedMAC)) {
			return nil, errors.New("HMAC verification failed")
		}
	}

	switch bitbucketEvent {
	case RepositoryReferenceChangedEvent:
		var pl RepositoryReferenceChangedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryModifiedEvent:
		var pl RepositoryModifiedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryForkedEvent:
		var pl RepositoryForkedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryCommentAddedEvent:
		var pl RepositoryCommentAddedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryCommentEditedEvent:
		var pl RepositoryCommentEditedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case RepositoryCommentDeletedEvent:
		var pl RepositoryCommentDeletedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestOpenedEvent:
		var pl PullRequestOpenedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestFromReferenceUpdatedEvent:
		var pl PullRequestFromReferenceUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestModifiedEvent:
		var pl PullRequestModifiedPayload
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
	case PullRequestDeletedEvent:
		var pl PullRequestDeletedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestReviewerUpdatedEvent:
		var pl PullRequestReviewerUpdatedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestReviewerApprovedEvent:
		var pl PullRequestReviewerApprovedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestReviewerUnapprovedEvent:
		var pl PullRequestReviewerUnapprovedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestReviewerNeedsWorkEvent:
		var pl PullRequestReviewerNeedsWorkPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCommentAddedEvent:
		var pl PullRequestCommentAddedPayload
		err = json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PullRequestCommentEditedEvent:
		var pl PullRequestCommentEditedPayload
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
