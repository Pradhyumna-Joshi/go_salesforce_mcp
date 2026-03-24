package tools

type ApexRestGetPayload struct {
	Path        string            `json:"path" jsonschema:"description=Apex REST resource path. Example: /MyApexClass/someEndpoint"`
	QueryParams map[string]string `json:"query_params,omitempty" jsonschema:"description=Optional URL query parameters. Example: {status:active version:2}"`
}

type ApexRestPostPayload struct {
	Path string         `json:"path" jsonschema:"description=Apex REST resource path. Example: /MyApexClass/someEndpoint"`
	Body map[string]any `json:"body" jsonschema:"description=JSON body to send to the Apex REST endpoint. Example: {accountName:Acme amount:5000}"`
}

type ApexRestPatchPayload struct {
	Path string         `json:"path" jsonschema:"description=Apex REST resource path including record identifier. Example: /MyApexClass/001XX000003GYkZYAW"`
	Body map[string]any `json:"body" jsonschema:"description=Fields to update via the Apex REST endpoint. Example: {status:Closed amount:9999}"`
}

type ApexExecuteAnonymousPayload struct {
	Body string `json:"body" jsonschema:"description=Apex code to execute anonymously via Tooling API. Example: System.debug('Hello World');"`
}

type ApexToolingQueryPayload struct {
	Soql string `json:"soql" jsonschema:"description=SOQL query against Tooling API objects. Example: SELECT Id Name Body FROM ApexClass WHERE NamespacePrefix = null LIMIT 10"`
}

// Apex Class

type GetApexClassPayload struct {
	Id   string `json:"id,omitempty" jsonschema:"description=Salesforce ApexClass Id. Example: 01p2800000A86adAAB"`
	Name string `json:"name,omitempty" jsonschema:"description=Apex class name (alternative to Id). Example: MyUtilityClass"`
}

type CreateApexClassPayload struct {
	Name string `json:"name" jsonschema:"description=API name for the new Apex class. Example: MyUtilityClass"`
	Body string `json:"body" jsonschema:"description=Full Apex source code for the class. Example: public class MyUtilityClass { public static void run() {} }"`
}

type UpdateApexClassPayload struct {
	Id   string `json:"id,omitempty" jsonschema:"description=Salesforce ApexClass Id to update"`
	Name string `json:"name,omitempty" jsonschema:"description=Apex class name (alternative to Id)"`
	Body string `json:"body" jsonschema:"description=Updated full Apex source code for the class"`
}

// Apex Trigger

type GetApexTriggerPayload struct {
	Id   string `json:"id,omitempty" jsonschema:"description=Salesforce ApexTrigger Id"`
	Name string `json:"name,omitempty" jsonschema:"description=Apex trigger name (alternative to Id). Example: AccountTrigger"`
}

type CreateApexTriggerPayload struct {
	Name          string `json:"name" jsonschema:"description=API name for the new Apex trigger. Example: AccountTrigger"`
	TableEnumOrId string `json:"table_enum_or_id" jsonschema:"description=SObject API name the trigger fires on. Example: Account Contact Opportunity"`
	Body          string `json:"body" jsonschema:"description=Full Apex source code for the trigger. Example: trigger AccountTrigger on Account (before insert) { }"`
}

type UpdateApexTriggerPayload struct {
	Id   string `json:"id,omitempty" jsonschema:"description=Salesforce ApexTrigger Id to update"`
	Name string `json:"name,omitempty" jsonschema:"description=Apex trigger name (alternative to Id)"`
	Body string `json:"body" jsonschema:"description=Updated full Apex source code for the trigger"`
}

// Logs

type ListApexLogsPayload struct {
	Limit int `json:"limit,omitempty" jsonschema:"description=Max number of log records to return. Defaults to 25 if omitted. Example: 10"`
}

type GetApexLogBodyPayload struct {
	Id string `json:"id" jsonschema:"description=ApexLog record Id whose raw body to retrieve. Example: 07L2800000BnZzKEAV"`
}

// Trace Flags

type GetTraceFlagPayload struct {
	Id string `json:"id" jsonschema:"description=TraceFlag record Id. Example: 7tf2800000BpQrtAAF"`
}

type CreateTraceFlagPayload struct {
	TracedEntityId   string `json:"traced_entity_id,omitempty" jsonschema:"description=Id of the user or Apex class to trace"`
	TracedEntityName string `json:"traced_entity_name,omitempty" jsonschema:"description=Name of user or Apex class (alternative to Id)"`
	DebugLevelId     string `json:"debug_level_id" jsonschema:"description=Id of the DebugLevel to apply"`
	LogType          string `json:"log_type" jsonschema:"description=Type of log to generate. Example: USER_DEBUG APEX_PROFILING CALLOUT"`
	ExpirationDate   string `json:"expiration_date" jsonschema:"description=ISO-8601 datetime when the flag expires. Example: 2025-12-31T23:59:59.000Z"`
}

type UpdateTraceFlagPayload struct {
	Id             string `json:"id" jsonschema:"description=TraceFlag record Id to update"`
	ExpirationDate string `json:"expiration_date" jsonschema:"description=New ISO-8601 expiry datetime"`
	DebugLevelId   string `json:"debug_level_id,omitempty" jsonschema:"description=Optional new DebugLevel Id"`
}

// Tests

type RunApexTestsAsyncPayload struct {
	ClassNames     []string `json:"class_names,omitempty" jsonschema:"description=List of Apex test class names to enqueue. Example: [MyClassTest AnotherClassTest]"`
	SuiteNames     []string `json:"suite_names,omitempty" jsonschema:"description=List of Apex test suite names to enqueue. Example: [MyTestSuite]"`
	MaxFailedTests int      `json:"max_failed_tests,omitempty" jsonschema:"description=Abort run after this many failures. Use -1 to run all. Example: 5"`
}

type GetApexTestResultPayload struct {
	AsyncApexJobId string `json:"async_apex_job_id" jsonschema:"description=Async job Id returned by runTestsAsynchronous. Example: 707XX0000002bCoQAI"`
}

type RunApexTestsSyncPayload struct {
	ClassNames  []string `json:"class_names" jsonschema:"description=List of Apex test class names to run synchronously. Example: [MyClassTest AnotherClassTest]"`
	TestMethods []string `json:"test_methods,omitempty" jsonschema:"description=Optional specific test method names. Example: [testInsert testUpdate]"`
}

// Code Coverage

type GetApexCodeCoveragePayload struct {
	ClassOrTriggerId string `json:"class_or_trigger_id,omitempty" jsonschema:"description=ApexClass or ApexTrigger Id. Example: 01p2800000A86adAAB"`
	ClassName        string `json:"class_name,omitempty" jsonschema:"description=Apex class name (alternative to Id). Example: MyClass"`
	TriggerName      string `json:"trigger_name,omitempty" jsonschema:"description=Apex trigger name (alternative to Id). Example: AccountTrigger"`
}

type GetOrgWideCoveragePayload struct{}

const apiVersion = "v60.0"
