package common

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/config"
)

var SFClient = &http.Client{}

func SalesforceRequest(method string, path string, data any) ([]byte, error) {

	url := config.Conf.Sfconfig.InstanceURL + path

	var body io.Reader
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)

	}

	r, err := http.NewRequest(method, url, body)
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
