package bitbucket_server

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

type PullRequestParticipant struct {
	Role               string `json:"role"`
	Status             string `json:"status"`
	LastReviewedCommit string `json:"lastReviewedCommit,omitempty"`
	Approved           bool   `json:"approved"`
	User               User   `json:"user"`
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
