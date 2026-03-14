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
		Handler: models.WrapHandler(QuerySObject),
	},
	{
		Tool: &mcp.Tool{
			Name:        "create_record",
			Description: "This tool is used to create any sobject record.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"object": map[string]any{
						"type":        "string",
						"description": "the specified sobject",
					},
					"fields": map[string]any{
						"type":        "object",
						"description": "the fields to create record",
					},
				},
			},
		},
		Handler: models.WrapHandler(CreateRecord),
	},
	{
		Tool: &mcp.Tool{
			Name:        "update_record",
			Description: "This tool is used to update any sobject record.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"id": map[string]any{
						"type":        "string",
						"description": "the record Id to update",
					},
					"object": map[string]any{
						"type":        "string",
						"description": "the specified sobject",
					},
					"fields": map[string]any{
						"type":        "object",
						"description": "the fields to update record",
					},
				},
			},
		},
		Handler: models.WrapHandler(UpdateRecord),
	},
}
