package cmd

import (
	"testing"

	initCmd "github.com/cloudquery/cloudquery/cmd/init"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/stretchr/testify/assert"
)

func Test_Init(t *testing.T) {
	testCases := []CommandTestCases{
		{
			Name:           "init-no-args",
			Command:        "init",
			ExpectError:    true,
			ExpectedOutput: "Error: requires at least 1 arg(s), only received 0",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testCommand(t, tc)
		})
	}
}

func Test_parseProviderCLIArg(t *testing.T) {
	type args struct {
		providerCLIArg string
	}
	tests := []struct {
		name        string
		args        args
		wantOrg     string
		wantName    string
		wantVersion string
		wantErr     bool
	}{
		{
			name:        "should return default org, latest version and name when given only name",
			args:        args{providerCLIArg: "aws"},
			wantOrg:     registry.DefaultOrganization,
			wantName:    "aws",
			wantVersion: "latest",
			wantErr:     false,
		},
		{
			name:        "should return default org, latest version and name when given name@latest",
			args:        args{providerCLIArg: "aws@latest"},
			wantOrg:     registry.DefaultOrganization,
			wantName:    "aws",
			wantVersion: "latest",
			wantErr:     false,
		},
		{
			name:        "should return default org, specific version and name when given name@version",
			args:        args{providerCLIArg: "aws@1.2.3"},
			wantOrg:     registry.DefaultOrganization,
			wantName:    "aws",
			wantVersion: "v1.2.3",
			wantErr:     false,
		},
		{
			name:        "should return org, latest version and name when given org/name",
			args:        args{providerCLIArg: "org/test"},
			wantOrg:     "org",
			wantName:    "test",
			wantVersion: "latest",
			wantErr:     false,
		},
		{
			name:        "should return org, latest version and name when given org/name@latest",
			args:        args{providerCLIArg: "org/test"},
			wantOrg:     "org",
			wantName:    "test",
			wantVersion: "latest",
			wantErr:     false,
		},
		{
			name:        "should return org, specific version and name when given org/name@version",
			args:        args{providerCLIArg: "org/test@1.2.3"},
			wantOrg:     "org",
			wantName:    "test",
			wantVersion: "v1.2.3",
			wantErr:     false,
		},
		{
			name:        "should fail when version doesn't follow semver",
			args:        args{providerCLIArg: "test@invalid"},
			wantOrg:     "",
			wantName:    "",
			wantVersion: "",
			wantErr:     true,
		},
		{
			name:        "should fail when version starts with a 'v'",
			args:        args{providerCLIArg: "org/test@v1.2.3"},
			wantOrg:     "org",
			wantName:    "test",
			wantVersion: "v1.2.3",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOrg, gotName, gotVersion, err := initCmd.ParseProviderCLIArg(tt.args.providerCLIArg)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.wantOrg, gotOrg)
			assert.Equal(t, tt.wantName, gotName)
			assert.Equal(t, tt.wantVersion, gotVersion)
		})
	}
}
