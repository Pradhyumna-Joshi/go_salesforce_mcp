package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Conf = LoadConfig()

type Config struct {
	ServerConfig ServerConfig
	Sfconfig     SalesforceConfig
}

type ServerConfig struct {
	Addr string
}

type SalesforceConfig struct {
	LoginURL     string
	TokenURL     string
	ClientID     string
	ClientSecret string
	GrantType    string
	RedirectURI  string
	AccessToken  string
	InstanceURL  string
}

func LoadConfig() *Config {

	if err := godotenv.Load("/Users/pradhyumnajoshi/go-dev/go_salesforce_mcp/.env"); err != nil {
		log.Println(".env not found")
	}

	log.Println("CLIENT ID", os.Getenv("SF_CLIENT_ID"))

	return &Config{
		ServerConfig: ServerConfig{
			Addr: os.Getenv("MCP_PORT_ADDR"),
		},
		Sfconfig: SalesforceConfig{
			LoginURL:     os.Getenv("SF_LOGIN_URL"),
			ClientID:     os.Getenv("SF_CLIENT_ID"),
			ClientSecret: os.Getenv("SF_CLIENT_SECRET"),
			GrantType:    os.Getenv("SF_GRANT_TYPE"),
			RedirectURI:  os.Getenv("SF_REDIRECT_URI"),
			TokenURL:     os.Getenv("SF_TOKEN_URL"),
		},
	}
}
