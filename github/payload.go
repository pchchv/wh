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
