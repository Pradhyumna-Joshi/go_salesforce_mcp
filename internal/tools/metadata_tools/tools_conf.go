package tools

import (
	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
)

var METADATA_TOOLS = []models.MCPTool{
	models.NewTool("get_sobject_fields",
		"CRITICAL: Call this before ANY SOQL query to verify custom field API names. Use this to prevent 'INVALID_FIELD' errors.",
		GetSObjectFields),

	models.NewTool("get_field_details",
		"Returns full metadata for a specific field (picklist values, lengths). Use this before updating a record.",
		GetFieldDetails),

	models.NewTool("get_child_relationships",
		"Identifies related child objects. Use this when you need to query related records (e.g., Contacts for an Account).",
		GetChildRelationships),

	models.NewTool("get_record_types",
		"Lists available Record Types for an object. Mandatory for creating new records in objects that use Record Types.",
		GetRecordTypes),

	models.NewTool("global_search",
		"Searches for a keyword across the entire Salesforce org. Use this when you have an ID but don't know the object type.",
		GlobalSearch),

	models.NewTool("get_system_context",
		"Returns current API limits and organization info. Use this to monitor usage and environment status.",
		GetSystemContext),
}
