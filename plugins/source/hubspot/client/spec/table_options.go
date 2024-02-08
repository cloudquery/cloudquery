package spec

type TableOptions map[string]*TableOptionsSpec

// Table options spec.
type TableOptionsSpec struct {
	// List of properties to sync. If empty, everything is synced.
	Properties []string `yaml:"properties,omitempty" json:"properties,omitempty" jsonschema:"minLength=1"`
	// List of associations to sync. If empty, everything is synced.
	Associations []string `yaml:"associations,omitempty" json:"associations,omitempty" jsonschema:"minLength=1"`
}

func (ts TableOptions) ForTable(name string) *TableOptionsSpec {
	return ts[name]
}

func (to *TableOptionsSpec) GetProperties() []string {
	if to == nil {
		return nil
	}
	return to.Properties
}

func (to *TableOptionsSpec) GetAssociations() []string {
	if to == nil {
		return nil
	}
	return to.Associations
}
