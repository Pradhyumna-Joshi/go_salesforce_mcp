package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	repo "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/repository"
)

func ApexRestGet(ctx context.Context, input ApexRestGetPayload) (string, error) {
	log.Printf("ApexRestGet Input received: %+v", input)

	path := ensureLeadingSlash(input.Path)
	endpoint := "/services/apexrest" + path

	if len(input.QueryParams) > 0 {
		params := url.Values{}
		for k, v := range input.QueryParams {
			params.Set(k, v)
		}
		endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	log.Println("ApexRestGet - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ApexRestPost(ctx context.Context, input ApexRestPostPayload) (string, error) {
	log.Printf("ApexRestPost Input received: %+v", input)

	path := ensureLeadingSlash(input.Path)
	endpoint := "/services/apexrest" + path

	body, err := repo.SalesforceRequest("POST", endpoint, input.Body)
	log.Println("ApexRestPost - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ApexRestPatch(ctx context.Context, input ApexRestPatchPayload) (string, error) {
	log.Printf("ApexRestPatch Input received: %+v", input)

	path := ensureLeadingSlash(input.Path)
	endpoint := "/services/apexrest" + path

	body, err := repo.SalesforceRequest("PATCH", endpoint, input.Body)
	log.Println("ApexRestPatch - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ApexExecuteAnonymous(ctx context.Context, input ApexExecuteAnonymousPayload) (string, error) {
	log.Printf("ApexExecuteAnonymous Input received: %+v", input)

	endpoint := fmt.Sprintf(
		"/services/data/v61.0/tooling/executeAnonymous?anonymousBody=%s",
		url.QueryEscape(input.Body),
	)

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	log.Println("ApexExecuteAnonymous - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ApexToolingQuery(ctx context.Context, input ApexToolingQueryPayload) (string, error) {
	log.Printf("ApexToolingQuery Input received: %+v", input)

	endpoint := fmt.Sprintf(
		"/services/data/v61.0/tooling/query?q=%s",
		url.QueryEscape(input.Soql),
	)

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	log.Println("ApexToolingQuery - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetApexClass(ctx context.Context, input GetApexClassPayload) (string, error) {
	log.Printf("GetApexClass Input received: %+v", input)

	id := input.Id

	if id == "" && input.Name != "" {
		soql := fmt.Sprintf("SELECT Id FROM ApexClass WHERE Name = '%s' LIMIT 1", input.Name)
		res, err := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
		if err != nil {
			return "", err
		}
		id = extractIdFromQuery(res)
	}

	if id == "" {
		return "", fmt.Errorf("either id or name must be provided")
	}

	endpoint := fmt.Sprintf("/services/data/%s/tooling/sobjects/ApexClass/%s", apiVersion, id)

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CreateApexClass(ctx context.Context, input CreateApexClassPayload) (string, error) {
	log.Printf("CreateApexClass Input received: %+v", input)

	endpoint := "/services/data/v61.0/tooling/sobjects/ApexClass"

	payload := map[string]any{
		"Name": input.Name,
		"Body": input.Body,
	}

	body, err := repo.SalesforceRequest("POST", endpoint, payload)
	log.Println("CreateApexClass - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func UpdateApexClass(ctx context.Context, input UpdateApexClassPayload) (string, error) {
	log.Printf("UpdateApexClass Input received: %+v", input)

	id := input.Id

	if id == "" && input.Name != "" {
		soql := fmt.Sprintf("SELECT Id FROM ApexClass WHERE Name = '%s' LIMIT 1", input.Name)
		res, err := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
		if err != nil {
			return "", err
		}
		id = extractIdFromQuery(res)
	}

	if id == "" {
		return "", fmt.Errorf("either id or name must be provided")
	}

	endpoint := fmt.Sprintf("/services/data/%s/tooling/sobjects/ApexClass/%s", apiVersion, id)

	payload := map[string]any{
		"Body": input.Body,
	}

	body, err := repo.SalesforceRequest("PATCH", endpoint, payload)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetApexTrigger(ctx context.Context, input GetApexTriggerPayload) (string, error) {
	log.Printf("GetApexTrigger Input received: %+v", input)

	id := input.Id

	if id == "" && input.Name != "" {
		soql := fmt.Sprintf("SELECT Id FROM ApexTrigger WHERE Name = '%s' LIMIT 1", input.Name)
		res, err := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
		if err != nil {
			return "", err
		}
		id = extractIdFromQuery(res)
	}

	if id == "" {
		return "", fmt.Errorf("either id or name must be provided")
	}

	endpoint := fmt.Sprintf("/services/data/%s/tooling/sobjects/ApexTrigger/%s", apiVersion, id)

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CreateApexTrigger(ctx context.Context, input CreateApexTriggerPayload) (string, error) {
	log.Printf("CreateApexTrigger Input received: %+v", input)

	endpoint := "/services/data/v61.0/tooling/sobjects/ApexTrigger"

	payload := map[string]any{
		"Name":          input.Name,
		"TableEnumOrId": input.TableEnumOrId,
		"Body":          input.Body,
	}

	body, err := repo.SalesforceRequest("POST", endpoint, payload)
	log.Println("CreateApexTrigger - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func UpdateApexTrigger(ctx context.Context, input UpdateApexTriggerPayload) (string, error) {
	log.Printf("UpdateApexTrigger Input received: %+v", input)

	id := input.Id

	if id == "" && input.Name != "" {
		soql := fmt.Sprintf("SELECT Id FROM ApexTrigger WHERE Name = '%s' LIMIT 1", input.Name)
		res, err := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
		if err != nil {
			return "", err
		}
		id = extractIdFromQuery(res)
	}

	if id == "" {
		return "", fmt.Errorf("either id or name must be provided")
	}

	endpoint := fmt.Sprintf("/services/data/%s/tooling/sobjects/ApexTrigger/%s", apiVersion, id)

	payload := map[string]any{
		"Body": input.Body,
	}

	body, err := repo.SalesforceRequest("PATCH", endpoint, payload)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ListApexLogs(ctx context.Context, input ListApexLogsPayload) (string, error) {
	log.Printf("ListApexLogs Input received: %+v", input)

	limit := input.Limit
	if limit <= 0 {
		limit = 25
	}

	soql := fmt.Sprintf(
		"SELECT Id,Application,DurationMilliseconds,Location,LogLength,LogUserId,Operation,Request,StartTime,Status FROM ApexLog ORDER BY StartTime DESC LIMIT %d",
		limit,
	)
	endpoint := fmt.Sprintf(
		"/services/data/v61.0/tooling/query?q=%s",
		url.QueryEscape(soql),
	)

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	log.Println("ListApexLogs - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetApexLogBody(ctx context.Context, input GetApexLogBodyPayload) (string, error) {
	log.Printf("GetApexLogBody Input received: %+v", input)

	endpoint := fmt.Sprintf(
		"/services/data/v61.0/tooling/sobjects/ApexLog/%s/Body",
		input.Id,
	)

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	log.Println("GetApexLogBody - Response length:", len(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetTraceFlag(ctx context.Context, input GetTraceFlagPayload) (string, error) {
	log.Printf("GetTraceFlag Input received: %+v", input)

	endpoint := fmt.Sprintf(
		"/services/data/v61.0/tooling/sobjects/TraceFlag/%s",
		input.Id,
	)

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	log.Println("GetTraceFlag - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CreateTraceFlag(ctx context.Context, input CreateTraceFlagPayload) (string, error) {
	log.Printf("CreateTraceFlag Input received: %+v", input)

	entityId := input.TracedEntityId

	if entityId == "" && input.TracedEntityName != "" {
		soql := fmt.Sprintf("SELECT Id FROM User WHERE Name = '%s' LIMIT 1", input.TracedEntityName)
		res, err := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
		if err != nil {
			return "", err
		}
		entityId = extractIdFromQuery(res)
	}

	if entityId == "" {
		return "", fmt.Errorf("traced_entity_id or traced_entity_name required")
	}

	endpoint := fmt.Sprintf("/services/data/%s/tooling/sobjects/TraceFlag", apiVersion)

	payload := map[string]any{
		"TracedEntityId": entityId,
		"DebugLevelId":   input.DebugLevelId,
		"LogType":        input.LogType,
		"ExpirationDate": input.ExpirationDate,
	}

	body, err := repo.SalesforceRequest("POST", endpoint, payload)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func UpdateTraceFlag(ctx context.Context, input UpdateTraceFlagPayload) (string, error) {
	log.Printf("UpdateTraceFlag Input received: %+v", input)

	endpoint := fmt.Sprintf(
		"/services/data/v61.0/tooling/sobjects/TraceFlag/%s",
		input.Id,
	)

	payload := map[string]any{
		"ExpirationDate": input.ExpirationDate,
	}
	if input.DebugLevelId != "" {
		payload["DebugLevelId"] = input.DebugLevelId
	}

	body, err := repo.SalesforceRequest("PATCH", endpoint, payload)
	log.Println("UpdateTraceFlag - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func RunApexTestsAsync(ctx context.Context, input RunApexTestsAsyncPayload) (string, error) {
	log.Printf("RunApexTestsAsync Input received: %+v", input)

	endpoint := "/services/data/v61.0/tooling/runTestsAsynchronous"

	payload := map[string]any{}
	if len(input.ClassNames) > 0 {
		payload["classNames"] = strings.Join(input.ClassNames, ",")
	}
	if len(input.SuiteNames) > 0 {
		payload["suiteNames"] = strings.Join(input.SuiteNames, ",")
	}
	if input.MaxFailedTests != 0 {
		payload["maxFailedTests"] = input.MaxFailedTests
	}

	body, err := repo.SalesforceRequest("POST", endpoint, payload)
	log.Println("RunApexTestsAsync - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetApexTestResult(ctx context.Context, input GetApexTestResultPayload) (string, error) {
	log.Printf("GetApexTestResult Input received: %+v", input)

	soql := fmt.Sprintf(
		"SELECT Id,ApexClass.Name,MethodName,Outcome,Message,StackTrace FROM ApexTestResult WHERE AsyncApexJobId = '%s'",
		input.AsyncApexJobId,
	)
	endpoint := fmt.Sprintf(
		"/services/data/v61.0/tooling/query?q=%s",
		url.QueryEscape(soql),
	)

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	log.Println("GetApexTestResult - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func RunApexTestsSync(ctx context.Context, input RunApexTestsSyncPayload) (string, error) {
	log.Printf("RunApexTestsSync Input received: %+v", input)

	endpoint := "/services/data/v61.0/tooling/runTestsSynchronous"

	payload := map[string]any{
		"classNames": strings.Join(input.ClassNames, ","),
	}
	if len(input.TestMethods) > 0 {
		payload["testMethods"] = strings.Join(input.TestMethods, ",")
	}

	body, err := repo.SalesforceRequest("POST", endpoint, payload)
	log.Println("RunApexTestsSync - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetApexCodeCoverage(ctx context.Context, input GetApexCodeCoveragePayload) (string, error) {
	log.Printf("GetApexCodeCoverage Input received: %+v", input)

	id := input.ClassOrTriggerId

	if id == "" {
		if input.ClassName != "" {
			soql := fmt.Sprintf("SELECT Id FROM ApexClass WHERE Name = '%s' LIMIT 1", input.ClassName)
			res, err := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
			if err != nil {
				return "", err
			}
			id = extractIdFromQuery(res)
		} else if input.TriggerName != "" {
			soql := fmt.Sprintf("SELECT Id FROM ApexTrigger WHERE Name = '%s' LIMIT 1", input.TriggerName)
			res, err := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
			if err != nil {
				return "", err
			}
			id = extractIdFromQuery(res)
		}
	}

	if id == "" {
		return "", fmt.Errorf("provide class_or_trigger_id, class_name, or trigger_name")
	}

	soql := fmt.Sprintf(
		"SELECT Id,ApexClassOrTriggerId,ApexClassOrTrigger.Name,NumLinesCovered,NumLinesUncovered,Coverage FROM ApexCodeCoverageAggregate WHERE ApexClassOrTriggerId = '%s'",
		id,
	)

	return ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
}

func GetOrgWideCoverage(ctx context.Context, _ GetOrgWideCoveragePayload) (string, error) {
	log.Println("GetOrgWideCoverage called")

	soql := "SELECT PercentCovered FROM ApexOrgWideCoverage"
	endpoint := fmt.Sprintf(
		"/services/data/v61.0/tooling/query?q=%s",
		url.QueryEscape(soql),
	)

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	log.Println("GetOrgWideCoverage - Response:", string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ensureLeadingSlash(path string) string {
	if !strings.HasPrefix(path, "/") {
		return "/" + path
	}
	return path
}

func extractIdFromQuery(response string) string {
	// minimal parsing (you can improve later)
	type record struct {
		Id string `json:"Id"`
	}
	type result struct {
		Records []record `json:"records"`
	}

	var r result
	_ = json.Unmarshal([]byte(response), &r)

	if len(r.Records) > 0 {
		return r.Records[0].Id
	}
	return ""
}
