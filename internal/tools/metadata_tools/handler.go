package tools

import (
	"context"
	"fmt"

	repo "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/repository"
)

func DescribeSObject(ctx context.Context, input SObjectRequest) (string, error) {

	url := fmt.Sprintf("/services/data/v61.0/sobjects/%s/describe", input.Object)

	body, err := repo.SalesforceRequest("GET", url, nil)

	if err != nil {
		return "", err
	}
	return string(body), nil
}

func ListSObjects(ctx context.Context, e any) (string, error) {

	url := "/services/data/v61.0/sobjects"

	body, err := repo.SalesforceRequest("GET", url, nil)

	if err != nil {
		return "", nil
	}
	return string(body), nil

}
