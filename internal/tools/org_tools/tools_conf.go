package tools

import (
	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
)

var ORG_TOOLS = []models.MCPTool{
	models.NewTool("list_orgs", "List all connected Salesforce orgs. Call this first if org alias is unknown. Show user the aliases and ask which org to use.", listOrgs),
	models.NewTool("connect_to_org", "Connect to a Salesforce org via browser login.Ask user: 'production' or 'sandbox'? Ask user: what alias to give this org? (e.g. prod-org, my-sandbox) After login, call list_orgs to confirm connection.", connectToOrg),
	models.NewTool("open_org", "Open a Salesforce org in the browser.", openOrg),
	models.NewTool("disconnect_org", "Logout and disconnect from a Salesforce org. Always confirm with user before calling this.", disconnectOrg),
}
