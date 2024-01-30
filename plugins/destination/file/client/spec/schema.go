package spec

import (
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/invopop/jsonschema"
)

func (Spec) JSONSchemaAlias() any {
	return SpecSchema{}
}

type SpecSchema struct {
	FileSpec  filetypes.FileSpec
	Directory string `json:"directory,omitempty" jsonschema:"oneof_required=directory"`
	Path      string `json:"path,omitempty" jsonschema:"oneof_required=path"`
	NoRotate  bool   `json:"no_rotate,omitempty"`

	BatchSize      *int64               `json:"batch_size"`
	BatchSizeBytes *int64               `json:"batch_size_bytes"`
	BatchTimeout   *configtype.Duration `json:"batch_timeout"`
}

// JSONSchemaExtend is required to actually represent the embedding of filetypes.FileSpec
func (SpecSchema) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Ref = sc.Properties.Value("FileSpec").Ref // this should work per ref allowed with additional stuff
	sc.Properties.Delete("FileSpec")
}
