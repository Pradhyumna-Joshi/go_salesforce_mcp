package tools

type Empty struct{}

type OrgResult struct {
	Result struct {
		AccessToken string `json:"accessToken"`
		InstanceURL string `json:"instanceUrl"`
		Username    string `json:"username"`
	} `json:"result"`
}
