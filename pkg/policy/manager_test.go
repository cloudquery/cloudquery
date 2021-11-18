package policy

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cloudquery/cloudquery/pkg/config"
)

func TestParseRemotePolicySource(t *testing.T) {
	tests := []struct {
		name      string
		policy    *config.Policy
		expected  *RemotePolicy
		wantErr   bool
		errString string
	}{
		{
			"repository with .git suffix and version",
			&config.Policy{
				Type:    config.Remote,
				Source:  "https://github.com/cloudquery/cloudquery.git",
				Version: "0.0.1",
			},
			&RemotePolicy{
				SourceControl: "https://github.com/",
				Organization:  "cloudquery",
				Repository:    "cloudquery",
				Version:       "0.0.1",
			},
			false,
			"",
		},
		{
			"repository with .git suffix and no version",
			&config.Policy{
				Type:   config.Remote,
				Source: "https://github.com/cloudquery/cloudquery.git",
			},
			&RemotePolicy{
				SourceControl: "https://github.com/",
				Organization:  "cloudquery",
				Repository:    "cloudquery",
				Version:       "",
			},
			false,
			"",
		},
		{
			"repository without .git suffix and version",
			&config.Policy{
				Type:    config.Remote,
				Source:  "https://github.com/cloudquery/cloudquery",
				Version: "0.0.1",
			},
			&RemotePolicy{
				SourceControl: "https://github.com/",
				Organization:  "cloudquery",
				Repository:    "cloudquery",
				Version:       "0.0.1",
			},
			false,
			"",
		},
		{
			"repository without .git suffix and no version",
			&config.Policy{
				Type:   config.Remote,
				Source: "https://github.com/cloudquery/cloudquery",
			},
			&RemotePolicy{
				SourceControl: "https://github.com/",
				Organization:  "cloudquery",
				Repository:    "cloudquery",
			},
			false,
			"",
		},
		{
			"repository without .git suffix and username",
			&config.Policy{
				Type:    config.Remote,
				Source:  "https://cq:cq@github.com/cloudquery/cloudquery",
				Version: "0.0.1",
			},
			&RemotePolicy{
				SourceControl: "https://cq:cq@github.com/",
				Organization:  "cloudquery",
				Repository:    "cloudquery",
				Version:       "0.0.1",
			},
			false,
			"",
		},
		{
			"repository with .git suffix and wrong path 1",
			&config.Policy{
				Type:    config.Remote,
				Source:  "https://cq:cq@github.com/cloudquery/cloudquery/cloud",
				Version: "0.0.1",
			},
			nil,
			true,
			"cloud not parse policy source url",
		},
		{
			"repository with .git suffix and wrong path 2",
			&config.Policy{
				Type:    config.Remote,
				Source:  "https://cq:cq@github.com/cloudquer",
				Version: "0.0.1",
			},
			nil,
			true,
			"cloud not parse policy source url",
		},
		{
			"repository from hub",
			&config.Policy{
				Type:    config.Hub,
				Source:  "aws-cis-1.2",
				Version: "0.0.1",
			},
			&RemotePolicy{
				SourceControl: "https://github.com/",
				Organization:  "cloudquery-policies",
				Repository:    "aws-cis-1.2",
				Version:       "0.0.1",
			},
			false,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			remotePolicy, err := ParsePolicyFromSource(tt.policy)

			if tt.wantErr != (err != nil) {
				t.Errorf("want errors is %v, but have %v, error details: %s", tt.wantErr, err != nil, err)
			}
			if tt.errString != "" {
				assert.Equal(t, err.Error(), tt.errString)
			}
			assert.Equal(t, tt.expected, remotePolicy)
		})
	}
}
