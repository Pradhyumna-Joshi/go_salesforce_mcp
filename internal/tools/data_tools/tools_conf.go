package tools

import (
	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var DATA_TOOLS = []models.MCPTools{
	{
		Tool: &mcp.Tool{
			Name:        "Query_Sobject",
			Description: "This tool is used to query any Sobject from salesforce org. Send a full SOQL Query",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"soql": map[string]any{
						"type":        "string",
						"description": "The SOQL query to execute",
					},
				},
			},
		},
		Handler: models.WrapHandler(QuerySObjectHandler),
	},
}
