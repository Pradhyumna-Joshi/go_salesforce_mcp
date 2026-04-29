package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	repo "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/repository"
)

func GetSObjectFields(ctx context.Context, input SObjectRequest) (string, error) {
	log.Printf("[MCP Tool] GetSObjectFields: Describing object: %s", input.Object)

	endpoint := fmt.Sprintf("/services/data/v61.0/sobjects/%s/describe", input.Object)
	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	if err != nil {
		log.Printf("[MCP Tool Error] GetSObjectFields failed for %s: %v", input.Object, err)
		return "", err
	}

	var fullResp struct {
		Fields []struct {
			Name  string `json:"name"`
			Label string `json:"label"`
			Type  string `json:"type"`
		} `json:"fields"`
	}

	if err := json.Unmarshal(body, &fullResp); err != nil {
		log.Printf("[MCP Tool Error] GetSObjectFields: JSON unmarshal error for %s: %v", input.Object, err)
		return "", err
	}

	res, _ := json.Marshal(fullResp.Fields)
	log.Printf("[MCP Tool Success] GetSObjectFields: Found %d fields for %s", len(fullResp.Fields), input.Object)
	return string(res), nil
}

func GetFieldDetails(ctx context.Context, input FieldInfoRequest) (string, error) {
	log.Printf("[MCP Tool] GetFieldDetails: Fetching details for %s.%s", input.Object, input.Field)

	endpoint := fmt.Sprintf("/services/data/v61.0/sobjects/%s/describe", input.Object)
	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	if err != nil {
		log.Printf("[MCP Tool Error] GetFieldDetails failed for %s: %v", input.Object, err)
		return "", err
	}

	var fullResp struct {
		Fields []map[string]interface{} `json:"fields"`
	}
	json.Unmarshal(body, &fullResp)

	for _, f := range fullResp.Fields {
		if f["name"] == input.Field {
			res, _ := json.Marshal(f)
			log.Printf("[MCP Tool Success] GetFieldDetails: Found details for %s.%s", input.Object, input.Field)
			return string(res), nil
		}
	}

	log.Printf("[MCP Tool Warning] GetFieldDetails: Field %s not found on object %s", input.Field, input.Object)
	return "Field not found", nil
}

func GetChildRelationships(ctx context.Context, input SObjectRequest) (string, error) {
	log.Printf("[MCP Tool] GetChildRelationships: Fetching relationships for %s", input.Object)

	endpoint := fmt.Sprintf("/services/data/v61.0/sobjects/%s/describe", input.Object)
	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	if err != nil {
		log.Printf("[MCP Tool Error] GetChildRelationships failed for %s: %v", input.Object, err)
		return "", err
	}

	var fullResp struct {
		Relationships []interface{} `json:"childRelationships"`
	}
	json.Unmarshal(body, &fullResp)
	res, _ := json.Marshal(fullResp.Relationships)

	log.Printf("[MCP Tool Success] GetChildRelationships: Retrieved relationships for %s", input.Object)
	return string(res), nil
}

func GetRecordTypes(ctx context.Context, input SObjectRequest) (string, error) {
	log.Printf("[MCP Tool] GetRecordTypes: Fetching Record Types for %s", input.Object)

	endpoint := fmt.Sprintf("/services/data/v61.0/sobjects/%s/describe", input.Object)
	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	if err != nil {
		log.Printf("[MCP Tool Error] GetRecordTypes failed for %s: %v", input.Object, err)
		return "", err
	}

	var fullResp struct {
		RecordTypes []interface{} `json:"recordTypeInfos"`
	}
	json.Unmarshal(body, &fullResp)
	res, _ := json.Marshal(fullResp.RecordTypes)

	log.Printf("[MCP Tool Success] GetRecordTypes: Retrieved record types for %s", input.Object)
	return string(res), nil
}

func GlobalSearch(ctx context.Context, input SearchRequest) (string, error) {
	log.Printf("[MCP Tool] GlobalSearch: Executing SOSL search: %s", input.Query)

	endpoint := fmt.Sprintf("/services/data/v61.0/search/?q=FIND+{%s}", input.Query)
	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	if err != nil {
		log.Printf("[MCP Tool Error] GlobalSearch failed: %v", err)
		return "", err
	}

	log.Printf("[MCP Tool Success] GlobalSearch: Search completed")
	return string(body), nil
}

func GetSystemContext(ctx context.Context, input EmptyRequest) (string, error) {
	log.Printf("[MCP Tool] GetSystemContext: Fetching Org Limits")

	endpoint := "/services/data/v61.0/limits"
	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	if err != nil {
		log.Printf("[MCP Tool Error] GetSystemContext failed: %v", err)
		return "", err
	}

	log.Printf("[MCP Tool Success] GetSystemContext: Limits retrieved")
	return string(body), nil
}
