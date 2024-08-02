package spec

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"text/template"
)

const (
	KindRemoveColumns    = "remove_columns"
	KindAddColumn        = "add_column"
	KindObfuscateColumns = "obfuscate_columns"
	KindChangeTableNames = "change_table_names"
)

type TransformationSpec struct {
	Kind    string   `json:"kind"`
	Tables  []string `json:"tables"` // per-transformation table glob patterns
	Columns []string `json:"columns"`
	Name    string   `json:"name"`
	Value   string   `json:"value"`

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

var kindToRequiredFields = map[string]map[string]struct{}{
	KindRemoveColumns:    {"Columns": {}},
	KindAddColumn:        {"Name": {}, "Value": {}},
	KindObfuscateColumns: {"Columns": {}},
	KindChangeTableNames: {"NewTableNameTemplate": {}},
}

var fieldsToCheck = []string{"Columns", "Name", "Value", "FromTableName", "NewTableNameTemplate"}

func (s *Spec) Validate() error {
	var err error
	for _, t := range s.TransformationSpecs {
		requiredFields, ok := kindToRequiredFields[t.Kind]
		if !ok {
			kinds := make([]string, 0, len(kindToRequiredFields))
			for k := range kindToRequiredFields {
				kinds = append(kinds, k)
			}
			return fmt.Errorf("unknown transformation kind: %s, supported kinds are: %s", t.Kind, strings.Join(kinds, ", "))
		}
		for fieldName := range requiredFields {
			value := reflect.ValueOf(t)
			if value == reflect.Zero(value.Type()) {
				panic(fmt.Sprintf("reflect.ValueOf(%v) is zero", t)) // this would be a nil on s.TransformationSpecs
			}
			fieldValue := value.FieldByName(fieldName)
			if !fieldValue.IsValid() {
				panic(fmt.Sprintf("field %s is not valid", fieldName)) // this would be a bug on kindToRequiredFields/fieldsToCheck
			}
			if fieldValue.Kind() == reflect.String && fieldValue.String() == "" {
				err = errors.Join(err, fmt.Errorf("'%s' field must be specified for %s transformation", fieldName, t.Kind))
			}
			if fieldValue.Kind() == reflect.Slice && fieldValue.Len() == 0 {
				err = errors.Join(err, fmt.Errorf("'%s' field must be specified for %s transformation", fieldName, t.Kind))
			}
		}
		for _, fieldName := range fieldsToCheck {
			if _, ok := requiredFields[fieldName]; ok {
				continue
			}
			fieldValue := reflect.ValueOf(t).FieldByName(fieldName)
			if fieldValue.Kind() == reflect.String && fieldValue.String() != "" {
				err = errors.Join(err, fmt.Errorf("'%s' field must not be specified for %s transformation", fieldName, t.Kind))
			}
			if fieldValue.Kind() == reflect.Slice && fieldValue.Len() > 0 {
				err = errors.Join(err, fmt.Errorf("'%s' field must not be specified for %s transformation", fieldName, t.Kind))
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
