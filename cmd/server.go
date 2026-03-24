package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/models"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MCPServer struct {
	tools        []models.MCPTool
	toolCategory []models.ToolCategory
}

type MCPServerMetadata struct {
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Toolcategory []models.ToolCategory `json:"tool_categories"`
}

func NewMCPServer(tools []models.MCPTool, toolCategory []models.ToolCategory) *MCPServer {
	return &MCPServer{
		tools:        tools,
		toolCategory: toolCategory,
	}
}

func (s *MCPServer) Run() error {

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "SalesforceMCP",
		Version: "1.0.0",
	}, nil)

	mcpServerMetadata := MCPServerMetadata{
		Name:         "SalesforceMCP",
		Description:  "A server that exposes Salesforce related tools via the Model Context Protocol (MCP). This server allows clients to interact with Salesforce data, metadata, and org management functionalities through a standardized interface.",
		Toolcategory: s.toolCategory,
	}

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_all_tools",
		Description: "List all available tools and their descriptions",
		InputSchema: nil,
	}, func(ctx context.Context, req *mcp.CallToolRequest, rawReqBody any) (*mcp.CallToolResult, any, error) {
		jsonReqBody, err := json.Marshal(mcpServerMetadata)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: string(jsonReqBody),
				},
			},
		}, nil, nil
	})

	for _, t := range s.tools {
		mcp.AddTool(server, t.Tool, t.Handler)
	}

	//return server.Run(context.Background(), &mcp.StdioTransport{})

	//FOR REMOTE MCP SERVER
	handler := mcp.NewStreamableHTTPHandler(func(r *http.Request) *mcp.Server {
		return server
	}, nil)

	http.Handle("/mcp", handler)

	log.Println("MCP server running on : 9000")

	return http.ListenAndServe(":9000", nil)

}

type NoInput struct{}
