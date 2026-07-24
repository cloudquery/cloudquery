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
	KindObfuscateColumnsExcept    = "obfuscate_columns_except"
	KindUppercase                 = "uppercase"
	KindLowercase                 = "lowercase"
	KindDropRows                  = "drop_rows"
)

const (
	UnmatchedDrop     = "drop"
	UnmatchedCollapse = "collapse"
	UnmatchedRedact   = "redact"
)

type TransformationSpec struct {
	Kind    string   `json:"kind"`
	Tables  []string `json:"tables"` // per-transformation table glob patterns
	Columns []string `json:"columns"`
	Name    string   `json:"name"`
	Value   *string  `json:"value"`

	// IncludeSHA controls whether obfuscation transformations append the SHA-256 hash
	// of the redacted value to the redaction marker. Defaults to true. Only meaningful
	// for the obfuscate_* transformation kinds.
	IncludeSHA *bool `json:"include_sha"`

	// Unmatched controls how obfuscate_columns_except handles fields that are not on the
	// allowlist: "drop" removes them (default), "collapse" replaces each non-allowlisted
	// object/array with a single redaction marker, and "redact" replaces every individual
	// leaf value with a marker. Only meaningful for the obfuscate_columns_except kind.
	Unmatched string `json:"unmatched"`

	// For change_table_names transformation
	NewTableNameTemplate string `json:"new_table_name_template"`
}

func (t TransformationSpec) ShouldIncludeSHA() bool {
	return t.IncludeSHA == nil || *t.IncludeSHA
}

func (t TransformationSpec) UnmatchedMode() string {
	if t.Unmatched == "" {
		return UnmatchedDrop
	}
	return t.Unmatched
}

func isObfuscateKind(kind string) bool {
	switch kind {
	case KindObfuscateColumns, KindObfuscateSensitiveColumns, KindObfuscateColumnsExcept:
		return true
	default:
		return false
	}
}

type Spec struct {
	TransformationSpecs []TransformationSpec `json:"transformations"`
}

func (s *Spec) SetDefaults() {
	for i := range s.TransformationSpecs {
		if len(s.TransformationSpecs[i].Tables) == 0 {
			s.TransformationSpecs[i].Tables = append(s.TransformationSpecs[i].Tables, "*")
		}
		if isObfuscateKind(s.TransformationSpecs[i].Kind) && s.TransformationSpecs[i].IncludeSHA == nil {
			t := true
			s.TransformationSpecs[i].IncludeSHA = &t
		}
	}
}

func (s *Spec) Validate() error {
	var err error
	for _, t := range s.TransformationSpecs {
		switch t.Kind {
		case KindRemoveColumns, KindAddPrimaryKeys, KindObfuscateColumns, KindObfuscateColumnsExcept:
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

		if t.IncludeSHA != nil && !isObfuscateKind(t.Kind) {
			err = errors.Join(err, fmt.Errorf("include_sha field must not be specified for %s transformation", t.Kind))
		}

		if t.Unmatched != "" {
			if t.Kind != KindObfuscateColumnsExcept {
				err = errors.Join(err, fmt.Errorf("unmatched field must not be specified for %s transformation", t.Kind))
			} else if t.Unmatched != UnmatchedDrop && t.Unmatched != UnmatchedCollapse && t.Unmatched != UnmatchedRedact {
				err = errors.Join(err, fmt.Errorf("unmatched must be one of %q, %q, %q for %s transformation", UnmatchedDrop, UnmatchedCollapse, UnmatchedRedact, t.Kind))
			}
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
