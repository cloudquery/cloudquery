package client

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

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
		{
			name: "neither private key and private key path specified in configuration",
			spec: Spec{
				AppAuth: []AppAuthSpec{
					{
						Org:            "org1",
						AppID:          "123",
						InstallationID: "456",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "both private key and private key path specified in configuration",
			spec: Spec{
				AppAuth: []AppAuthSpec{
					{
						Org:            "org1",
						AppID:          "123",
						PrivateKeyPath: "path/to/private/key",
						PrivateKey:     "private key",
						InstallationID: "456",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "valid repo",
			spec: Spec{
				AccessToken: "token",
				Repos:       []string{"myorg/myrepo"},
			},
			wantErr: false,
		},
		{
			name: "repo validation error - should be <org>/<repo>",
			spec: Spec{
				AccessToken: "token",
				Repos:       []string{"bad-repo"},
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
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec",
			Spec: `{}`,
			Err:  true,
		},
		{
			Name: "valid spec - orgs",
			Spec: `{"access_token": "token", "orgs": ["org1", "org2"]}`,
		},
		{
			Name: "valid spec - repos",
			Spec: `{"access_token": "token", "repos": ["org1/repo1", "org2/repo2"]}`,
		},
		{
			Name: "valid spec - app auth",
			Spec: `{"app_auth": [{"org": "org1", "app_id": "123", "private_key_path": "path/to/private/key", "installation_id": "456"}], "orgs": ["org1", "org2"]}`,
		},
		{
			Name: "invalid spec - missing orgs or repos",
			Spec: `{"access_token": "token"}`,
			Err:  true,
		},
		{
			Name: "missing base url in enterprise settings",
			Spec: `{"access_token":"token", "enterprise": {"upload_url": "https://upload.example.com"}}`,
			Err:  true,
		},
		{
			Name: "missing upload url in enterprise settings",
			Spec: `{"access_token":"token", "enterprise": {"base_url": "https://base.example.com"}}`,
			Err:  true,
		},
		{
			Name: "missing org in app auth configuration",
			Spec: `{"orgs": ["org1", "org2"],"app_auth": [{"app_id": "123", "private_key_path": "path/to/private/key", "installation_id": "456"}]}`,
			Err:  true,
		},
		{
			Name: "neither private key and private key path specified in configuration",
			Spec: `{"orgs": ["org1", "org2"],"app_auth": [{"org":"org1", "app_id": "123", "installation_id": "456"}]}`,
			Err:  true,
		},
		{
			Name: "both private key and path specified in configuration",
			Spec: `{"orgs": ["org1", "org2"],"app_auth": [{"org":"org1", "app_id": "123", "private_key_path":"path/to/private/key","private_key":"key", "installation_id": "456"}]}`,
			Err:  true,
		},
		{
			Name: "valid repo",
			Spec: `{"access_token": "token", "repos": ["org1/repo1", "org2/repo2"]}`,
			Err:  false,
		},
		{
			Name: "repo validation error - should be <org>/<repo>",
			Spec: `{"access_token": "token", "repos": ["bad-repo"]}`,
			Err:  true,
		},
	})
}
