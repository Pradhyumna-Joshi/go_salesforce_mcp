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
	log.Printf("[MCP Tool] ApexRestGet: Path=%s, Params=%v", input.Path, input.QueryParams)
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
	if err != nil {
		log.Printf("[MCP Tool Error] ApexRestGet failed: %v", err)
		return "", err
	}
	log.Printf("[MCP Tool Success] ApexRestGet: Received %d bytes", len(body))
	return string(body), nil
}

func ApexRestPost(ctx context.Context, input ApexRestPostPayload) (string, error) {
	log.Printf("[MCP Tool] ApexRestPost: Path=%s", input.Path)
	path := ensureLeadingSlash(input.Path)
	endpoint := "/services/apexrest" + path

	body, err := repo.SalesforceRequest("POST", endpoint, input.Body)
	if err != nil {
		log.Printf("[MCP Tool Error] ApexRestPost failed: %v", err)
		return "", err
	}
	return string(body), nil
}

func ApexRestPatch(ctx context.Context, input ApexRestPatchPayload) (string, error) {
	log.Printf("[MCP Tool] ApexRestPatch: Path=%s", input.Path)
	path := ensureLeadingSlash(input.Path)
	endpoint := "/services/apexrest" + path

	body, err := repo.SalesforceRequest("PATCH", endpoint, input.Body)
	if err != nil {
		log.Printf("[MCP Tool Error] ApexRestPatch failed: %v", err)
		return "", err
	}
	return string(body), nil
}

func GetApexClass(ctx context.Context, input GetApexClassPayload) (string, error) {
	log.Printf("[MCP Tool] GetApexClass: ID=%s, Name=%s", input.Id, input.Name)
	id := input.Id

	if id == "" && input.Name != "" {
		soql := fmt.Sprintf("SELECT Id FROM ApexClass WHERE Name = '%s' LIMIT 1", input.Name)
		res, err := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
		if err != nil {
			return "", err
		}
		id = extractIdFromQuery(res)
		log.Printf("[MCP Tool] GetApexClass: Resolved Name %s to ID %s", input.Name, id)
	}

	if id == "" {
		return "", fmt.Errorf("either id or name must be provided")
	}

	endpoint := fmt.Sprintf("/services/data/%s/tooling/sobjects/ApexClass/%s", ApiVersion, id)
	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	return string(body), err
}

func CreateApexClass(ctx context.Context, input CreateApexClassPayload) (string, error) {
	log.Printf("[MCP Tool] CreateApexClass: Name=%s", input.Name)
	endpoint := fmt.Sprintf("/services/data/%s/tooling/sobjects/ApexClass", ApiVersion)
	payload := map[string]any{"Name": input.Name, "Body": input.Body}

	body, err := repo.SalesforceRequest("POST", endpoint, payload)
	if err != nil {
		log.Printf("[MCP Tool Error] CreateApexClass failed: %v", err)
		return "", err
	}
	return string(body), nil
}

func UpdateApexClass(ctx context.Context, input UpdateApexClassPayload) (string, error) {
	log.Printf("[MCP Tool] UpdateApexClass: ID=%s, Name=%s", input.Id, input.Name)
	id := input.Id

	if id == "" && input.Name != "" {
		soql := fmt.Sprintf("SELECT Id FROM ApexClass WHERE Name = '%s' LIMIT 1", input.Name)
		res, _ := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
		id = extractIdFromQuery(res)
	}

	endpoint := fmt.Sprintf("/services/data/%s/tooling/sobjects/ApexClass/%s", ApiVersion, id)
	payload := map[string]any{"Body": input.Body}
	body, err := repo.SalesforceRequest("PATCH", endpoint, payload) // Note: Tooling API often uses PATCH or PUT
	return string(body), err
}

func GetApexTrigger(ctx context.Context, input GetApexTriggerPayload) (string, error) {
	log.Printf("[MCP Tool] GetApexTrigger: ID=%s, Name=%s", input.Id, input.Name)
	id := input.Id

	if id == "" && input.Name != "" {
		soql := fmt.Sprintf("SELECT Id FROM ApexTrigger WHERE Name = '%s' LIMIT 1", input.Name)
		res, _ := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
		id = extractIdFromQuery(res)
	}

	endpoint := fmt.Sprintf("/services/data/%s/tooling/sobjects/ApexTrigger/%s", ApiVersion, id)
	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	return string(body), err
}

func ApexExecuteAnonymous(ctx context.Context, input ApexExecuteAnonymousPayload) (string, error) {
	log.Printf("[MCP Tool] ApexExecuteAnonymous: Executing script snippet")
	endpoint := fmt.Sprintf("/services/data/%s/tooling/executeAnonymous?anonymousBody=%s",
		ApiVersion, url.QueryEscape(input.Body))

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	if err != nil {
		log.Printf("[MCP Tool Error] ApexExecuteAnonymous failed: %v", err)
		return "", err
	}
	return string(body), nil
}

func ListApexLogs(ctx context.Context, input ListApexLogsPayload) (string, error) {
	limit := input.Limit
	if limit <= 0 {
		limit = 25
	}
	log.Printf("[MCP Tool] ListApexLogs: Requesting %d latest logs", limit)

	soql := fmt.Sprintf("SELECT Id,Application,DurationMilliseconds,Operation,Request,StartTime,Status FROM ApexLog ORDER BY StartTime DESC LIMIT %d", limit)
	return ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
}

func GetApexLogBody(ctx context.Context, input GetApexLogBodyPayload) (string, error) {
	log.Printf("[MCP Tool] GetApexLogBody: ID=%s", input.Id)
	endpoint := fmt.Sprintf("/services/data/%s/tooling/sobjects/ApexLog/%s/Body", ApiVersion, input.Id)
	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}
	log.Printf("[MCP Tool Success] GetApexLogBody: Retrieved %d bytes", len(body))
	return string(body), nil
}

func RunApexTestsAsync(ctx context.Context, input RunApexTestsAsyncPayload) (string, error) {
	log.Printf("[MCP Tool] RunApexTestsAsync: Classes=%v, Suites=%v", input.ClassNames, input.SuiteNames)
	endpoint := fmt.Sprintf("/services/data/%s/tooling/runTestsAsynchronous", ApiVersion)

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
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetApexCodeCoverage(ctx context.Context, input GetApexCodeCoveragePayload) (string, error) {
	log.Printf("[MCP Tool] GetApexCodeCoverage: Checking coverage for target")
	id := input.ClassOrTriggerId

	if id == "" {
		if input.ClassName != "" {
			soql := fmt.Sprintf("SELECT Id FROM ApexClass WHERE Name = '%s' LIMIT 1", input.ClassName)
			res, _ := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
			id = extractIdFromQuery(res)
		} else if input.TriggerName != "" {
			soql := fmt.Sprintf("SELECT Id FROM ApexTrigger WHERE Name = '%s' LIMIT 1", input.TriggerName)
			res, _ := ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
			id = extractIdFromQuery(res)
		}
	}

	if id == "" {
		return "", fmt.Errorf("could not resolve Id from provided names")
	}

	soql := fmt.Sprintf("SELECT Id,ApexClassOrTrigger.Name,NumLinesCovered,NumLinesUncovered FROM ApexCodeCoverageAggregate WHERE ApexClassOrTriggerId = '%s'", id)
	return ApexToolingQuery(ctx, ApexToolingQueryPayload{Soql: soql})
}

func ApexToolingQuery(ctx context.Context, input ApexToolingQueryPayload) (string, error) {
	log.Printf("[MCP Tool] ApexToolingQuery: %s", input.Soql)
	endpoint := fmt.Sprintf("/services/data/%s/tooling/query?q=%s", ApiVersion, url.QueryEscape(input.Soql))
	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	return string(body), err
}

func ensureLeadingSlash(path string) string {
	if !strings.HasPrefix(path, "/") {
		return "/" + path
	}
	return path
}

func extractIdFromQuery(response string) string {
	type record struct {
		Id string `json:"Id"`
	}
	type result struct {
		Records []record `json:"records"`
	}
	var r result
	if err := json.Unmarshal([]byte(response), &r); err != nil {
		return ""
	}
	if len(r.Records) > 0 {
		return r.Records[0].Id
	}
	return ""
}
