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
