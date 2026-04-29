package tools

import (
	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
)

var APEX_TOOLS = []models.MCPTool{
	models.NewTool(
		"apex_rest_get",
		"Call custom GET Apex REST services. Use this when the org has custom integration endpoints defined with @HttpGet.",
		ApexRestGet,
	),
	models.NewTool(
		"apex_rest_post",
		"Call custom POST Apex REST services. Best for complex operations or record creation handled by custom @HttpPost methods.",
		ApexRestPost,
	),
	models.NewTool(
		"apex_rest_patch",
		"Call custom PATCH Apex REST services. Use this for partial updates to custom integration endpoints defined with @HttpPatch.",
		ApexRestPatch,
	),

	models.NewTool(
		"apex_class_get",
		"Retrieve Apex class source code and metadata. You can provide either the 'id' OR the 'name'; the tool will automatically resolve the ID if only the Name is provided.",
		GetApexClass,
	),
	models.NewTool(
		"apex_class_create",
		"Deploy a new Apex class to the org. Requires valid class body and name.",
		CreateApexClass,
	),
	models.NewTool(
		"apex_class_update",
		"Update source code for an existing Apex class. You can provide 'id' or 'name'. This replaces the entire class body.",
		UpdateApexClass,
	),
	models.NewTool(
		"apex_trigger_get",
		"Retrieve Apex trigger source code and object bindings. Supports resolution by 'id' or 'name'.",
		GetApexTrigger,
	),

	models.NewTool(
		"apex_execute_anonymous",
		"Execute a block of Apex code anonymously. Use this to test logic snippets, run one-off scripts, or verify behavior without creating a class.",
		ApexExecuteAnonymous,
	),
	models.NewTool(
		"apex_tooling_query",
		"Perform SOQL queries specifically against Tooling API objects (e.g., ApexClass, ApexLog, TraceFlag). Use this for developer-meta-queries.",
		ApexToolingQuery,
	),
	models.NewTool(
		"apex_logs_list",
		"Fetch the most recent debug logs from the org. Returns IDs, status, and duration. Use this to find a LogId after an error occurs.",
		ListApexLogs,
	),
	models.NewTool(
		"apex_log_get_body",
		"Retrieve the full raw text of a specific debug log by its ID. Essential for reading System.debug output.",
		GetApexLogBody,
	),

	models.NewTool(
		"apex_tests_run_async",
		"Enqueue an asynchronous test run for classes or suites. Returns a job ID; you must check results later.",
		RunApexTestsAsync,
	),
	models.NewTool(
		"apex_code_coverage_get",
		"Get code coverage percentages for a specific class or trigger. Supports lookup by 'class_or_trigger_id', 'class_name', or 'trigger_name'.",
		GetApexCodeCoverage,
	),
}
