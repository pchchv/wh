package azure

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	ImageURL    string `json:"imageUrl"`
	UniqueName  string `json:"uniqueName"`
	DisplayName string `json:"displayName"`
}

type Account struct {
	ID string `json:"id"`
}

type Commit struct {
	URL      string `json:"url"`
	CommitID string `json:"commitId"`
}

type Message struct {
	Text     string `json:"text"`
	HTML     string `json:"html"`
	Markdown string `json:"markdown"`
}

type Project struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type PushedBy struct {
	ID          string `json:"id"`
	UniqueName  string `json:"uniqueName"`
	DisplayName string `json:"displayName"`
}

type Repository struct {
	ID            string  `json:"id"`
	URL           string  `json:"url"`
	Name          string  `json:"name"`
	Project       Project `json:"project"`
	RemoteURL     string  `json:"remoteUrl"`
	DefaultBranch string  `json:"defaultBranch"`
}

type Reviewer struct {
	Vote        int    `json:"vote"`
	ID          string `json:"id"`
	URL         string `json:"url"`
	ImageURL    string `json:"imageUrl"`
	UniqueName  string `json:"uniqueName"`
	ReviewerURL string `json:"reviewerUrl"`
	DisplayName string `json:"displayName"`
	IsContainer bool   `json:"isContainer"`
}

type Request struct {
	ID           int    `json:"id"`
	URL          string `json:"url"`
	RequestedFor User   `json:"requestedFor"`
}

type RefUpdate struct {
	Name        string `json:"name"`
	NewObjectID string `json:"newObjectId"`
	OldObjectID string `json:"oldObjectId"`
}

type Resource struct {
	PushID     int         `json:"pushId"`
	URL        string      `json:"url"`
	Date       string      `json:"date"`
	Commits    []Commit    `json:"commits"`
	PushedBy   PushedBy    `json:"pushedBy"`
	Repository Repository  `json:"repository"`
	RefUpdates []RefUpdate `json:"refUpdates"`
}

type Log struct {
	URL         string `json:"url"`
	Type        string `json:"type"`
	DownloadURL string `json:"downloadUrl"`
}

type Drop struct {
	URL         string `json:"url"`
	Type        string `json:"type"`
	Location    string `json:"location"`
	DownloadURL string `json:"downloadUrl"`
}

type Queue struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	QueueType string `json:"queueType"`
}

type BuildDefinition struct {
	ID             int    `json:"id"`
	BatchSize      int    `json:"batchSize"`
	DefinitionType string `json:"definitionType"`
	TriggerType    string `json:"triggerType"`
	Name           string `json:"name"`
	URL            string `json:"url"`
}

type ResourceContainers struct {
	Account    Account `json:"account"`
	Project    Account `json:"project"`
	Collection Account `json:"collection"`
}

type Build struct {
	ID                 int             `json:"id"`
	HasDiagnostics     bool            `json:"hasDiagnostics"`
	RetainIndefinitely bool            `json:"retainIndefinitely"`
	URI                string          `json:"uri"`
	URL                string          `json:"url"`
	Reason             string          `json:"reason"`
	Status             string          `json:"status"`
	BuildNumber        string          `json:"buildNumber"`
	DropLocation       string          `json:"dropLocation"`
	SourceGetVersion   string          `json:"sourceGetVersion"`
	Log                Log             `json:"log"`
	Drop               Drop            `json:"drop"`
	Queue              Queue           `json:"queue"`
	Requests           []Request       `json:"requests"`
	StartTime          Date            `json:"startTime"`
	FinishTime         Date            `json:"finishTime"`
	Definition         BuildDefinition `json:"definition"`
	LastChangedBy      User            `json:"lastChangedBy"`
}

type PullRequest struct {
	PullRequestID         int        `json:"pullRequestId"`
	URL                   string     `json:"url"`
	Title                 string     `json:"title"`
	MergeID               string     `json:"mergeId"`
	Description           string     `json:"description"`
	MergeStatus           string     `json:"mergeStatus"`
	SourceRefName         string     `json:"sourceRefName"`
	TargetRefName         string     `json:"targetRefName"`
	Status                string     `json:"status"`
	Commits               []Commit   `json:"commits"`
	Reviewers             []Reviewer `json:"reviewers"`
	CreatedBy             User       `json:"createdBy"`
	ClosedDate            Date       `json:"closedDate"`
	Repository            Repository `json:"repository"`
	CreationDate          Date       `json:"creationDate"`
	LastMergeCommit       Commit     `json:"lastMergeCommit"`
	LastMergeSourceCommit Commit     `json:"lastMergeSourceCommit"`
	LastMergeTargetCommit Commit     `json:"lastMergeTargetCommit"`
}

// Azure DevOps does not send an event header, this BasicEvent is provided to get the EventType.
type BasicEvent struct {
	ID          string `json:"id"`
	Scope       string `json:"scope"`
	PublisherID string `json:"publisherId"`
	CreatedDate Date   `json:"createdDate"`
	EventType   Event  `json:"eventType"`
}

// git.push
type GitPushEvent struct {
	ID                 string             `json:"id"`
	Scope              string             `json:"scope"`
	EventType          string             `json:"eventType"`
	CreatedDate        string             `json:"createdDate"`
	PublisherID        string             `json:"publisherId"`
	ResourceVersion    string             `json:"resourceVersion"`
	Message            Message            `json:"message"`
	DetailedMessage    Message            `json:"detailedMessage"`
	ResourceContainers ResourceContainers `json:"resourceContainers"`
	Resource           Resource           `json:"resource"`
}

// git.pullrequest.*
// git.pullrequest.merged
// git.pullrequest.created
// git.pullrequest.updated
type GitPullRequestEvent struct {
	ID                 string      `json:"id"`
	Scope              string      `json:"scope"`
	PublisherID        string      `json:"publisherId"`
	ResourceVersion    string      `json:"resourceVersion"`
	CreatedDate        Date        `json:"createdDate"`
	EventType          Event       `json:"eventType"`
	Message            Message     `json:"message"`
	Resource           PullRequest `json:"resource"`
	DetailedMessage    Message     `json:"detailedMessage"`
	ResourceContainers interface{} `json:"resourceContainers"`
}

// build.complete
type BuildCompleteEvent struct {
	ID                 string      `json:"id"`
	Scope              string      `json:"scope"`
	PublisherID        string      `json:"publisherId"`
	ResourceVersion    string      `json:"resourceVersion"`
	CreatedDate        Date        `json:"createdDate"`
	EventType          Event       `json:"eventType"`
	Message            Message     `json:"message"`
	Resource           Build       `json:"resource"`
	DetailedMessage    Message     `json:"detailedMessage"`
	ResourceContainers interface{} `json:"resourceContainers"`
}

type Date time.Time

func (b Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(b).Format(time.RFC3339Nano))), nil
}

func (b *Date) UnmarshalJSON(p []byte) error {
	if t, err := time.Parse(time.RFC3339Nano, strings.Replace(string(p), "\"", "", -1)); err != nil {
		return err
	} else {
		*b = Date(t)
	}

	return nil
}
