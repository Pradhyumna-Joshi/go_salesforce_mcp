package tools

import (
	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
)

var METADATA_TOOLS = []models.MCPTool{
	models.NewTool("describe_sobject", "Describe a Salesforce SObject schema including fields, types, and relationships. Use this before create or update to know available fields. Example: Account Contact Opportunity", DescribeSObject),
	models.NewTool("list_sobjects", "List all available Salesforce SObject API names in the org. Use this when you don't know the SObject name to query or modify.", ListSObjects),
}
