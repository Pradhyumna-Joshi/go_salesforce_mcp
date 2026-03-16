package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/config"
	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
	"github.com/pkg/browser"
)

func LoginToOrgHandler(ctx context.Context, e any) (string, error) {

	log.Println("CLIENT ID", config.Conf.Sfconfig.ClientID)
	log.Println("CLIENT SECRET", config.Conf.Sfconfig.ClientSecret)
	data := url.Values{}
	data.Add("response_type", "code")
	data.Add("client_id", config.Conf.Sfconfig.ClientID)
	data.Add("redirect_uri", "http://localhost:8080/callback")

	authURL := config.Conf.Sfconfig.LoginURL + "?" + data.Encode()

	// open this url in the browser
	if err := browser.OpenURL(authURL); err != nil {
		return "Please open this URL manually: " + authURL, nil
	}
	return "Browser opened for Salesforce login. Please complete authentication.", nil
}

func HandleCallBack(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")

	log.Println("code", code)

	data := url.Values{}
	data.Set("grant_type", config.Conf.Sfconfig.GrantType)
	data.Set("client_id", config.Conf.Sfconfig.ClientID)
	data.Set("client_secret", config.Conf.Sfconfig.ClientSecret)
	data.Set("redirect_uri", config.Conf.Sfconfig.RedirectURI)
	data.Set("code", code)

	resp, err := http.PostForm(config.Conf.Sfconfig.TokenURL, data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Token exchange failed: ", http.StatusInternalServerError)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}
	log.Println("Salesforce raw response:", string(body))

	var token models.TokenResponse
	if err := json.Unmarshal(body, &token); err != nil {
		http.Error(w, "Something went wrong "+err.Error(), http.StatusInternalServerError)
		return
	}

	config.Conf.Sfconfig.AccessToken = token.AccessToken
	config.Conf.Sfconfig.InstanceURL = token.InstanceURL

	log.Println("access token", token.AccessToken)
	log.Println("instance url", token.InstanceURL)

	fmt.Fprintln(w, "Login successful! ")

}

/*

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

func DisconnectSalesforceOrg(ctx context.Context, e Empty) (string, error) {

	cmd := exec.Command("sf", "org", "logout")

	body, err := cmd.Output()
	if err != nil {
		return "", err
	}

	log.Println(string(body))

	return string(body), nil
}
*/
