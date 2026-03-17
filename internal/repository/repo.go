package repository

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/config"
)

var SFClient = &http.Client{}

func SalesforceRequest(httpMethod string, path string, reqBody any) ([]byte, error) {

	url := config.Conf.Sfconfig.InstanceURL + path

	var body io.Reader
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)
	}

	httpRequest, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Authorization", "Bearer "+config.Conf.Sfconfig.AccessToken)
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := SFClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()
	return io.ReadAll(httpResponse.Body)
}
