package tools

import (
	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
)

var DATA_TOOLS = []models.MCPTool{
	models.NewTool("query_sobject", "Query Salesforce records using SOQL", QuerySObject),
	models.NewTool("create_record", "Create a new Salesforce record for a given SObject", CreateRecord),
	models.NewTool("update_record", "Update an existing Salesforce record by Id", UpdateRecord),
}
