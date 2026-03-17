package models

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"strings"

	"github.com/invopop/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MCPTool struct {
	Tool    *mcp.Tool
	Handler func(context.Context, *mcp.CallToolRequest, any) (*mcp.CallToolResult, any, error)
}

func RequestResponseWrapper[T any](toolHandler func(context.Context, T) (string, error)) func(context.Context, *mcp.CallToolRequest, any) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, req *mcp.CallToolRequest, rawReqBody any) (*mcp.CallToolResult, any, error) {

		var structReqBody T

		jsonReqBody, err := json.Marshal(rawReqBody)
		if err != nil {
			return nil, nil, err
		}

		if err := json.Unmarshal(jsonReqBody, &structReqBody); err != nil {
			return nil, nil, err
		}

		resBody, err := toolHandler(ctx, structReqBody)
		if err != nil {
			return nil, nil, err
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: resBody,
				},
			},
		}, nil, nil
	}
}

func NewTool[T any](toolName string, toolDescription string, toolHandler func(ctx context.Context, input T) (string, error)) MCPTool {
	var inputStructWrapper T

	properties := map[string]any{}
	required := []string{}

	if reflect.TypeOf(inputStructWrapper) != nil && reflect.TypeOf(inputStructWrapper).Kind() == reflect.Struct && reflect.TypeOf(inputStructWrapper).NumField() > 0 {
		reflector := &jsonschema.Reflector{}
		rootSchema := reflector.Reflect(inputStructWrapper)

		structSchema := rootSchema
		if rootSchema.Ref != "" && rootSchema.Definitions != nil {
			defName := strings.TrimPrefix(rootSchema.Ref, "#/$defs/")
			if resolved, ok := rootSchema.Definitions[defName]; ok {
				structSchema = resolved
			}
		}

		if structSchema.Properties != nil {
			for pair := structSchema.Properties.Oldest(); pair != nil; pair = pair.Next() {
				propJSON, _ := json.Marshal(pair.Value)
				var propMap map[string]any
				_ = json.Unmarshal(propJSON, &propMap)
				properties[pair.Key] = propMap
			}
		}

		required = structSchema.Required
	}

	inputSchema := map[string]any{
		"type":       "object",
		"properties": properties,
	}
	if len(required) > 0 {
		inputSchema["required"] = required
	}

	log.Printf("Generated input schema for tool %s: %+v\n", toolName, inputSchema)

	return MCPTool{
		Tool: &mcp.Tool{
			Name:        toolName,
			Description: toolDescription,
			InputSchema: inputSchema,
		},
		Handler: RequestResponseWrapper(toolHandler),
	}
}

/* Example of generated input schema for CreateRecordPayload:
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$ref": "#/$defs/CreateRecordPayload",
  "$defs": {
    "CreateRecordPayload": {
      "type": "object",
      "properties": {
        "sobject": { "type": "string" },
        "fields": {
          "type": "object",
          "additionalProperties": true
        }
      },
      "required": ["sobject", "fields"]
    }
  }
}
*/
