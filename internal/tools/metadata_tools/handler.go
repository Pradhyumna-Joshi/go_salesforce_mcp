package tools

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/config"
)

var SFClient = &http.Client{}

// Generic salesforce request generator
func SalesforceMetadataRequest(method string, path string, data io.Reader) ([]byte, error) {

	url := config.Conf.Sfconfig.InstanceURL + path

	r, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	r.Header.Set("Authorization", "Bearer "+config.Conf.Sfconfig.AccessToken)

	resp, err := SFClient.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)

}

func DescribeSObject(ctx context.Context, input SObjectRequest) (string, error) {

	path := fmt.Sprintf("/services/data/v61.0/sobjects/%s/describe", input.Object)

	body, err := SalesforceMetadataRequest("GET", path, nil)

	if err != nil {
		return "", err
	}
	return string(body), nil
}

func ListSObjects(ctx context.Context, e any) (string, error) {

	path := "/services/data/v61.0/sobjects"

	body, err := SalesforceMetadataRequest("GET", path, nil)

	if err != nil {
		return "", nil
	}
	return string(body), nil

}
