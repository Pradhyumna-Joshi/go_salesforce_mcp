package tools

type SObjectRequest struct {
	Object string `json:"object" jsonschema:"description=Salesforce SObject API name. Example: Account Contact Opportunity Lead Case"`
}
