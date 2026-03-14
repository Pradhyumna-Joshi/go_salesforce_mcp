package main

import (
	"log"
	"net/http"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MCPServer struct {
	tools []models.MCPTools
}

func NewMCPServer(tools []models.MCPTools) *MCPServer {
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

	handler := mcp.NewSSEHandler(func(r *http.Request) *mcp.Server {
		return server
	}, nil)

	http.Handle("/mcp", handler)

	log.Println("MCP server running on : 8080")

	return http.ListenAndServe(":8080", nil)
}
