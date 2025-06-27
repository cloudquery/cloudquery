package spec

import (
	"testing"
)

func TestSetDefaults(t *testing.T) {
	tests := []struct {
		name     string
		input    Spec
		expected Spec
	}{
		{
			name: "Adds * as default pattern",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindRemoveColumns, Columns: []string{"col1"}},
				},
			},
			expected: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindRemoveColumns, Tables: []string{"*"}, Columns: []string{"col1"}},
				},
			},
		},
		{
			name: "Leaves as is if pattern is already set",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindRemoveColumns, Tables: []string{"table1"}, Columns: []string{"col1"}},
				},
			},
			expected: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindRemoveColumns, Tables: []string{"table1"}, Columns: []string{"col1"}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.input.SetDefaults()
			if len(tt.input.TransformationSpecs) != len(tt.expected.TransformationSpecs) {
				t.Fatalf("Expected %d transformation specs, got %d", len(tt.expected.TransformationSpecs), len(tt.input.TransformationSpecs))
			}
			for i, spec := range tt.input.TransformationSpecs {
				if len(spec.Tables) != len(tt.expected.TransformationSpecs[i].Tables) {
					t.Errorf("Expected tables %v, got %v", tt.expected.TransformationSpecs[i].Tables, spec.Tables)
				}
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   Spec
		wantErr bool
	}{
		{
			name: "ValidRemoveColumns",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindRemoveColumns, Columns: []string{"col1"}},
				},
			},
			wantErr: false,
		},
		{
			name: "InvalidRemoveColumnsNoColumns",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindRemoveColumns},
				},
			},
			wantErr: true,
		},
		{
			name: "ValidAddColumn",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindAddColumn, Name: "new_col", Value: &[]string{"default"}[0]},
				},
			},
			wantErr: false,
		},
		{
			name: "InvalidAddColumnNoName",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindAddColumn, Value: &[]string{"default"}[0]},
				},
			},
			wantErr: true,
		},
		{
			name: "ValidAddTimestampColumn",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindAddTimestampColumn, Name: "col1"},
				},
			},
			wantErr: false,
		},
		{
			name: "InvalidAddTimestampColumn",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindAddTimestampColumn},
				},
			},
			wantErr: true,
		},
		{
			name: "InvalidAddTimestampColumnValue",
			input: Spec{
				TransformationSpecs: []TransformationSpec{

					{Kind: KindAddTimestampColumn, Value: &[]string{"default"}[0]},
				},
			},
			wantErr: true,
		},
		{
			name: "ValidObfuscateColumns",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindObfuscateColumns, Columns: []string{"col1"}},
				},
			},
			wantErr: false,
		},
		{
			name: "InvalidObfuscateColumnsNoColumns",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindObfuscateColumns},
				},
			},
			wantErr: true,
		},
		{
			name: "ValidRenameColumn",
			input: Spec{
				TransformationSpecs: []TransformationSpec{

					{Kind: KindRenameColumn, Name: "old_col", Value: &[]string{"new_col"}[0]},
				},
			},
			wantErr: false,
		},
		{
			name: "InvalidRenameColumnColumnsProvided",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindRenameColumn, Columns: []string{"col1"}},
				},
			},
			wantErr: true,
		},
		{
			name: "InvalidRenameColumnNoNameNoValue",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindRenameColumn},
				},
			},
			wantErr: true,
		},
		{
			name: "InvalidTransformationKind",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: "invalid_kind"},
				},
			},
			wantErr: true,
		},
		{
			name: "InvalidLowercaseEmptyColumns",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindLowercase, Columns: []string{}},
				},
			},
			wantErr: true,
		},
		{
			name: "ValidLowercase",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindLowercase, Columns: []string{"col1"}},
				},
			},
			wantErr: false,
		},

		{
			name: "ValidDropRows",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindDropRows, Columns: []string{"col1"}},
					{Kind: KindDropRows, Columns: []string{"col1"}, Value: &[]string{"value"}[0]},
				},
			},
			wantErr: false,
		},
		{
			name: "InvalidDropRows",
			input: Spec{
				TransformationSpecs: []TransformationSpec{
					{Kind: KindDropRows},
					{Kind: KindDropRows, Value: &[]string{"value"}[0]},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
