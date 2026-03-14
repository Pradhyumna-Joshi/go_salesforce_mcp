package tools

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/config"
)

func LoginToOrgHandler(ctx context.Context, e any) (string, error) {

	cmd := exec.Command("sf", "org", "login", "web", "--set-default", "--json")

	body, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	LoadSalesforceSession()

	log.Println(config.Conf.Sfconfig.AccessToken)
	log.Println(config.Conf.Sfconfig.InstanceURL)

	return string(body), nil
}

func LoadSalesforceSession() error {

	cmd := exec.Command("sf", "org", "display", "--json")

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	var org OrgResult

	err = json.Unmarshal(out, &org)
	if err != nil {
		return err
	}

	config.Conf.Sfconfig.AccessToken = org.Result.AccessToken
	config.Conf.Sfconfig.InstanceURL = org.Result.InstanceURL
	config.Conf.Sfconfig.UserName = org.Result.Username

	return nil
}

func ListAllOrgs(ctx context.Context, e Empty) (string, error) {

	cmd := exec.Command("sf", "org", "list")

	body, err := cmd.Output()
	if err != nil {
		return "", err
	}

	log.Println(string(body))

	return string(body), nil
}

func OpenSalesforceOrg(ctx context.Context, e Empty) (string, error) {

	cmd := exec.Command("sf", "org", "open")

	body, err := cmd.Output()
	if err != nil {
		return "", err
	}

	log.Println(string(body))

	return string(body), nil
}
