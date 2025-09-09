package gitlab

import (
	"crypto/sha512"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	// GitLab hook types.
	TagEvents                 Event  = "Tag Push Hook"
	JobEvents                 Event  = "Job Hook"
	PushEvents                Event  = "Push Hook"
	BuildEvents               Event  = "Build Hook"
	IssuesEvents              Event  = "Issue Hook"
	CommentEvents             Event  = "Note Hook"
	ReleaseEvents             Event  = "Release Hook"
	WikiPageEvents            Event  = "Wiki Page Hook"
	PipelineEvents            Event  = "Pipeline Hook"
	DeploymentEvents          Event  = "Deployment Hook"
	SystemHookEvents          Event  = "System Hook"
	MergeRequestEvents        Event  = "Merge Request Hook"
	ConfidentialIssuesEvents  Event  = "Confidential Issue Hook"
	ConfidentialCommentEvents Event  = "Confidential Note Hook"
	objectTag                 string = "tag_push"
	objectPush                string = "push"
	objectBuild               string = "build"
	objectMergeRequest        string = "merge_request"
	eventKeyCreate            string = "key_create"
	eventUserCreate           string = "user_create"
	eventUserRename           string = "user_rename"
	eventKeyDestroy           string = "key_destroy"
	eventUserDestroy          string = "user_destroy"
	eventGroupRename          string = "group_rename"
	eventGroupCreate          string = "group_create"
	eventGroupDestroy         string = "group_destroy"
	eventProjectCreate        string = "project_create"
	eventProjectRename        string = "project_rename"
	eventProjectUpdate        string = "project_update"
	eventUserAddToTeam        string = "user_add_to_team"
	eventProjectDestroy       string = "project_destroy"
	eventUserAddToGroup       string = "user_add_to_group"
	eventUserFailedLogin      string = "user_failed_login"
	eventProjectTransfer      string = "project_transfer"
	eventUserUpdateForTeam    string = "user_update_for_team"
	eventUserRemoveFromTeam   string = "user_remove_from_team"
	eventUserUpdateForGroup   string = "user_update_for_group"
	eventUserRemoveFromGroup  string = "user_remove_from_group"
)

// Options is a namespace variable for configuration options.
var Options = WebhookOptions{}

// Event defines a GitLab hook event type by the X-Gitlab-Event Header.
type Event string

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	secretHash []byte
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

	// шf a secret set is existing, it is necessary to check it in a constant time
	if len(hook.secretHash) > 0 {
		tokenHash := sha512.Sum512([]byte(r.Header.Get("X-Gitlab-Token")))
		if subtle.ConstantTimeCompare(tokenHash[:], hook.secretHash[:]) == 0 {
			return nil, errors.New("X-Gitlab-Token validation failed")
		}
	}

	event := r.Header.Get("X-Gitlab-Event")
	if len(event) == 0 {
		return nil, errors.New("missing X-Gitlab-Event Header")
	}

	gitLabEvent := Event(event)
	payload, err := io.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, errors.New("error parsing payload")
	}

	return eventParsing(gitLabEvent, events, payload)
}

func eventParsing(gitLabEvent Event, events []Event, payload []byte) (interface{}, error) {
	var found bool
	for _, evt := range events {
		if evt == gitLabEvent {
			found = true
			break
		}
	}

	// event not defined to be parsed
	if !found {
		return nil, errors.New("event not defined to be parsed")
	}

	switch gitLabEvent {
	case PushEvents:
		var pl PushEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case TagEvents:
		var pl TagEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case ConfidentialIssuesEvents:
		var pl ConfidentialIssueEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case IssuesEvents:
		var pl IssueEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case ConfidentialCommentEvents:
		var pl ConfidentialCommentEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case CommentEvents:
		var pl CommentEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case MergeRequestEvents:
		var pl MergeRequestEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case WikiPageEvents:
		var pl WikiPageEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case PipelineEvents:
		var pl PipelineEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case BuildEvents:
		var pl BuildEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case JobEvents:
		var pl JobEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		if err != nil {
			return nil, err
		}

		if pl.ObjectKind == objectBuild {
			return eventParsing(BuildEvents, events, payload)
		}

		return pl, nil
	case DeploymentEvents:
		var pl DeploymentEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		if err != nil {
			return nil, err
		}

		return pl, nil
	case SystemHookEvents:
		var pl SystemHookPayload
		err := json.Unmarshal([]byte(payload), &pl)
		if err != nil {
			return nil, err
		}

		switch pl.ObjectKind {
		case objectPush:
			return eventParsing(PushEvents, events, payload)
		case objectTag:
			return eventParsing(TagEvents, events, payload)
		case objectMergeRequest:
			return eventParsing(MergeRequestEvents, events, payload)
		default:
			switch pl.EventName {
			case objectPush:
				return eventParsing(PushEvents, events, payload)
			case objectTag:
				return eventParsing(TagEvents, events, payload)
			case objectMergeRequest:
				return eventParsing(MergeRequestEvents, events, payload)
			case eventProjectCreate:
				var pl ProjectCreatedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventProjectDestroy:
				var pl ProjectDestroyedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventProjectRename:
				var pl ProjectRenamedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventProjectTransfer:
				var pl ProjectTransferredEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventProjectUpdate:
				var pl ProjectUpdatedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventUserAddToTeam:
				var pl TeamMemberAddedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventUserRemoveFromTeam:
				var pl TeamMemberRemovedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventUserUpdateForTeam:
				var pl TeamMemberUpdatedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventUserCreate:
				var pl UserCreatedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventUserDestroy:
				var pl UserRemovedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventUserFailedLogin:
				var pl UserFailedLoginEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventUserRename:
				var pl UserRenamedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventKeyCreate:
				var pl KeyAddedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventKeyDestroy:
				var pl KeyRemovedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventGroupCreate:
				var pl GroupCreatedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventGroupDestroy:
				var pl GroupRemovedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventGroupRename:
				var pl GroupRenamedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventUserAddToGroup:
				var pl GroupMemberAddedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventUserRemoveFromGroup:
				var pl GroupMemberRemovedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			case eventUserUpdateForGroup:
				var pl GroupMemberUpdatedEventPayload
				err := json.Unmarshal([]byte(payload), &pl)
				return pl, err
			default:
				return nil, fmt.Errorf("unknown system hook event %s", gitLabEvent)
			}
		}
	case ReleaseEvents:
		var pl ReleaseEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	default:
		return nil, fmt.Errorf("unknown event %s", gitLabEvent)
	}
}

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

// WebhookOptions is a namespace for configuration option methods.
type WebhookOptions struct{}
