package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
	apex_tools "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/tools/apex_tools"
	data_tools "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/tools/data_tools"
	lwc_aura_tools "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/tools/lwc_aura_tools"
	mdt_tools "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/tools/metadata_tools"
	org_tools "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/tools/org_tools"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	PORT := ":" + getEnv("SF_MCP_PORT", "9000")

	lstTools := make([]models.MCPTool, 0)
	lstTools = append(lstTools, org_tools.ORG_TOOLS...)
	lstTools = append(lstTools, data_tools.DATA_TOOLS...)
	lstTools = append(lstTools, mdt_tools.METADATA_TOOLS...)
	lstTools = append(lstTools, apex_tools.APEX_TOOLS...)
	lstTools = append(lstTools, lwc_aura_tools.UI_COMPONENTS_TOOLS...)

	lstToolCategories := make([]models.ToolCategory, 0)
	lstToolCategories = append(lstToolCategories, models.ToolCategory{
		Name:        "OrgManagementTools",
		Description: "Tools for managing Salesforce org connections, including connecting to new orgs, listing connected orgs, and disconnecting from orgs.",
		Tools:       org_tools.ORG_TOOLS,
	})
	lstToolCategories = append(lstToolCategories, models.ToolCategory{
		Name:        "DataTools",
		Description: "Tools for querying and manipulating Salesforce data, including creating, updating, and retrieving records.",
		Tools:       data_tools.DATA_TOOLS,
	})
	lstToolCategories = append(lstToolCategories, models.ToolCategory{
		Name:        "MetadataTools",
		Description: "Tools for retrieving and manipulating Salesforce metadata, including describing objects and fields.",
		Tools:       mdt_tools.METADATA_TOOLS,
	})
	lstToolCategories = append(lstToolCategories, models.ToolCategory{
		Name:        "ApexTools",
		Description: "Tools for interacting with Salesforce Apex and Tooling API. Includes executing Apex code, managing Apex classes and triggers, running tests, retrieving debug logs, managing trace flags, and analyzing code coverage. These tools are primarily used for development, debugging, and runtime execution of Apex logic within Salesforce.",
		Tools:       apex_tools.APEX_TOOLS,
	})
	lstToolCategories = append(lstToolCategories, models.ToolCategory{
		Name:        "LwcAuraTools",
		Description: "Tools for managing and retrieving Salesforce Frontend components. Includes fetching source code and metadata for Lightning Web Components (LWC) and Aura Components (bundles, controllers, and markup). Use these tools when the user asks about UI logic, component resources, or frontend-to-backend dependencies.",
		Tools:       lwc_aura_tools.UI_COMPONENTS_TOOLS,
	})

	srv := &http.Server{Addr: PORT}
	server := NewMCPServer(lstTools, lstToolCategories)

	if err := server.Run(srv); err != nil {
		log.Println("Failed to run Salesforce MCP Server")
		log.Fatal(err)
	}

}

func getEnv(key, defaultValue string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return defaultValue
	}
	return val
}
