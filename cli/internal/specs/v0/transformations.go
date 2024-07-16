package specs

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/cloudquery/cloudquery/cli/transformations"
	"github.com/invopop/jsonschema"
)

type TransformationSpec struct {
	Kind TransformationKind `json:"kind,omitempty"`

	// For TransformationKindAddField, TransformationKindRemoveField
	FieldName string `json:"field_name,omitempty"`

	// For TransformationKindAddField, TransformationKindDataOnlyCustom, TransformationKindCustom
	SQL string `json:"sql,omitempty"`

	// For TransformationKindCustom
	Schema []ArrowSchemaSpec `json:"schema,omitempty"`

	// For TransformationKindAddField
	DataType string            `json:"data_type,omitempty"`
	Nullable bool              `json:"nullable,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (m *TransformationSpec) ToTransformation() (transformations.Transformation, error) {
	switch m.Kind {
	case TransformationKindCustom:
		schema, err := schemaFromFields(m.Schema)
		if err != nil {
			return transformations.Transformation{}, err
		}
		return transformations.NewCustomTransformation(m.SQL, schema)
	case TransformationKindDataOnlyCustom:
		return transformations.NewCustomDataOnlyTransformation(m.SQL)
	case TransformationKindAddField:
		arrowSchemaSpec := ArrowSchemaSpec{
			FieldName: m.FieldName,
			DataType:  m.DataType,
			Nullable:  m.Nullable,
			Metadata:  m.Metadata,
		}
		field, err := arrowSchemaSpec.ToField()
		if err != nil {
			return transformations.Transformation{}, err
		}
		return transformations.NewAddFieldTransformation(m.FieldName, m.SQL, field)
	case TransformationKindRemoveField:
		return transformations.NewRemoveFieldTransformation(m.FieldName)
	case TransformationKindUpdateField:
		return transformations.NewUpdateFieldTransformation(m.FieldName, m.SQL)
	default:
		return transformations.Transformation{}, fmt.Errorf("unknown transformation kind: %s", m.Kind)
	}
}

func (m *TransformationSpec) Validate() error {
	if m == nil {
		return nil
	}
	var (
		kindToFields = map[TransformationKind]map[string]bool{
			TransformationKindAddField:       {"field_name": true, "sql": true, "data_type": true, "nullable": true, "metadata": true},
			TransformationKindRemoveField:    {"field_name": true},
			TransformationKindUpdateField:    {"field_name": true, "sql": true},
			TransformationKindDataOnlyCustom: {"sql": true},
			TransformationKindCustom:         {"sql": true, "schema": true},
		}
	)
	if m.Kind == TransformationKindUnknown {
		return fmt.Errorf("kind is required")
	}
	if m.FieldName == "" && kindToFields[m.Kind]["field_name"] {
		return fmt.Errorf("field_name is required for transformation kind %s", m.Kind)
	}
	if m.FieldName != "" && !kindToFields[m.Kind]["field_name"] {
		return fmt.Errorf("field_name is not allowed for transformation kind %s", m.Kind)
	}
	if m.SQL == "" && kindToFields[m.Kind]["sql"] {
		return fmt.Errorf("sql is required for transformation kind %s", m.Kind)
	}
	if m.SQL != "" && !kindToFields[m.Kind]["sql"] {
		return fmt.Errorf("sql is not allowed for transformation kind %s", m.Kind)
	}
	if len(m.Schema) == 0 && kindToFields[m.Kind]["schema"] {
		return fmt.Errorf("schema is required for transformation kind %s", m.Kind)
	}
	if len(m.Schema) > 0 && !kindToFields[m.Kind]["schema"] {
		return fmt.Errorf("schema is not allowed for transformation kind %s", m.Kind)
	}
	if m.DataType == "" && kindToFields[m.Kind]["data_type"] {
		return fmt.Errorf("data_type is required for transformation kind %s", m.Kind)
	}
	if m.DataType != "" && !kindToFields[m.Kind]["data_type"] {
		return fmt.Errorf("data_type is not allowed for transformation kind %s", m.Kind)
	}
	if m.Nullable && !kindToFields[m.Kind]["nullable"] {
		return fmt.Errorf("nullable is not allowed for transformation kind %s", m.Kind)
	}
	if m.Metadata != nil && !kindToFields[m.Kind]["metadata"] {
		return fmt.Errorf("metadata is not allowed for transformation kind %s", m.Kind)
	}
	if len(m.Schema) > 0 {
		_, err := schemaFromFields(m.Schema)
		if err != nil {
			return err
		}
	}

	return nil
}

type ArrowSchemaSpec struct {
	FieldName string            `json:"field_name,omitempty"`
	DataType  string            `json:"data_type,omitempty"`
	Nullable  bool              `json:"nullable,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

func (s ArrowSchemaSpec) Validate() error {
	if s.FieldName == "" {
		return fmt.Errorf("schema's name is required")
	}
	if s.DataType == "" {
		return fmt.Errorf("schema's datatype is required")
	}
	if _, err := arrowTypeFromString(s.DataType); err != nil {
		return err
	}
	return nil
}

func (s ArrowSchemaSpec) ToField() (arrow.Field, error) {
	if err := s.Validate(); err != nil {
		return arrow.Field{}, err
	}

	dt, err := arrowTypeFromString(s.DataType)
	if err != nil {
		return arrow.Field{}, err
	}
	keys := make([]string, 0, len(s.Metadata))
	values := make([]string, 0, len(s.Metadata))
	for k, v := range s.Metadata {
		keys = append(keys, k)
		values = append(values, v)
	}
	return arrow.Field{
		Name:     s.FieldName,
		Type:     dt,
		Nullable: s.Nullable,
		Metadata: arrow.NewMetadata(keys, values),
	}, nil
}

func schemaFromFields(fields []ArrowSchemaSpec) (*arrow.Schema, error) {
	arrowFields := make([]arrow.Field, len(fields))
	for i, f := range fields {
		arrowField, err := f.ToField()
		if err != nil {
			return nil, err
		}
		arrowFields[i] = arrowField
	}
	return arrow.NewSchema(arrowFields, nil), nil
}

func arrowTypeFromString(t string) (arrow.DataType, error) {
	switch t {
	case "null":
		return arrow.Null, nil
	case "bool":
		return arrow.FixedWidthTypes.Boolean, nil
	case "int8":
		return arrow.PrimitiveTypes.Int8, nil
	case "int16":
		return arrow.PrimitiveTypes.Int16, nil
	case "int32":
		return arrow.PrimitiveTypes.Int32, nil
	case "int64":
		return arrow.PrimitiveTypes.Int64, nil
	case "uint8":
		return arrow.PrimitiveTypes.Uint8, nil
	case "uint16":
		return arrow.PrimitiveTypes.Uint16, nil
	case "uint32":
		return arrow.PrimitiveTypes.Uint32, nil
	case "uint64":
		return arrow.PrimitiveTypes.Uint64, nil
	case "float32":
		return arrow.PrimitiveTypes.Float32, nil
	case "float64":
		return arrow.PrimitiveTypes.Float64, nil
	case "binary":
		return arrow.BinaryTypes.Binary, nil
	case "string", "utf8":
		return arrow.BinaryTypes.String, nil
	case "uuid":
		return &arrow.FixedSizeBinaryType{ByteWidth: 16}, nil
	case "timestamp":
		return &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"}, nil
	default:
		return nil, fmt.Errorf("schema's data type is invalid: %s", t)
	}
}

type TransformationKind int

const (
	TransformationKindUnknown TransformationKind = iota
	TransformationKindCustom
	TransformationKindDataOnlyCustom
	TransformationKindAddField
	TransformationKindRemoveField
	TransformationKindUpdateField
)

var (
	AllTransformationKinds = [...]string{
		TransformationKindUnknown:        "unknown",
		TransformationKindCustom:         "custom",
		TransformationKindDataOnlyCustom: "data_only_custom",
		TransformationKindAddField:       "add_field",
		TransformationKindRemoveField:    "remove_field",
		TransformationKindUpdateField:    "update_field",
	}
)

func (m TransformationKind) String() string {
	return AllTransformationKinds[m]
}

func (m TransformationKind) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(m.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (m *TransformationKind) UnmarshalJSON(data []byte) (err error) {
	var transformationKind string
	if err := json.Unmarshal(data, &transformationKind); err != nil {
		return err
	}
	if *m, err = TransformationKindFromString(transformationKind); err != nil {
		return err
	}
	return nil
}

func (TransformationKind) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Type = "string"
	sc.Enum = make([]any, len(AllTransformationKinds))
	for i, k := range AllTransformationKinds {
		sc.Enum[i] = k
	}
}

func TransformationKindFromString(s string) (TransformationKind, error) {
	for m, str := range AllTransformationKinds {
		if s == str {
			return TransformationKind(m), nil
		}
	}
	return TransformationKindUnknown, fmt.Errorf("invalid migrate mode: %s", s)
}
