package tools

type QueryPayload struct {
	Soql string `json:"soql"`
}

type QueryResponse struct {
	Response string
}

type CreateRecordPayload struct {
	Object string         `json:"object"`
	Fields map[string]any `json:"fields"`
}

type UpdateRecordPayload struct {
	Id     string         `json:"id"`
	Object string         `json:"object"`
	Fields map[string]any `json:"fields"`
}
