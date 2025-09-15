// azure devops does not send an event header, this BasicEvent is provided to get the EventType
package azure

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
