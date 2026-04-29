package tools

const ApiVersion = "v60.0"

type EmptyRequest struct{}

type ApexRestGetPayload struct {
	Path        string            `json:"path" jsonschema:"description=Apex REST resource path. Example: /MyApexClass/someEndpoint"`
	QueryParams map[string]string `json:"query_params,omitempty" jsonschema:"description=Optional URL query parameters. Example: {status:active version:2}"`
}

type ApexRestPostPayload struct {
	Path string `json:"path" jsonschema:"description=Apex REST resource path. Example: /MyApexClass/someEndpoint"`
	Body any    `json:"body" jsonschema:"description=JSON body to send. Can be a map {key:value} or an array [{},{}]."`
}

type ApexRestPatchPayload struct {
	Path string `json:"path" jsonschema:"description=Apex REST resource path including record identifier. Example: /MyApexClass/001XX000003GYkZYAW"`
	Body any    `json:"body" jsonschema:"description=Fields to update via the Apex REST endpoint. Can be a map or an array."`
}

type ApexExecuteAnonymousPayload struct {
	Body string `json:"body" jsonschema:"description=Apex code to execute anonymously. Example: System.debug('Hello World');"`
}

type ApexToolingQueryPayload struct {
	Soql string `json:"soql" jsonschema:"description=SOQL query against Tooling API objects. Example: SELECT Id, Name, Body FROM ApexClass WHERE NamespacePrefix = null LIMIT 10"`
}

type GetApexClassPayload struct {
	Id   string `json:"id,omitempty" jsonschema:"description=Salesforce ApexClass Id. Example: 01p2800000A86adAAB"`
	Name string `json:"name,omitempty" jsonschema:"description=Apex class name (alternative to Id). Example: MyUtilityClass"`
}

type CreateApexClassPayload struct {
	Name string `json:"name" jsonschema:"description=API name for the new Apex class. Example: MyUtilityClass"`
	Body string `json:"body" jsonschema:"description=Full Apex source code for the class. Must include class declaration."`
}

type UpdateApexClassPayload struct {
	Id   string `json:"id,omitempty" jsonschema:"description=Salesforce ApexClass Id to update"`
	Name string `json:"name,omitempty" jsonschema:"description=Apex class name (alternative to Id)"`
	Body string `json:"body" jsonschema:"description=Updated full Apex source code for the class. Replaces existing code."`
}

type GetApexTriggerPayload struct {
	Id   string `json:"id,omitempty" jsonschema:"description=Salesforce ApexTrigger Id"`
	Name string `json:"name,omitempty" jsonschema:"description=Apex trigger name (alternative to Id). Example: AccountTrigger"`
}

type CreateApexTriggerPayload struct {
	Name          string `json:"name" jsonschema:"description=API name for the new Apex trigger. Example: AccountTrigger"`
	TableEnumOrId string `json:"table_enum_or_id" jsonschema:"description=SObject API name the trigger fires on. Example: Account, Contact, Opportunity"`
	Body          string `json:"body" jsonschema:"description=Full Apex source code for the trigger. Example: trigger AccountTrigger on Account (before insert) { }"`
}

type UpdateApexTriggerPayload struct {
	Id            string `json:"id,omitempty" jsonschema:"description=Salesforce ApexTrigger Id to update"`
	Name          string `json:"name,omitempty" jsonschema:"description=Apex trigger name (alternative to Id)"`
	TableEnumOrId string `json:"table_enum_or_id,omitempty" jsonschema:"description=SObject API name. Recommended to ensure correct trigger resolution."`
	Body          string `json:"body" jsonschema:"description=Updated full Apex source code for the trigger."`
}

type ListApexLogsPayload struct {
	Limit int `json:"limit,omitempty" jsonschema:"description=Max number of log records to return. Defaults to 25. Example: 10"`
}

type GetApexLogBodyPayload struct {
	Id string `json:"id" jsonschema:"description=ApexLog record Id. Example: 07L2800000BnZzKEAV"`
}

type GetTraceFlagPayload struct {
	Id string `json:"id" jsonschema:"description=TraceFlag record Id. Example: 7tf2800000BpQrtAAF"`
}

type CreateTraceFlagPayload struct {
	TracedEntityId   string `json:"traced_entity_id,omitempty" jsonschema:"description=Id of the user or Apex class to trace."`
	TracedEntityName string `json:"traced_entity_name,omitempty" jsonschema:"description=Name of user or Apex class (alternative to Id)."`
	DebugLevelId     string `json:"debug_level_id" jsonschema:"description=Id of the DebugLevel to apply (e.g., 7dlXX000000...)."`
	LogType          string `json:"log_type" jsonschema:"description=Type of log. Example: DEVELOPER_LOG, USER_DEBUG, APEX_PROFILING."`
	ExpirationDate   string `json:"expiration_date" jsonschema:"description=ISO-8601 datetime string. Must be in the future. Example: 2026-12-31T23:59:59Z"`
}

type UpdateTraceFlagPayload struct {
	Id             string `json:"id" jsonschema:"description=TraceFlag record Id to update."`
	ExpirationDate string `json:"expiration_date" jsonschema:"description=New ISO-8601 expiry datetime."`
	DebugLevelId   string `json:"debug_level_id,omitempty" jsonschema:"description=Optional new DebugLevel Id."`
}

type RunApexTestsAsyncPayload struct {
	ClassNames     []string `json:"class_names,omitempty" jsonschema:"description=List of Apex test class names. Example: [MyClassTest, AnotherClassTest]"`
	SuiteNames     []string `json:"suite_names,omitempty" jsonschema:"description=List of Apex test suite names. Example: [MyTestSuite]"`
	MaxFailedTests int      `json:"max_failed_tests,omitempty" jsonschema:"description=Abort after N failures. Use -1 or omit for no limit. Example: 5"`
}

type GetApexTestResultPayload struct {
	AsyncApexJobId string `json:"async_apex_job_id" jsonschema:"description=Async job Id returned by async test run. Example: 707XX0000002bCoQAI"`
}

type RunApexTestsSyncPayload struct {
	ClassNames  []string `json:"class_names" jsonschema:"description=List of Apex test class names to run synchronously."`
	TestMethods []string `json:"test_methods,omitempty" jsonschema:"description=Optional specific test method names to run."`
}

type GetApexCodeCoveragePayload struct {
	ClassOrTriggerId string `json:"class_or_trigger_id,omitempty" jsonschema:"description=ApexClass or ApexTrigger Id."`
	ClassName        string `json:"class_name,omitempty" jsonschema:"description=Apex class name (alternative to Id)."`
	TriggerName      string `json:"trigger_name,omitempty" jsonschema:"description=Apex trigger name (alternative to Id)."`
}

type GetOrgWideCoveragePayload struct{}
