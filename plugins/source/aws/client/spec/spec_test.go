package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
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
				Organization: &Organization{
					ChildAccountRoleName: "test",
					OrganizationUnits:    []string{"ou-1234-12345678"},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid org",
			spec: &Spec{
				Organization: &Organization{
					ChildAccountRoleName: "test",
					OrganizationUnits:    []string{"123"},
				},
			},
			wantErr: true,
		},
		{
			name: "missing member account role name",
			spec: &Spec{
				Organization: &Organization{},
			},
			wantErr: true,
		},
		{
			name: "valid skip ou",
			spec: &Spec{
				Organization: &Organization{
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
				Organization: &Organization{
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
				Organization: &Organization{
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

func TestSpecJSONSchema(t *testing.T) {
	// Accounts, Organization, TableOptions & EventBasedSync are tested separately
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "empty accounts",
			Spec: `{"accounts":[]}`,
		},
		{
			Name: "null accounts",
			Spec: `{"accounts":null}`,
		},
		{
			Name: "bad accounts",
			Err:  true,
			Spec: `{"accounts":123}`,
		},
		{
			Name: "empty accounts entry", // detailed is tested separately
			Err:  true,
			Spec: `{"accounts":[{}]}`,
		},
		{
			Name: "null accounts entry",
			Err:  true,
			Spec: `{"accounts":[null]}`,
		},
		{
			Name: "bad accounts entry",
			Err:  true,
			Spec: `{"accounts":[123]}`,
		},
		{
			Name: "empty org",
			Err:  true, // missing member_role_name
			Spec: `{"org":{}}`,
		},
		{
			Name: "null org",
			Spec: `{"org":null}`,
		},
		{
			Name: "bad org",
			Err:  true,
			Spec: `{"org":123}`,
		},
		// We check that accounts aren't present together with org, though
		{
			Name: "accounts with org",
			Err:  true,
			Spec: `{"org":{"member_role_name":"abc"},"accounts":[{"id":"abc"}]}`,
		},
		{
			Name: "filled in accounts with null org",
			Spec: `{"org":null,"accounts":[{"id":"abc"}]}`,
		},
		{
			Name: "filled in org with null accounts",
			Spec: `{"org":{"member_role_name":"abc"},"accounts":null}`,
		},
		{
			Name: "filled in org with empty accounts",
			Spec: `{"org":{"member_role_name":"abc"},"accounts":[]}`,
		},
		{
			Name: "empty regions",
			Spec: `{"regions":[]}`,
		},
		{
			Name: "null regions",
			Spec: `{"regions":null}`,
		},
		{
			Name: "bad regions type",
			Err:  true,
			Spec: `{"regions":123}`,
		},
		{
			Name: "bad region type",
			Err:  true,
			Spec: `{"regions":[123]}`,
		},
		{
			Name: "empty region",
			Err:  true,
			Spec: `{"regions":["a","b",""]}`,
		},
		{
			Name: "proper regions",
			Spec: `{"regions":["a","b","c"]}`,
		},
		{
			Name: "proper aws_debug",
			Spec: `{"aws_debug":true}`,
		},
		{
			Name: "bad aws_debug",
			Err:  true,
			Spec: `{"aws_debug":123}`,
		},
		{
			Name: "null aws_debug",
			Err:  true,
			Spec: `{"aws_debug":null}`,
		},
		{
			Name: "proper max_retries",
			Spec: `{"max_retries":123}`,
		},
		{
			Name: "bad max_retries",
			Err:  true,
			Spec: `{"max_retries":true}`,
		},
		{
			Name: "null max_retries",
			Spec: `{"max_retries":null}`,
		},
		{
			Name: "proper max_backoff",
			Spec: `{"max_backoff":123}`,
		},
		{
			Name: "bad max_backoff",
			Err:  true,
			Spec: `{"max_backoff":true}`,
		},
		{
			Name: "null max_backoff",
			Spec: `{"max_backoff":null}`,
		},
		{
			Name: "null custom_endpoint_url",
			Err:  true,
			Spec: `{"custom_endpoint_url":null}`,
		},
		{
			Name: "bad custom_endpoint_url",
			Err:  true,
			Spec: `{"custom_endpoint_url":123}`,
		},
		{
			Name: "proper custom_endpoint_url & dependent",
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_signing_region":     "region",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "custom_endpoint_url without custom_endpoint_partition_id",
			Err:  true,
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_signing_region":     "region",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "custom_endpoint_url with empty custom_endpoint_partition_id",
			Err:  true,
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_partition_id":       "",
  "custom_endpoint_signing_region":     "region",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "custom_endpoint_url with null custom_endpoint_partition_id",
			Err:  true,
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_partition_id":       null,
  "custom_endpoint_signing_region":     "region",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "custom_endpoint_url with bad custom_endpoint_partition_id",
			Err:  true,
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_partition_id":       123,
  "custom_endpoint_signing_region":     "region",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "custom_endpoint_url without custom_endpoint_signing_region",
			Err:  true,
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "custom_endpoint_url with empty custom_endpoint_signing_region",
			Err:  true,
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_signing_region":     "",    
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "custom_endpoint_url with null custom_endpoint_signing_region",
			Err:  true,
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_signing_region":     null,    
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "custom_endpoint_url with bad custom_endpoint_signing_region",
			Err:  true,
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_signing_region":     123,    
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "custom_endpoint_url without custom_endpoint_hostname_immutable",
			Err:  true,
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_signing_region":     "region"
}
`,
		},
		{
			Name: "custom_endpoint_url with null custom_endpoint_hostname_immutable",
			Err:  true,
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_signing_region":     "region",
  "custom_endpoint_hostname_immutable": null
}
`,
		},
		{
			Name: "custom_endpoint_url with bad custom_endpoint_hostname_immutable",
			Err:  true,
			Spec: `
{
  "custom_endpoint_url":                "url",
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_signing_region":     "region",
  "custom_endpoint_hostname_immutable": 123
}
`,
		},
		{
			Name: "empty custom_endpoint_url",
			Spec: `
{
  "custom_endpoint_url":                "",
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_signing_region":     "region",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "empty custom_endpoint_url without custom_endpoint_partition_id",
			Spec: `
{
  "custom_endpoint_url":                "",
  "custom_endpoint_signing_region":     "region",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "empty custom_endpoint_url without custom_endpoint_signing_region",
			Spec: `
{
  "custom_endpoint_url":                "",
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "empty custom_endpoint_url without custom_endpoint_hostname_immutable",
			Spec: `
{
  "custom_endpoint_url":                "",
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_signing_region":     "region"
}
`,
		},
		{
			Name: "no custom_endpoint_url",
			Spec: `
{
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_signing_region":     "region",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "no custom_endpoint_url without custom_endpoint_partition_id",
			Spec: `
{
  "custom_endpoint_signing_region":     "region",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "no custom_endpoint_url without custom_endpoint_signing_region",
			Spec: `
{
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_hostname_immutable": true
}
`,
		},
		{
			Name: "no custom_endpoint_url without custom_endpoint_hostname_immutable",
			Spec: `
{
  "custom_endpoint_partition_id":       "id",
  "custom_endpoint_signing_region":     "region"
}
`,
		},
		{
			Name: "proper initialization_concurrency",
			Spec: `{"initialization_concurrency":123}`,
		},
		{
			Name: "zero initialization_concurrency",
			Err:  true,
			Spec: `{"initialization_concurrency":0}`,
		},
		{
			Name: "bad initialization_concurrency",
			Err:  true,
			Spec: `{"initialization_concurrency":-1}`,
		},
		{
			Name: "float initialization_concurrency",
			Err:  true,
			Spec: `{"initialization_concurrency":4.5}`,
		},
		{
			Name: "null initialization_concurrency",
			Err:  true,
			Spec: `{"initialization_concurrency":null}`,
		},
		{
			Name: "proper concurrency",
			Spec: `{"concurrency":123}`,
		},
		{
			Name: "zero concurrency",
			Err:  true,
			Spec: `{"concurrency":0}`,
		},
		{
			Name: "bad concurrency",
			Err:  true,
			Spec: `{"concurrency":-1}`,
		},
		{
			Name: "float concurrency",
			Err:  true,
			Spec: `{"concurrency":4.5}`,
		},
		{
			Name: "null concurrency",
			Err:  true,
			Spec: `{"concurrency":null}`,
		},
		{
			Name: "false use_paid_apis",
			Spec: `{"use_paid_apis":false}`,
		},
		{
			Name: "true use_paid_apis",
			Spec: `{"use_paid_apis":true}`,
		},
		{
			Name: "null use_paid_apis",
			Err:  true,
			Spec: `{"use_paid_apis":null}`,
		},
		{
			Name: "bad use_paid_apis type",
			Err:  true,
			Spec: `{"use_paid_apis":123}`,
		},
		// Scheduler tests are included for completeness’s sake, but should be done in scheduler package instead
		{
			Name: "dfs scheduler",
			Spec: `{"scheduler":"dfs"}`,
		},
		{
			Name: "round-robin scheduler",
			Spec: `{"scheduler":"round-robin"}`,
		},
		{
			Name: "shuffle scheduler",
			Spec: `{"scheduler":"shuffle"}`,
		},
		{
			Name: "empty scheduler",
			Err:  true,
			Spec: `{"scheduler":""}`,
		},
		{
			Name: "bad scheduler",
			Err:  true,
			Spec: `{"scheduler":"bad"}`,
		},
		{
			Name: "bad scheduler type",
			Err:  true,
			Spec: `{"scheduler":123}`,
		},
		{
			Name: "null scheduler type",
			Err:  false,
			Spec: `{"scheduler":null}`,
		},
	})
}
