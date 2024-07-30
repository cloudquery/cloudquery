package spec

import (
	"errors"
	"fmt"
)

const (
	KindRemoveColumns    = "remove_columns"
	KindAddColumn        = "add_column"
	KindObfuscateColumns = "obfuscate_columns"
)

type TransformationSpec struct {
	Kind    string   `json:"kind"`
	Tables  []string `json:"tables"` // per-transformation table glob patterns
	Columns []string `json:"columns"`
	Name    string   `json:"name"`
	Value   string   `json:"value"`
}

type Spec struct {
	TransformationSpecs []TransformationSpec `json:"transformations"`
}

func (s *Spec) SetDefaults() {
	for i := range s.TransformationSpecs {
		if len(s.TransformationSpecs[i].Tables) == 0 {
			s.TransformationSpecs[i].Tables = append(s.TransformationSpecs[i].Tables, "*")
		}
	}
}

func (s *Spec) Validate() error {
	var err error
	for _, t := range s.TransformationSpecs {
		switch t.Kind {
		case KindRemoveColumns:
			if len(t.Columns) == 0 {
				err = errors.Join(err, fmt.Errorf("'columns' field must be specified for remove_columns transformation"))
			}
			if t.Name != "" {
				err = errors.Join(err, fmt.Errorf("'name' field must not be specified for remove_columns transformation"))
			}
			if t.Value != "" {
				err = errors.Join(err, fmt.Errorf("'value' field must not be specified for remove_columns transformation"))
			}
		case KindAddColumn:
			if t.Name == "" {
				err = errors.Join(err, fmt.Errorf("'name' field must be specified for add_column transformation"))
			}
			if t.Value == "" {
				err = errors.Join(err, fmt.Errorf("'value' field must be specified for add_column transformation"))
			}
			if len(t.Columns) > 0 {
				err = errors.Join(err, fmt.Errorf("'columns' field must not be specified for add_column transformation"))
			}
		case KindObfuscateColumns:
			if len(t.Columns) == 0 {
				err = errors.Join(err, fmt.Errorf("'columns' field must be specified for obfuscate_columns transformation"))
			}
			if t.Name != "" {
				err = errors.Join(err, fmt.Errorf("'name' field must not be specified for obfuscate_columns transformation"))
			}
			if t.Value != "" {
				err = errors.Join(err, fmt.Errorf("'value' field must not be specified for obfuscate_columns transformation"))
			}
		default:
			err = errors.Join(err, fmt.Errorf("invalid transformation kind: %s; must be one of: remove_columns, add_column, obfuscate_columns", t.Kind))
		}
	}
	return err
}
