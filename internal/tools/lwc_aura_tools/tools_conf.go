package tools

import "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"

var UI_COMPONENTS_TOOLS = []models.MCPTool{
	models.NewTool(
		"aura_definition_get",
		"Retrieve Aura component source code. Requires bundle_name and aura_type (COMPONENT, CONTROLLER, STYLE).",
		GetAuraDefinition,
	),

	models.NewTool(
		"lwc_bundle_list_files",
		"Lists all files (HTML, JS, CSS) within an LWC bundle. Use this to find the resource_id needed for updates.",
		ListLWCWithResources,
	),
	models.NewTool(
		"lwc_resource_get",
		"Retrieve the source code of a specific LWC file by its bundle name and file path.",
		GetLWCResource,
	),
	models.NewTool(
		"lwc_resource_update",
		"Update the source code of a specific LWC file. Requires the resource_id (found via lwc_bundle_list_files).",
		UpdateLWCResource,
	),
}
