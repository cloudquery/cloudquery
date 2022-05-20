package terraform

type ShowOutput struct {
	FormatVersion    string           `json:"format_version"`
	TerraformVersion string           `json:"terraform_version"`
	Values           ShowOutputValues `json:"values"`
}

type ShowOutputValues struct {
	RootModule RootModule `json:"root_module,omitempty"`
}

type Resource struct {
	Address       string                 `json:"address,omitempty"`
	Mode          string                 `json:"mode,omitempty"`
	Type          string                 `json:"type,omitempty"`
	Name          string                 `json:"name,omitempty"`
	ProviderName  string                 `json:"provider_name,omitempty"`
	SchemaVersion int                    `json:"schema_version"`
	Values        map[string]interface{} `json:"values,omitempty"`
}

type RootModule struct {
	Resources    []Resource    `json:"resources,omitempty"`
	ChildModules []ChildModule `json:"child_modules,omitempty"`
}

type ChildModule struct {
	Address      string        `json:"address,omitempty"`
	Resources    []Resource    `json:"resources,omitempty"`
	ChildModules []ChildModule `json:"child_modules,omitempty"`
}

// GetAllResources returns all resources sorted by provider in the given show output
func (o *ShowOutput) GetAllResourcesByProvider() map[string][]Resource {
	result := make(map[string][]Resource)
	for _, r := range o.Values.RootModule.Resources {
		result[r.ProviderName] = append(result[r.ProviderName], r)
	}
	for _, r := range o.Values.RootModule.ChildModules {
		subChild := getAllResourcesByProvider(&r)
		result = mergeResourcesMap(result, subChild)
	}
	return result
}

func mergeResourcesMap(m1, m2 map[string][]Resource) map[string][]Resource {
	for k, v := range m2 {
		m1[k] = append(m1[k], v...)
	}
	return m1
}

func getAllResourcesByProvider(childModule *ChildModule) map[string][]Resource {
	result := make(map[string][]Resource)
	for _, r := range childModule.Resources {
		result[r.ProviderName] = append(result[r.ProviderName], r)
	}
	for _, r := range childModule.ChildModules {
		subChild := getAllResourcesByProvider(&r)
		result = mergeResourcesMap(result, subChild)
	}
	return result
}
