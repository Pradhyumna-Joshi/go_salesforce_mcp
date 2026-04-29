package tools

type GetAuraDefinitionPayload struct {
	BundleName string `json:"bundle_name" jsonschema:"description=Name of the Aura Bundle. Example: MyAuraComponent"`
	AuraType   string `json:"aura_type,omitempty" jsonschema:"description=Specific file type to retrieve: COMPONENT, CONTROLLER, HELPER, STYLE, RENDERER, DOCUMENTATION, DESIGN, SVG. Defaults to COMPONENT."`
}

type CreateAuraBundlePayload struct {
	Name        string `json:"name" jsonschema:"description=Name of the new Aura bundle."`
	Description string `json:"description,omitempty" jsonschema:"description=Optional description."`
}

type ListLWCWithResourcesPayload struct {
	Name string `json:"name,omitempty" jsonschema:"description=Optional name to filter a specific LWC bundle."`
}

type GetLWCResourcePayload struct {
	BundleName string `json:"bundle_name" jsonschema:"description=Name of the LWC bundle."`
	FilePath   string `json:"file_path" jsonschema:"description=Path to the specific file. Example: lwc/myComponent/myComponent.js"`
}

type UpdateLWCResourcePayload struct {
	ResourceId string `json:"resource_id" jsonschema:"description=The Tooling ID of the specific LightningComponentResource (the file)."`
	Source     string `json:"source" jsonschema:"description=The full updated source code for the file."`
}
