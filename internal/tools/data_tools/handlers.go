package tools

import (
	"context"
	"fmt"
	"log"
	"net/url"

	repo "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/repository"
)

func QuerySObject(ctx context.Context, query QueryPayload) (string, error) {
	log.Printf("[MCP Tool] QuerySObject: Executing SOQL: %s", query.Soql)

	endpoint := fmt.Sprintf(
		"/services/data/v59.0/query?q=%s",
		url.QueryEscape(query.Soql),
	)

	body, err := repo.SalesforceRequest("GET", endpoint, nil)
	if err != nil {
		log.Printf("[MCP Tool Error] QuerySObject failed: %v", err)
		return "", err
	}

	log.Printf("[MCP Tool Success] QuerySObject returned: %s", string(body))
	return string(body), nil
}

func CreateRecord(ctx context.Context, input CreateRecordPayload) (string, error) {
	log.Printf("[MCP Tool] CreateRecord: Creating %s record with fields: %+v", input.Object, input.Fields)

	endpoint := fmt.Sprintf(
		"/services/data/v61.0/sobjects/%s",
		url.QueryEscape(input.Object),
	)

	body, err := repo.SalesforceRequest("POST", endpoint, input.Fields)
	if err != nil {
		log.Printf("[MCP Tool Error] CreateRecord failed for %s: %v", input.Object, err)
		return "", err
	}

	log.Printf("[MCP Tool Success] CreateRecord created %s successfully: %s", input.Object, string(body))
	return string(body), nil
}

func UpdateRecord(ctx context.Context, input UpdateRecordPayload) (string, error) {
	log.Printf("[MCP Tool] UpdateRecord: Updating %s (ID: %s) with fields: %+v", input.Object, input.Id, input.Fields)

	endpoint := fmt.Sprintf(
		"/services/data/v61.0/sobjects/%s/%s",
		url.QueryEscape(input.Object),
		input.Id,
	)

	body, err := repo.SalesforceRequest("PATCH", endpoint, input.Fields)
	if err != nil {
		log.Printf("[MCP Tool Error] UpdateRecord failed for %s ID %s: %v", input.Object, input.Id, err)
		return "", err
	}

	log.Printf("[MCP Tool Success] UpdateRecord modified %s ID %s successfully", input.Object, input.Id)
	return string(body), nil
}
