package tools

import (
	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
)

var APEX_TOOLS = []models.MCPTool{
	// Apex REST
	models.NewTool(
		"apex_rest_get",
		"Execute a GET request on a custom Apex REST endpoint to retrieve data",
		ApexRestGet,
	),
	models.NewTool(
		"apex_rest_post",
		"Execute a POST request on a custom Apex REST endpoint to create records",
		ApexRestPost,
	),
	models.NewTool(
		"apex_rest_patch",
		"Execute a PATCH request on a custom Apex REST endpoint to update records",
		ApexRestPatch,
	),

	// Apex Execution & Tooling Query
	models.NewTool(
		"apex_execute_anonymous",
		"Execute anonymous Apex code using the Tooling API and return execution results",
		ApexExecuteAnonymous,
	),
	models.NewTool(
		"apex_tooling_query",
		"Execute a SOQL query against Tooling API objects like ApexClass, ApexTrigger, ApexLog, and TraceFlag",
		ApexToolingQuery,
	),

	// Apex Classes
	models.NewTool(
		"apex_class_get",
		"Retrieve an Apex class by Id or Name using the Tooling API",
		GetApexClass,
	),
	models.NewTool(
		"apex_class_create",
		"Create a new Apex class with the provided name and body",
		CreateApexClass,
	),
	models.NewTool(
		"apex_class_update",
		"Update an existing Apex class by Id or Name with new source code",
		UpdateApexClass,
	),

	// Apex Triggers
	models.NewTool(
		"apex_trigger_get",
		"Retrieve an Apex trigger by Id or Name using the Tooling API",
		GetApexTrigger,
	),
	models.NewTool(
		"apex_trigger_create",
		"Create a new Apex trigger for a specified object",
		CreateApexTrigger,
	),
	models.NewTool(
		"apex_trigger_update",
		"Update an existing Apex trigger by Id or Name with new source code",
		UpdateApexTrigger,
	),

	// Logs
	models.NewTool(
		"apex_logs_list",
		"Retrieve recent Apex debug logs ordered by most recent first",
		ListApexLogs,
	),
	models.NewTool(
		"apex_log_get_body",
		"Retrieve the raw body of a specific Apex debug log by Id",
		GetApexLogBody,
	),

	// Trace Flags
	models.NewTool(
		"trace_flag_get",
		"Retrieve a TraceFlag configuration by Id",
		GetTraceFlag,
	),
	models.NewTool(
		"trace_flag_create",
		"Create a TraceFlag to enable debug logging for a user or Apex entity (Id or Name)",
		CreateTraceFlag,
	),
	models.NewTool(
		"trace_flag_update",
		"Update an existing TraceFlag, such as extending expiration or changing debug level",
		UpdateTraceFlag,
	),

	// Apex Tests
	models.NewTool(
		"apex_tests_run_async",
		"Run Apex test classes or suites asynchronously and return a job Id",
		RunApexTestsAsync,
	),
	models.NewTool(
		"apex_tests_get_result",
		"Retrieve results of an asynchronous Apex test run using AsyncApexJob Id",
		GetApexTestResult,
	),
	models.NewTool(
		"apex_tests_run_sync",
		"Run Apex test classes synchronously and return immediate results",
		RunApexTestsSync,
	),

	// Code Coverage
	models.NewTool(
		"apex_code_coverage_get",
		"Retrieve line-level code coverage for an Apex class or trigger using Id, class name, or trigger name",
		GetApexCodeCoverage,
	),
	models.NewTool(
		"apex_code_coverage_org",
		"Retrieve overall Apex code coverage percentage for the Salesforce org",
		GetOrgWideCoverage,
	),
}
