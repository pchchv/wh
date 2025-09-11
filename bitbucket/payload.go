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
