package spec

import (
	"encoding/json"
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/stretchr/testify/require"
)

func TestSpecValidate(t *testing.T) {
	tests := []struct {
		name    string
		spec    *Spec
		wantErr bool
	}{
		{
			name: "valid accounts",
			spec: &Spec{
				Accounts: []Account{
					{ID: "123456789012"},
					{ID: "cq-playground"},
				},
			},
			wantErr: false,
		},
		{
			name: "valid org",
			spec: &Spec{
				Organization: &Org{
					ChildAccountRoleName: "test",
					OrganizationUnits:    []string{"ou-1234-12345678"},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid org",
			spec: &Spec{
				Organization: &Org{
					ChildAccountRoleName: "test",
					OrganizationUnits:    []string{"123"},
				},
			},
			wantErr: true,
		},
		{
			name: "missing member account role name",
			spec: &Spec{
				Organization: &Org{},
			},
			wantErr: true,
		},
		{
			name: "valid skip ou",
			spec: &Spec{
				Organization: &Org{
					ChildAccountRoleName:    "test",
					OrganizationUnits:       []string{"ou-1234-12345678"},
					SkipOrganizationalUnits: []string{"ou-1234-45678901"},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid skip ou",
			spec: &Spec{
				Organization: &Org{
					ChildAccountRoleName:    "test",
					OrganizationUnits:       []string{"ou-1234-12345678"},
					SkipOrganizationalUnits: []string{"456"},
				},
			},
			wantErr: true,
		},
		{
			name: "both account and org",
			spec: &Spec{
				Accounts: []Account{
					{ID: "123456789012"},
				},
				Organization: &Org{
					ChildAccountRoleName: "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.spec.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Spec.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJSONSchema(t *testing.T) {
	validator, err := plugin.JSONSchemaValidator(JSONSchema)
	require.NoError(t, err)

	type testCase struct {
		name string
		spec string
		err  bool
	}

	for _, tc := range []testCase{} {
		t.Run(tc.name, func(t *testing.T) {
			var v any
			require.NoError(t, json.Unmarshal([]byte(tc.spec), &v))
			err := validator.Validate(v)
			if tc.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestEnsureJSONSchema(t *testing.T) {
	data, err := jsonschema.Generate(new(Spec))
	require.NoError(t, err)
	require.JSONEqf(t, string(data), JSONSchema, "new schema should be:\n%s\n", string(data))
}
