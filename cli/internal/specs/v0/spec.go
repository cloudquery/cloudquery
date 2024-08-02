package specs

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

// Warnings is a map of field name to string, used mainly for deprecation notices.
type Warnings map[string]string

type Spec struct {
	// CloudQuery plugin kind
	Kind Kind `json:"kind" jsonschema:"required"`

	// CloudQuery plugin (top-level) spec
	Spec any `json:"spec" jsonschema:"required"`
}

func (s *Spec) UnmarshalJSON(data []byte) error {
	var t struct {
		Kind Kind `json:"kind"`
		Spec any  `json:"spec"`
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	dec.UseNumber()
	if err := dec.Decode(&t); err != nil {
		return err
	}
	s.Kind = t.Kind
	switch s.Kind {
	case KindSource:
		s.Spec = new(Source)
	case KindDestination:
		s.Spec = new(Destination)
	case KindTransformer:
		s.Spec = new(Transformer)
	default:
		return fmt.Errorf("unknown kind %s", s.Kind)
	}
	b, err := json.Marshal(t.Spec)
	if err != nil {
		return err
	}
	dec = json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	dec.DisallowUnknownFields()
	return dec.Decode(s.Spec)
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	// delete & obtain the values
	source, _ := sc.Properties.Delete("Source")
	destination, _ := sc.Properties.Delete("Destination")
	transformer, _ := sc.Properties.Delete("Transformer")

	sc.AllOf = []*jsonschema.Schema{
		{
			// `kind: source` implies source spec
			If: &jsonschema.Schema{
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()
					kind := *sc.Properties.Value("kind")
					kind.Const = "source"
					kind.Enum = nil
					properties.Set("kind", &kind)
					return properties
				}(),
			},
			Then: &jsonschema.Schema{
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()
					properties.Set("spec", source)
					return properties
				}(),
			},
		},
		{
			// `kind: destination` implies destination spec
			If: &jsonschema.Schema{
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()
					kind := *sc.Properties.Value("kind")
					kind.Const = "destination"
					kind.Enum = nil
					properties.Set("kind", &kind)
					return properties
				}(),
			},
			Then: &jsonschema.Schema{
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()
					properties.Set("spec", destination)
					return properties
				}(),
			},
		},
		{
			// `kind: transformer` implies transformer spec
			If: &jsonschema.Schema{
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()
					kind := *sc.Properties.Value("kind")
					kind.Const = "transformer"
					kind.Enum = nil
					properties.Set("kind", &kind)
					return properties
				}(),
			},
			Then: &jsonschema.Schema{
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()
					properties.Set("spec", transformer)
					return properties
				}(),
			},
		},
	}
}

func SpecUnmarshalYamlStrict(b []byte, spec *Spec) error {
	jb, err := yaml.YAMLToJSON(b)
	if err != nil {
		return fmt.Errorf("failed to convert yaml to json: %w", err)
	}
	dec := json.NewDecoder(bytes.NewReader(jb))
	dec.DisallowUnknownFields()
	dec.UseNumber()
	if err := dec.Decode(spec); err != nil {
		return fmt.Errorf("failed to decode spec: %w", err)
	}
	return nil
}
