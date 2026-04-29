package tools

import (
	"context"
	"fmt"
	"log"

	repo "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/repository"
	apex_tools "github.com/Pradhyumna-Joshi/go_salesforce_mcp/internal/tools/apex_tools"
)

func GetAuraDefinition(ctx context.Context, input GetAuraDefinitionPayload) (string, error) {
	log.Printf("[MCP Tool] GetAuraDefinition: Bundle=%s, Type=%s", input.BundleName, input.AuraType)

	defType := input.AuraType
	if defType == "" {
		defType = "COMPONENT"
	}

	soql := fmt.Sprintf("SELECT Id, Source, LastModifiedDate FROM AuraDefinition WHERE AuraDefinitionBundle.DeveloperName = '%s' AND DefType = '%s'",
		input.BundleName, defType)

	return apex_tools.ApexToolingQuery(ctx, apex_tools.ApexToolingQueryPayload{Soql: soql})
}

func ListLWCWithResources(ctx context.Context, input ListLWCWithResourcesPayload) (string, error) {
	log.Printf("[MCP Tool] ListLWCWithResources: Filter=%s", input.Name)

	soql := "SELECT Id, LightningComponentBundle.DeveloperName, FilePath, Format, LastModifiedDate FROM LightningComponentResource"
	if input.Name != "" {
		soql += fmt.Sprintf(" WHERE LightningComponentBundle.DeveloperName = '%s'", input.Name)
	}

	return apex_tools.ApexToolingQuery(ctx, apex_tools.ApexToolingQueryPayload{Soql: soql})
}

func GetLWCResource(ctx context.Context, input GetLWCResourcePayload) (string, error) {
	log.Printf("[MCP Tool] GetLWCResource: Bundle=%s, File=%s", input.BundleName, input.FilePath)

	soql := fmt.Sprintf("SELECT Id, Source FROM LightningComponentResource WHERE LightningComponentBundle.DeveloperName = '%s' AND FilePath = '%s'",
		input.BundleName, input.FilePath)

	return apex_tools.ApexToolingQuery(ctx, apex_tools.ApexToolingQueryPayload{Soql: soql})
}

func UpdateLWCResource(ctx context.Context, input UpdateLWCResourcePayload) (string, error) {
	log.Printf("[MCP Tool] UpdateLWCResource: ResourceID=%s", input.ResourceId)

	endpoint := fmt.Sprintf("/services/data/%s/tooling/sobjects/LightningComponentResource/%s", apex_tools.ApiVersion, input.ResourceId)

	payload := map[string]any{
		"Source": input.Source,
	}

	body, err := repo.SalesforceRequest("PATCH", endpoint, payload)
	if err != nil {
		log.Printf("[MCP Tool Error] UpdateLWCResource failed: %v", err)
		return "", err
	}

	log.Printf("[MCP Tool Success] UpdateLWCResource: Resource %s updated", input.ResourceId)
	return string(body), nil
}
