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
