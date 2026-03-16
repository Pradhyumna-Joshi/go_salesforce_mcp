package main

import (
	"log"
	"net/http"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
	org_tools "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/tools/org_tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MCPServer struct {
	tools []models.MCPTool
}

func NewMCPServer(tools []models.MCPTool) *MCPServer {
	return &MCPServer{tools}
}

func (s *MCPServer) Run() error {

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "SalesforceMCP",
		Version: "1.0.0",
	}, nil)

	for _, t := range s.tools {
		mcp.AddTool(server, t.Tool, t.Handler)
	}

	//return server.Run(context.Background(), &mcp.StdioTransport{})

	//FOR REMOTE MCP SERVER
	handler := mcp.NewStreamableHTTPHandler(func(r *http.Request) *mcp.Server {
		return server
	}, nil)

	http.Handle("/mcp", handler)

	http.HandleFunc("/callback", org_tools.HandleCallBack)

	log.Println("MCP server running on : 8080")

	return http.ListenAndServe(":8080", nil)
}
