package spec

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/require"
)

func TestAccountJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{"accounts":[]}`,
		},
		{
			Name: "null",
			Spec: `{"accounts":null}`,
		},
		{
			Name: "proper",
			Spec: func() string {
				var input Account
				require.NoError(t, faker.FakeObject(&input))

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.RoleARN = randomARN.String()

				return `{"accounts":[` + jsonschema.WithRemovedKeys(t, &input) + `]}`
			}(),
		},
		{
			Name: "bad role_arn",
			Err:  true,
			Spec: func() string {
				var input Account
				require.NoError(t, faker.FakeObject(&input))
				return `{"accounts":[` + jsonschema.WithRemovedKeys(t, &input) + `]}`
			}(),
		},
		{
			Name: "id missing",
			Err:  true,
			Spec: func() string {
				var input Account
				require.NoError(t, faker.FakeObject(&input))

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.RoleARN = randomARN.String()

				return `{"accounts":[` + jsonschema.WithRemovedKeys(t, &input, "id") + `]}`
			}(),
		},
		{
			Name: "empty region",
			Err:  true,
			Spec: func() string {
				var input Account
				require.NoError(t, faker.FakeObject(&input))

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.RoleARN = randomARN.String()

				input.Regions = []string{""}
				return `{"accounts":[` + jsonschema.WithRemovedKeys(t, &input) + `]}`
			}(),
		},
		{
			Name: "accounts with null org",
			Spec: func() string {
				var account Account
				require.NoError(t, faker.FakeObject(&account))

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				account.RoleARN = randomARN.String()

				return `{"org":null,"accounts":[` + jsonschema.WithRemovedKeys(t, &account) + `]}`
			}(),
		},
	})
}
