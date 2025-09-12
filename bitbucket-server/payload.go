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
