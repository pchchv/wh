package docker

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Docker hook types (only one for now).
const BuildEvent Event = "build"

// parse error
var ErrParsingPayload = errors.New("error parsing payload")

// Event defines a Docker hook event type.
type Event string

// BuildPayload a docker hub build notice.
//
// https://docs.docker.com/docker-hub/webhooks/
type BuildPayload struct {
	CallbackURL string `json:"callback_url"`
	PushData    struct {
		Tag      string   `json:"tag"`
		Pusher   string   `json:"pusher"`
		Images   []string `json:"images"`
		PushedAt float32  `json:"pushed_at"`
	} `json:"push_data"`
	Repository struct {
		CommentCount    int     `json:"comment_count"`
		StarCount       int     `json:"star_count"`
		DateCreated     float32 `json:"date_created"`
		FullDescription string  `json:"full_description"`
		Description     string  `json:"description"`
		Dockerfile      string  `json:"dockerfile"`
		Namespace       string  `json:"namespace"`
		RepoName        string  `json:"repo_name"`
		RepoURL         string  `json:"repo_url"`
		Status          string  `json:"status"`
		Owner           string  `json:"owner"`
		Name            string  `json:"name"`
		IsOfficial      bool    `json:"is_official"`
		IsPrivate       bool    `json:"is_private"`
		IsTrusted       bool    `json:"is_trusted"`
	} `json:"repository"`
}

// Webhook instance contains all methods needed to process events.
type Webhook struct{}

// New creates and returns a WebHook instance.
func New() (*Webhook, error) {
	return new(Webhook), nil
}

// Parse verifies and parses the events specified and returns the payload object or an error.
func (hook Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
	defer func() {
		_, _ = io.Copy(io.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if r.Method != http.MethodPost {
		return nil, errors.New("invalid HTTP Method")
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	var pl BuildPayload
	if err = json.Unmarshal([]byte(payload), &pl); err != nil {
		return nil, ErrParsingPayload
	}

	return pl, nil
}
