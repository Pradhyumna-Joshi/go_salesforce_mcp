package tools

type ConnectToOrgInput struct {
	OrgType string `json:"orgType" jsonschema:"description=orgType: production or sandbox"`
	Alias   string `json:"alias" jsonschema:",description=alias for this org. Lowercase, no spaces. Example: prod-org, my-sandbox"`
}

type OpenOrgInput struct {
	Alias string `json:"alias" jsonschema:",description=alias of the org to open. Use list_orgs tool if alias is unknown."`
}

type DisconnectOrgInput struct {
	Alias string `json:"alias" jsonschema:",description=alias of the org to disconnect. Use list_orgs tool if alias is unknown."`
}

type OrgResult struct {
	Result struct {
		AccessToken string `json:"accessToken"`
		InstanceURL string `json:"instanceUrl"`
		Username    string `json:"username"`
	} `json:"result"`
}
