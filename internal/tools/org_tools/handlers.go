package tools

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/config"
)

func connectToOrg(ctx context.Context, inputReq ConnectToOrgInput) (string, error) {

	if inputReq.OrgType == "" {
		return "Invalid orgType, Must be 'production' or 'sandbox'.", nil
	}

	if inputReq.Alias == "" {
		return "Invalid alias: Alias cannot be empty.", nil
	}

	var cmd *exec.Cmd
	prodUrl := "https://login.salesforce.com"
	sandboxUrl := "https://test.salesforce.com"

	switch inputReq.OrgType {
	case "production":
		cmd = exec.Command("sf", "org", "login", "web", "--set-default", "--alias", inputReq.Alias, "--instance-url", prodUrl, "--json")
	case "sandbox":
		cmd = exec.Command("sf", "org", "login", "web", "--set-default", "--alias", inputReq.Alias, "--instance-url", sandboxUrl, "--json")
	default:
		return "Invalid orgType, Must be 'production' or 'sandbox'.", nil
	}

	resBody, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	LoadSalesforceSession()

	log.Println(config.Conf.Sfconfig.AccessToken)
	log.Println(config.Conf.Sfconfig.InstanceURL)

	return string(resBody), nil
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

func listOrgs(ctx context.Context, e any) (string, error) {

	cmd := exec.Command("sf", "org", "list", "--json")

	resBody, err := cmd.Output()
	if err != nil {
		return "", err
	}

	log.Println(string(resBody))

	return string(resBody), nil
}

func openOrg(ctx context.Context, inputReq OpenOrgInput) (string, error) {
	log.Printf("openOrg Input received: %+v", inputReq)
	if inputReq.Alias == "" {
		return "Invalid alias: Run list_orgs to see available orgs.", nil
	}

	cmd := exec.Command("sf", "org", "open", "--target-org", inputReq.Alias)

	resBody, err := cmd.Output()
	if err != nil {
		return "", err
	}

	log.Println(string(resBody))

	return string(resBody), nil
}

func disconnectOrg(ctx context.Context, inputReq DisconnectOrgInput) (string, error) {
	log.Printf("disconnectOrg Input received: %+v", inputReq)
	if inputReq.Alias == "" {
		return "Invalid alias: Run list_orgs to see available orgs.", nil
	}
	cmd := exec.Command("sf", "org", "logout", "--target-org", inputReq.Alias, "--no-prompt", "--json")

	resBody, err := cmd.Output()

	if err != nil {
		return "", err
	}

	log.Println(string(resBody))

	return string(resBody), nil
}
