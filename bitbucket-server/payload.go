package bitbucket_server

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
