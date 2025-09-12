package bitbucket_server

import (
	"fmt"
	"strings"
	"time"
)

type DiagnosticsPingPayload struct{}

type User struct {
	ID           uint64                 `json:"id"`
	Active       bool                   `json:"active"`
	Name         string                 `json:"name"`
	Slug         string                 `json:"slug"`
	Type         string                 `json:"type"`
	DisplayName  string                 `json:"displayName"`
	EmailAddress string                 `json:"emailAddress"`
	Links        map[string]interface{} `json:"links"`
}

type Comment struct {
	ID                  uint64                   `json:"id"`
	Version             uint64                   `json:"version"`
	CreatedDate         uint64                   `json:"createdDate"`
	UpdatedDate         uint64                   `json:"updatedDate"`
	Author              User                     `json:"author"`
	Text                string                   `json:"text"`
	Tasks               []map[string]interface{} `json:"tasks"`
	Comments            []map[string]interface{} `json:"comments"`
	Properties          map[string]interface{}   `json:"properties,omitempty"`
	PermittedOperations map[string]interface{}   `json:"permittedOperations,omitempty"`
}

type Project struct {
	ID     uint64                 `json:"id"`
	Key    string                 `json:"key"`
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Public *bool                  `json:"public,omitempty"`
	Owner  User                   `json:"owner"`
	Links  map[string]interface{} `json:"links"`
}

type PullRequest struct {
	ID           uint64                   `json:"id"`
	Version      uint64                   `json:"version"`
	ClosedDate   uint64                   `json:"closedDate,omitempty"`
	CreatedDate  uint64                   `json:"createdDate"`
	UpdatedDate  uint64                   `json:"updatedDate,omitempty"`
	Title        string                   `json:"title"`
	State        string                   `json:"state"`
	Description  string                   `json:"description,omitempty"`
	Open         bool                     `json:"open"`
	Closed       bool                     `json:"closed"`
	Locked       bool                     `json:"locked"`
	ToRef        RepositoryReference      `json:"toRef"`
	FromRef      RepositoryReference      `json:"fromRef"`
	Author       PullRequestParticipant   `json:"author"`
	Reviewers    []PullRequestParticipant `json:"reviewers"`
	Participants []PullRequestParticipant `json:"participants"`
	Links        map[string]interface{}   `json:"links"`
	Properties   map[string]interface{}   `json:"properties,omitempty"`
}

type PullRequestParticipant struct {
	Role               string `json:"role"`
	Status             string `json:"status"`
	LastReviewedCommit string `json:"lastReviewedCommit,omitempty"`
	Approved           bool   `json:"approved"`
	User               User   `json:"user"`
}

type PullRequestReviewerUpdatedPayload struct {
	Date             Date        `json:"date"`
	Actor            User        `json:"actor"`
	EventKey         Event       `json:"eventKey"`
	PullRequest      PullRequest `json:"pullRequest"`
	AddedReviewers   []User      `json:"addedReviewers"`
	RemovedReviewers []User      `json:"removedReviewers"`
}

type PullRequestReviewerApprovedPayload struct {
	Date           Date                   `json:"date"`
	Actor          User                   `json:"actor"`
	EventKey       Event                  `json:"eventKey"`
	PullRequest    PullRequest            `json:"pullRequest"`
	Participant    PullRequestParticipant `json:"participant"`
	PreviousStatus string                 `json:"previousStatus"`
}

type PullRequestReviewerUnapprovedPayload struct {
	Date           Date                   `json:"date"`
	Actor          User                   `json:"actor"`
	EventKey       Event                  `json:"eventKey"`
	PullRequest    PullRequest            `json:"pullRequest"`
	Participant    PullRequestParticipant `json:"participant"`
	PreviousStatus string                 `json:"previousStatus"`
}

type PullRequestReviewerNeedsWorkPayload struct {
	Date           Date                   `json:"date"`
	Actor          User                   `json:"actor"`
	EventKey       Event                  `json:"eventKey"`
	PullRequest    PullRequest            `json:"pullRequest"`
	Participant    PullRequestParticipant `json:"participant"`
	PreviousStatus string                 `json:"previousStatus"`
}

type PullRequestCommentAddedPayload struct {
	Date            Date        `json:"date"`
	Actor           User        `json:"actor"`
	Comment         Comment     `json:"comment"`
	EventKey        Event       `json:"eventKey"`
	PullRequest     PullRequest `json:"pullRequest"`
	CommentParentID uint64      `json:"commentParentId,omitempty"`
}

type PullRequestCommentEditedPayload struct {
	Date            Date        `json:"date"`
	Actor           User        `json:"actor"`
	Comment         Comment     `json:"comment"`
	EventKey        Event       `json:"eventKey"`
	PullRequest     PullRequest `json:"pullRequest"`
	CommentParentID string      `json:"commentParentId,omitempty"`
	PreviousComment string      `json:"previousComment"`
}

type PullRequestCommentDeletedPayload struct {
	Date            Date        `json:"date"`
	Actor           User        `json:"actor"`
	Comment         Comment     `json:"comment"`
	EventKey        Event       `json:"eventKey"`
	PullRequest     PullRequest `json:"pullRequest"`
	CommentParentID uint64      `json:"commentParentId,omitempty"`
}

type PullRequestDeclinedPayload struct {
	Date        Date        `json:"date"`
	Actor       User        `json:"actor"`
	EventKey    Event       `json:"eventKey"`
	PullRequest PullRequest `json:"pullRequest"`
}

type PullRequestDeletedPayload struct {
	Date        Date        `json:"date"`
	Actor       User        `json:"actor"`
	EventKey    Event       `json:"eventKey"`
	PullRequest PullRequest `json:"pullRequest"`
}

type PullRequestMergedPayload struct {
	Date        Date        `json:"date"`
	Actor       User        `json:"actor"`
	EventKey    Event       `json:"eventKey"`
	PullRequest PullRequest `json:"pullRequest"`
}

type PullRequestFromReferenceUpdatedPayload struct {
	Date             Date        `json:"date"`
	Actor            User        `json:"actor"`
	EventKey         Event       `json:"eventKey"`
	PullRequest      PullRequest `json:"pullRequest"`
	PreviousFromHash string      `json:"previousFromHash"`
}

type Repository struct {
	ID            uint64                 `json:"id"`
	Public        bool                   `json:"public"`
	Forkable      bool                   `json:"forkable"`
	Slug          string                 `json:"slug"`
	Name          string                 `json:"name"`
	ScmID         string                 `json:"scmId"`
	State         string                 `json:"state"`
	StatusMessage string                 `json:"statusMessage"`
	Project       Project                `json:"project"`
	Origin        *Repository            `json:"origin,omitempty"`
	Links         map[string]interface{} `json:"links"`
}

type RepositoryChange struct {
	Type        string              `json:"type"`
	ToHash      string              `json:"toHash"`
	FromHash    string              `json:"fromHash"`
	ReferenceID string              `json:"refId"`
	Reference   RepositoryReference `json:"ref"`
}

type RepositoryReference struct {
	ID           string     `json:"id"`
	Type         string     `json:"type,omitempty"`
	DisplayID    string     `json:"displayId"`
	LatestCommit string     `json:"latestCommit,omitempty"`
	Repository   Repository `json:"repository,omitempty"`
}

type Date time.Time

func (b Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(b).Format("2006-01-02T15:04:05Z0700"))), nil
}

func (b *Date) UnmarshalJSON(p []byte) error {
	if t, err := time.Parse("2006-01-02T15:04:05Z0700", strings.Replace(string(p), "\"", "", -1)); err != nil {
		return err
	} else {
		*b = Date(t)
	}

	return nil
}
