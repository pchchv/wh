package bitbucket

import "time"

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

// Comment is the common Bitbucket Comment Sub Entity.
type Comment struct {
	ID     int64 `json:"id"`
	Parent struct {
		ID int64 `json:"id"`
	} `json:"parent"`
	Content struct {
		Raw    string `json:"raw"`
		HTML   string `json:"html"`
		Markup string `json:"markup"`
	} `json:"content"`
	Inline struct {
		Path string `json:"path"`
		From *int64 `json:"from"`
		To   int64  `json:"to"`
	} `json:"inline"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	Links     struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
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

// RepoPushPayload is the Bitbucket repo:push payload.
type RepoPushPayload struct {
	Actor      Owner      `json:"actor"`
	Repository Repository `json:"repository"`
	Push       struct {
		Changes []struct {
			New struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Target struct {
					Type    string    `json:"type"`
					Hash    string    `json:"hash"`
					Author  Owner     `json:"author"`
					Message string    `json:"message"`
					Date    time.Time `json:"date"`
					Parents []struct {
						Type  string `json:"type"`
						Hash  string `json:"hash"`
						Links struct {
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
							HTML struct {
								Href string `json:"href"`
							} `json:"html"`
						} `json:"links"`
					} `json:"parents"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						HTML struct {
							Href string `json:"href"`
						} `json:"html"`
					} `json:"links"`
				} `json:"target"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Commits struct {
						Href string `json:"href"`
					} `json:"commits"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"new"`
			Old struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Target struct {
					Type    string    `json:"type"`
					Hash    string    `json:"hash"`
					Author  Owner     `json:"author"`
					Message string    `json:"message"`
					Date    time.Time `json:"date"`
					Parents []struct {
						Type  string `json:"type"`
						Hash  string `json:"hash"`
						Links struct {
							Self struct {
								Href string `json:"href"`
							} `json:"self"`
							HTML struct {
								Href string `json:"href"`
							} `json:"html"`
						} `json:"links"`
					} `json:"parents"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						HTML struct {
							Href string `json:"href"`
						} `json:"html"`
					} `json:"links"`
				} `json:"target"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Commits struct {
						Href string `json:"href"`
					} `json:"commits"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"old"`
			Links struct {
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Diff struct {
					Href string `json:"href"`
				} `json:"diff"`
				Commits struct {
					Href string `json:"href"`
				} `json:"commits"`
			} `json:"links"`
			Created bool `json:"created"`
			Forced  bool `json:"forced"`
			Closed  bool `json:"closed"`
			Commits []struct {
				Hash    string `json:"hash"`
				Type    string `json:"type"`
				Message string `json:"message"`
				Author  Owner  `json:"author"`
				Links   struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
			} `json:"commits"`
			Truncated bool `json:"truncated"`
		} `json:"changes"`
	} `json:"push"`
}

// RepoCommitStatusCreatedPayload is the Bitbucket repo:commit_status_created payload.
type RepoCommitStatusCreatedPayload struct {
	Actor        Owner      `json:"actor"`
	Repository   Repository `json:"repository"`
	CommitStatus struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		State       string    `json:"state"`
		Key         string    `json:"key"`
		URL         string    `json:"url"`
		Type        string    `json:"type"`
		CreatedOn   time.Time `json:"created_on"`
		UpdatedOn   time.Time `json:"updated_on"`
		Links       struct {
			Commit struct {
				Href string `json:"href"`
			} `json:"commit"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"commit_status"`
}

// RepoCommitStatusUpdatedPayload is the Bitbucket repo:commit_status_updated payload.
type RepoCommitStatusUpdatedPayload struct {
	Actor        Owner      `json:"actor"`
	Repository   Repository `json:"repository"`
	CommitStatus struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		State       string    `json:"state"`
		Key         string    `json:"key"`
		URL         string    `json:"url"`
		Type        string    `json:"type"`
		CreatedOn   time.Time `json:"created_on"`
		UpdatedOn   time.Time `json:"updated_on"`
		Links       struct {
			Commit struct {
				Href string `json:"href"`
			} `json:"commit"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"commit_status"`
}
