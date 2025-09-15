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
