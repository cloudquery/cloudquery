package client

import "testing"

func TestSpec_Validate(t *testing.T) {
	tests := []struct {
		name    string
		spec    Spec
		wantErr bool
	}{
		{
			name:    "empty spec",
			spec:    Spec{},
			wantErr: true,
		},
		{
			name: "valid spec - orgs",
			spec: Spec{
				AccessToken: "token",
				Orgs:        []string{"org1", "org2"},
			},
		},
		{
			name: "valid spec - repos",
			spec: Spec{
				AccessToken: "token",
				Repos:       []string{"org1/repo1", "org2/repo2"},
			},
		},
		{
			name: "valid spec - app auth",
			spec: Spec{
				AppAuth: []AppAuthSpec{
					{
						Org:            "org1",
						AppID:          "123",
						PrivateKeyPath: "path/to/private/key",
						InstallationID: "456",
					},
				},
				Orgs: []string{"org1", "org2"},
			},
		},
		{
			name: "invalid spec - missing orgs or repos",
			spec: Spec{
				AccessToken: "token",
			},
			wantErr: true,
		},
		{
			name: "missing base url in enterprise settings",
			spec: Spec{
				EnterpriseSettings: &EnterpriseSettings{
					UploadURL: "https://upload.example.com",
				},
			},
			wantErr: true,
		},
		{
			name: "missing upload url in enterprise settings",
			spec: Spec{
				EnterpriseSettings: &EnterpriseSettings{
					BaseURL: "https://base.example.com",
				},
			},
			wantErr: true,
		},
		{
			name: "missing org in app auth configuration",
			spec: Spec{
				AppAuth: []AppAuthSpec{
					{
						AppID:          "123",
						PrivateKeyPath: "path/to/private/key",
						PrivateKey:     "private key",
						InstallationID: "456",
					},
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
