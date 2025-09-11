package bitbucket

// Owner is the common Bitbucket Owner Sub Entity.
type Owner struct {
	UUID        string `json:"uuid"`
	Type        string `json:"type"`
	NickName    string `json:"nickname"`
	AccountID   string `json:"account_id"`
	DisplayName string `json:"display_name"`
	Links       struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
	} `json:"links"`
}

// Project is the common Bitbucket Project Sub Entity.
type Project struct {
	Type    string `json:"type"`
	UUID    string `json:"uuid"`
	Project string `json:"project"`
	Links   struct {
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
	} `json:"links"`
	Key string `json:"key"`
}

// Repository is the common Bitbucket Repository Sub Entity.
type Repository struct {
	Type      string `json:"type"`
	Scm       string `json:"scm"`
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Website   string `json:"website"`
	FullName  string `json:"full_name"`
	IsPrivate bool   `json:"is_private"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
	} `json:"links"`
	Project Project `json:"project"`
	Owner   Owner   `json:"owner"`
}

// RepoForkPayload is the Bitbucket repo:fork payload.
type RepoForkPayload struct {
	Actor      Owner      `json:"actor"`
	Fork       Repository `json:"fork"`
	Repository Repository `json:"repository"`
}

// RepoUpdatedPayload is the Bitbucket repo:updated payload.
type RepoUpdatedPayload struct {
	Actor      Owner      `json:"actor"`
	Repository Repository `json:"repository"`
	Changes    struct {
		Name struct {
			New string `json:"new"`
			Old string `json:"old"`
		} `json:"name"`
		Website struct {
			New string `json:"new"`
			Old string `json:"old"`
		} `json:"website"`
		Language struct {
			New string `json:"new"`
			Old string `json:"old"`
		} `json:"language"`
		Links struct {
			New struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"new"`
			Old struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"old"`
		} `json:"links"`
		Description struct {
			New string `json:"new"`
			Old string `json:"old"`
		} `json:"description"`
		FullName struct {
			New string `json:"new"`
			Old string `json:"old"`
		} `json:"full_name"`
	} `json:"changes"`
}
