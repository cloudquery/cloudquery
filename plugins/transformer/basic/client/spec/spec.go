package spec

import (
	"errors"
	"fmt"
	"text/template"
)

const (
	KindRemoveColumns             = "remove_columns"
	KindAddColumn                 = "add_column"
	KindObfuscateColumns          = "obfuscate_columns"
	KindChangeTableNames          = "change_table_names"
	KindAddTimestampColumn        = "add_current_timestamp_column"
	KindRenameColumn              = "rename_column"
	KindAddPrimaryKeys            = "add_primary_keys"
	KindObfuscateSensitiveColumns = "obfuscate_sensitive_columns"
	KindUppercase                 = "uppercase"
	KindLowercase                 = "lowercase"
	KindDropRows                  = "drop_rows"
)

type TransformationSpec struct {
	Kind    string   `json:"kind"`
	Tables  []string `json:"tables"` // per-transformation table glob patterns
	Columns []string `json:"columns"`
	Name    string   `json:"name"`
	Value   *string  `json:"value"`

	// For change_table_names transformation
	NewTableNameTemplate string `json:"new_table_name_template"`
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
		case KindRemoveColumns, KindAddPrimaryKeys:
			if len(t.Columns) == 0 {
				err = errors.Join(err, fmt.Errorf("'%s' field must be specified for %s transformation", "columns", t.Kind))
			}
			if t.Name != "" || (t.Value != nil && *t.Value != "") || t.NewTableNameTemplate != "" {
				err = errors.Join(err, fmt.Errorf("name/value/new_table_name_template fields must not be specified for %s transformation", t.Kind))
			}
		case KindAddColumn:
			if t.Name == "" || t.Value == nil || *t.Value == "" {
				err = errors.Join(err, fmt.Errorf("'%s' and '%s' fields must be specified for %s transformation", "name", "value", t.Kind))
			}
			if t.NewTableNameTemplate != "" {
				err = errors.Join(err, fmt.Errorf("new_table_name_template field must not be specified for %s transformation", t.Kind))
			}
		case KindAddTimestampColumn:
			if t.Name == "" {
				err = errors.Join(err, fmt.Errorf("'%s' field must be specified for %s transformation", "name", t.Kind))
			}
			if (t.Value != nil && *t.Value != "") || len(t.Columns) > 0 || t.NewTableNameTemplate != "" {
				err = errors.Join(err, fmt.Errorf("value/columns/new_table_name_template fields must not be specified for %s transformation", t.Kind))
			}
			if t.NewTableNameTemplate != "" {
				err = errors.Join(err, fmt.Errorf("new_table_name_template field must not be specified for %s transformation", t.Kind))
			}
		case KindObfuscateColumns:
			if len(t.Columns) == 0 {
				err = errors.Join(err, fmt.Errorf("'%s' field must be specified for %s transformation", "columns", t.Kind))
			}
			if t.Name != "" || (t.Value != nil && *t.Value != "") || t.NewTableNameTemplate != "" {
				err = errors.Join(err, fmt.Errorf("name/value/new_table_name_template fields must not be specified for %s transformation", t.Kind))
			}
		case KindObfuscateSensitiveColumns:
			if len(t.Columns) != 0 {
				err = errors.Join(err, fmt.Errorf("'%s' field must not be specified for %s transformation", "columns", t.Kind))
			}
			if t.Name != "" || (t.Value != nil && *t.Value != "") || t.NewTableNameTemplate != "" {
				err = errors.Join(err, fmt.Errorf("name/value/new_table_name_template fields must not be specified for %s transformation", t.Kind))
			}
		case KindChangeTableNames:
			if t.NewTableNameTemplate == "" {
				err = errors.Join(err, fmt.Errorf("'%s' field must be specified for %s transformation", "new_table_name_template", t.Kind))
			}
			if t.Name != "" || (t.Value != nil && *t.Value != "") || len(t.Columns) > 0 {
				err = errors.Join(err, fmt.Errorf("name/value/columns fields must not be specified for %s transformation", t.Kind))
			}
		case KindRenameColumn:
			if t.Name == "" || t.Value == nil || *t.Value == "" {
				err = errors.Join(err, fmt.Errorf("'%s' and '%s' fields must be specified for %s transformation", "name", "value", t.Kind))
			}
			if t.NewTableNameTemplate != "" {
				err = errors.Join(err, fmt.Errorf("new_table_name_template field must not be specified for %s transformation", t.Kind))
			}
			if len(t.Columns) > 0 {
				err = errors.Join(err, fmt.Errorf("columns field must not be specified for %s transformation", t.Kind))
			}
		case KindLowercase, KindUppercase:
			if t.Value != nil && *t.Value != "" {
				err = errors.Join(err, fmt.Errorf("value field must be empty for %s transformation", t.Kind))
			}
			if len(t.Columns) == 0 {
				err = errors.Join(err, fmt.Errorf("'%s' field must be specified for %s transformation", "columns", t.Kind))
			}
			if t.Name != "" || t.NewTableNameTemplate != "" {
				err = errors.Join(err, fmt.Errorf("name/new_table_name_template fields must not be specified for %s transformation", t.Kind))
			}
		case KindDropRows:
			if len(t.Columns) == 0 {
				err = errors.Join(err, fmt.Errorf("'columns' must be specified for %s transformation", t.Kind))
			}

		default:
			err = errors.Join(err, fmt.Errorf("unknown transformation kind: %s", t.Kind))
		}

		// Non-trivial validations
		if t.Kind == KindChangeTableNames {
			if _, tplErr := template.New("table_name").Parse(t.NewTableNameTemplate); err != nil {
				err = errors.Join(err, fmt.Errorf("error parsing new_table_name_template: %v", tplErr))
			}
		}
	}

	return err
}
