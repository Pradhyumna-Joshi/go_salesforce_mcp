package tools

import (
	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var ORG_TOOLS = []models.MCPTool{
	{
		Tool: &mcp.Tool{
			Name:        "Login_to_Salesforce_org",
			Description: "This tool is used to login to the salesforce org",
		},
		Handler: models.WrapHandler(LoginToOrgHandler),
	},
	{
		Tool: &mcp.Tool{
			Name:        "list_all_orgs",
			Description: "This tool is used to list all salesforce org",
		},
		Handler: models.WrapHandler(ListAllOrgs),
	},
	{
		Tool: &mcp.Tool{
			Name:        "open_salesforce_org",
			Description: "This tool is used open salesforce org in default browser",
		},
		Handler: models.WrapHandler(OpenSalesforceOrg),
	},
	{
		Tool: &mcp.Tool{
			Name:        "disconnect_salesforce_org",
			Description: "This tool is used disconnect salesforce org",
		},
		Handler: models.WrapHandler(DisconnectSalesforceOrg),
	},
}
