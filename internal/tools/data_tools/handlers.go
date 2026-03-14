package tools

import (
	"context"
	"fmt"
	"net/url"

	repo "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/repository"
)

func QuerySObject(ctx context.Context, query QueryPayload) (string, error) {

	url := fmt.Sprintf(
		"/services/data/v59.0/query?q=%s",
		url.QueryEscape(query.Soql),
	)

	body, err := repo.SalesforceRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	return string(body), nil

}

func CreateRecord(ctx context.Context, input CreateRecordPayload) (string, error) {

	url := fmt.Sprintf(
		"/services/data/v61.0/sobjects/%s",
		url.QueryEscape(input.Object),
	)

	body, err := repo.SalesforceRequest("POST", url, input.Fields)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func UpdateRecord(ctx context.Context, input UpdateRecordPayload) (string, error) {

	url := fmt.Sprintf(
		"/services/data/v61.0/sobjects/%s/%s",
		url.QueryEscape(input.Object),
		input.Id,
	)

	body, err := repo.SalesforceRequest("PATCH", url, input.Fields)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
