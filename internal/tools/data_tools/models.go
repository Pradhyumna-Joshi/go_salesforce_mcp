package tools

type QueryRequest struct {
	Soql string `json:"soql"`
}

type QueryResponse struct {
	Response string
}
