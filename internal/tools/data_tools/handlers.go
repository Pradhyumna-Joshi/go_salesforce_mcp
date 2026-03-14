package tools

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/config"
)

func QuerySObjectHandler(ctx context.Context, query QueryRequest) (string, error) {

	url := fmt.Sprintf(
		"%s/services/data/v59.0/query?q=%s",
		config.Conf.Sfconfig.InstanceURL,
		url.QueryEscape(query.Soql),
	)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("Authorization", "Bearer "+config.Conf.Sfconfig.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	log.Println("RESPONSE==>", string(body))
	return string(body), nil

}
