package spec

import (
	_ "embed"

	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func (s Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	s.FileSpec.JSONSchemaExtend(sc) // need to call manually

	strValueIsSet := func(property string) *jsonschema.Schema {
		return &jsonschema.Schema{
			Title: "`" + property + "` value is set",
			Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
				p := *sc.Properties.Value(property)
				p.Default = nil
				p.Description = ""
				p.MinLength = &([]uint64{1}[0])
				properties := orderedmap.New[string, *jsonschema.Schema]()
				properties.Set(property, &p)
				return properties
			}(),
			Required: []string{property},
		}
	}
	usernamePresent := strValueIsSet("sasl_username")
	passwordPresent := strValueIsSet("sasl_password")

	sc.AllOf = append(sc.AllOf,
		&jsonschema.Schema{
			Title: "Require `sasl_password` when `sasl_username` is set",
			If:    usernamePresent,
			Then:  passwordPresent,
		},
		&jsonschema.Schema{
			Title: "Require `sasl_username` when `sasl_password` is set",
			If:    passwordPresent,
			Then:  usernamePresent,
		},
	)
}

//go:embed schema.json
var JSONSchema string
