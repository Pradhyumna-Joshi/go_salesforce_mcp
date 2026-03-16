package models

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MCPTool struct {
	Tool    *mcp.Tool
	Handler func(context.Context, *mcp.CallToolRequest, any) (*mcp.CallToolResult, any, error)
}

func WrapHandler[T any](
	f func(context.Context, T) (string, error)) func(context.Context, *mcp.CallToolRequest, any) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, input any) (*mcp.CallToolResult, any, error) {

		var data T

		b, err := json.Marshal(input)
		if err != nil {
			return nil, nil, err
		}

		if err := json.Unmarshal(b, &data); err != nil {
			return nil, nil, err
		}

		out, err := f(ctx, data)
		if err != nil {
			return nil, nil, err
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: out,
				},
			},
		}, nil, nil

	}
}
