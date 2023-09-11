package packages

type PackageJSON struct {
	SchemaVersion int              `json:"schema_version"`
	Kind          string           `json:"kind"`
	Properties    PluginProperties `json:"properties"`
	Artifacts     string           `json:"artifacts"`
	Docs          string           `json:"docs"`
	Tables        string           `json:"tables"`
	Protocols     []string         `json:"protocols"`
	Targets       []PluginTarget   `json:"targets"`
}

type PluginProperties struct {
	Source      bool `json:"source"`
	Destination bool `json:"destination"`
}

type PluginTarget struct {
	Name string `json:"name"`
	OS   string `json:"os"`
	Arch string `json:"arch"`
}
