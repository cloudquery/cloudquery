package spec

import "testing"

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
