package github

// Team contains GitHub's Team information.
type Team struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	URL             string `json:"url"`
	Slug            string `json:"slug"`
	NodeID          string `json:"node_id"`
	Permission      string `json:"permission"`
	MembersURL      string `json:"members_url"`
	RepositoriesURL string `json:"repositories_url"`
	Parent          *Team  `json:"parent,omitempty"`
}

// Label contains Issue's Label information.
type Label struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Color   string `json:"color"`
	NodeID  string `json:"node_id"`
	Default bool   `json:"default"`
}

// Parent contains GitHub's parent information.
type Parent struct {
	URL string `json:"url"`
	Sha string `json:"sha"`
}
