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
	KindObfuscateColumnsExcept    = "obfuscate_columns_except"
	KindRemoveColumnsExcept       = "remove_columns_except"
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

	// Optional redaction controls, applying only to obfuscate_columns,
	// obfuscate_columns_except and obfuscate_sensitive_columns transformations.
	// When unset, obfuscation output is byte-identical to the historical default.
	Redaction *Redaction `json:"redaction"`
}

// PlaintextRedaction configures the redacted value used for STRING and BINARY columns,
// as well as JSON-path (col.a.b) values.
type PlaintextRedaction struct {
	Message     string `json:"message"`
	IncludeHash bool   `json:"include_hash"`
}

// JSONRedaction configures the redacted value used when a whole JSON column is obfuscated.
type JSONRedaction struct {
	Key         string `json:"key"`
	Message     string `json:"message"`
	IncludeHash bool   `json:"include_hash"`
}

// Redaction bundles the plaintext and whole-JSON redaction options. Both must be set together.
type Redaction struct {
	Plaintext *PlaintextRedaction `json:"plaintext"`
	JSON      *JSONRedaction      `json:"json"`
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
		case KindRemoveColumns, KindAddPrimaryKeys, KindObfuscateColumns, KindRemoveColumnsExcept:
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
		case KindObfuscateColumnsExcept:
			// Unlike obfuscate_columns, the keep-list may be empty (empty = obfuscate every obfuscatable column).
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

		if t.Redaction != nil {
			err = errors.Join(err, validateRedaction(t))
		}
	}

	return err
}

func validateRedaction(t TransformationSpec) error {
	var err error
	switch t.Kind {
	case KindObfuscateColumns, KindObfuscateColumnsExcept, KindObfuscateSensitiveColumns:
	default:
		err = errors.Join(err, fmt.Errorf("redaction field is only valid for obfuscate_columns, obfuscate_columns_except and obfuscate_sensitive_columns transformations, not %s", t.Kind))
	}
	if t.Redaction.Plaintext == nil || t.Redaction.JSON == nil {
		return errors.Join(err, errors.New("both 'plaintext' and 'json' must be specified when 'redaction' is set"))
	}
	if t.Redaction.Plaintext.Message == "" {
		err = errors.Join(err, errors.New("redaction.plaintext.message must not be empty"))
	}
	if t.Redaction.JSON.Message == "" {
		err = errors.Join(err, errors.New("redaction.json.message must not be empty"))
	}
	if t.Redaction.JSON.Key == "" {
		err = errors.Join(err, errors.New("redaction.json.key must not be empty"))
	}
	return err
}
