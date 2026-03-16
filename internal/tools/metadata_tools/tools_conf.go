package tools

import (
	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var METADATA_TOOLS = []models.MCPTool{
	{
		Tool: &mcp.Tool{
			Name:        "describe_sobject",
			Description: "This tool is used to describe the sobject data",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"object": map[string]any{
						"type":        "string",
						"description": "The SObject name to describe",
					},
				},
			},
		},
		Handler: models.WrapHandler(DescribeSObject),
	},
	{
		Tool: &mcp.Tool{
			Name:        "list_sobjects",
			Description: "This tool is used to list all sobjects",
		},
		Handler: models.WrapHandler(ListSObjects),
	},
}
