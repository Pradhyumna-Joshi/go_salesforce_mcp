package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Conf = LoadConfig()

type Config struct {
	ServerConfig ServerConfig
	Sfconfig    SalesforceConfig
}

type ServerConfig struct {
	Addr string
}

type SalesforceConfig struct {
	AccessToken string
	InstanceURL string
	UserName    string
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
	}
}
