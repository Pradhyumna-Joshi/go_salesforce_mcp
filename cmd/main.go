package main

import (
	"log"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
	data_tools "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/tools/data_tools"
	mdt_tools "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/tools/metadata_tools"
	org_tools "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/tools/org_tools"
)

func main() {

	lstTools := make([]models.MCPTools, 0)
	lstTools = append(lstTools, org_tools.ORG_TOOLS...)
	lstTools = append(lstTools, data_tools.DATA_TOOLS...)
	lstTools = append(lstTools, mdt_tools.METADATA_TOOLS...)

	server := NewMCPServer(lstTools)

	if err := server.Run(); err != nil {
		log.Println("Failed to run server")
	}

}
