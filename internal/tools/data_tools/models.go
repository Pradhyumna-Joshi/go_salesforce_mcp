package tools

type QueryPayload struct {
	Soql string `json:"soql" jsonschema:"description=The SOQL query to execute. Example: SELECT Id Name FROM Account LIMIT 10"`
}

type QueryResponse struct {
	Response string
}

type CreateRecordPayload struct {
	Object string         `json:"object" jsonschema:"description=Salesforce SObject API name. Example: Account Contact Lead"`
	Fields map[string]any `json:"fields" jsonschema:"description=Map of field API names to values. Example: {Name:Acme Phone:123-456-7890}"`
}

type UpdateRecordPayload struct {
	Id     string         `json:"id" jsonschema:"description=Salesforce record Id to update. Example: 001XX000003GYkZYAW"`
	Object string         `json:"object" jsonschema:"description=Salesforce SObject API name. Example: Account Contact Lead"`
	Fields map[string]any `json:"fields" jsonschema:"description=Only the fields to update with new values. Example: {Phone:999-999-9999}"`
}
