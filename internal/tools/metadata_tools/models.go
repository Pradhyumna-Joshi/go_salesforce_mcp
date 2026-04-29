package tools

type EmptyRequest struct{}

type SObjectRequest struct {
	Object string `json:"object" jsonschema:"description=API name of the SObject (e.g., Account, Case)"`
}

type FieldInfoRequest struct {
	Object string `json:"object" jsonschema:"description=API name of the SObject"`
	Field  string `json:"field" jsonschema:"description=API name of the specific field"`
}

type SearchRequest struct {
	Query string `json:"query" jsonschema:"description=The search term or SOSL query string"`
}
