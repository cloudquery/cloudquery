package spec

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/require"
)

func TestOrgJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true,
			Spec: `{"org":{}}`,
		},
		{
			Name: "null",
			Spec: `{"org":null}`,
		},
		{
			Name: "bad",
			Err:  true,
			Spec: `{"org":123}`,
		},
		{
			Name: "proper",
			Spec: func() string {
				var input Org
				require.NoError(t, faker.FakeObject(&input))

				ou := []string{"ou-abcdefg123-qwerty789", "r-qwerty789"}
				input.OrganizationUnits = ou
				input.SkipOrganizationalUnits = ou

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.AdminAccount.RoleARN = randomARN.String()
				input.MemberCredentials.RoleARN = randomARN.String()

				return `{"org":` + jsonschema.WithRemovedKeys(t, &input) + `}`
			}(),
		},
		{
			Name: "missing member_role_name",
			Err:  true,
			Spec: func() string {
				var input Org
				require.NoError(t, faker.FakeObject(&input))

				ou := []string{"ou-abcdefg123-qwerty789", "r-qwerty789"}
				input.OrganizationUnits = ou
				input.SkipOrganizationalUnits = ou

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.AdminAccount.RoleARN = randomARN.String()
				input.MemberCredentials.RoleARN = randomARN.String()

				return `{"org":` + jsonschema.WithRemovedKeys(t, &input, "member_role_name") + `}`
			}(),
		},
		{
			Name: "empty member_role_name",
			Err:  true,
			Spec: func() string {
				var input Org
				require.NoError(t, faker.FakeObject(&input))

				ou := []string{"ou-abcdefg123-qwerty789", "r-qwerty789"}
				input.OrganizationUnits = ou
				input.SkipOrganizationalUnits = ou

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.AdminAccount.RoleARN = randomARN.String()
				input.MemberCredentials.RoleARN = randomARN.String()

				input.ChildAccountRoleName = ""

				return `{"org":` + jsonschema.WithRemovedKeys(t, &input) + `}`
			}(),
		},
		{
			Name: "null organization_units",
			Spec: func() string {
				var input Org
				require.NoError(t, faker.FakeObject(&input))

				ou := []string{"ou-abcdefg123-qwerty789", "r-qwerty789"}
				input.OrganizationUnits = nil
				input.SkipOrganizationalUnits = ou

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.AdminAccount.RoleARN = randomARN.String()
				input.MemberCredentials.RoleARN = randomARN.String()

				return `{"org":` + jsonschema.WithRemovedKeys(t, &input) + `}`
			}(),
		},
		{
			Name: "empty organization_units",
			Spec: func() string {
				var input Org
				require.NoError(t, faker.FakeObject(&input))

				ou := []string{"ou-abcdefg123-qwerty789", "r-qwerty789"}
				input.OrganizationUnits = []string{}
				input.SkipOrganizationalUnits = ou

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.AdminAccount.RoleARN = randomARN.String()
				input.MemberCredentials.RoleARN = randomARN.String()

				return `{"org":` + jsonschema.WithRemovedKeys(t, &input) + `}`
			}(),
		},
		{
			Name: "bad organization_units",
			Err:  true,
			Spec: func() string {
				var input Org
				require.NoError(t, faker.FakeObject(&input))

				ou := []string{"ou-abcdefg123-qwerty789", "r-qwerty789"}
				input.SkipOrganizationalUnits = ou

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.AdminAccount.RoleARN = randomARN.String()
				input.MemberCredentials.RoleARN = randomARN.String()

				return `{"org":` + jsonschema.WithRemovedKeys(t, &input) + `}`
			}(),
		},
		{
			Name: "null skip_organization_units",
			Spec: func() string {
				var input Org
				require.NoError(t, faker.FakeObject(&input))

				ou := []string{"ou-abcdefg123-qwerty789", "r-qwerty789"}
				input.OrganizationUnits = ou
				input.SkipOrganizationalUnits = nil

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.AdminAccount.RoleARN = randomARN.String()
				input.MemberCredentials.RoleARN = randomARN.String()

				return `{"org":` + jsonschema.WithRemovedKeys(t, &input) + `}`
			}(),
		},
		{
			Name: "empty skip_organization_units",
			Spec: func() string {
				var input Org
				require.NoError(t, faker.FakeObject(&input))

				ou := []string{"ou-abcdefg123-qwerty789", "r-qwerty789"}
				input.OrganizationUnits = ou
				input.SkipOrganizationalUnits = []string{}

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.AdminAccount.RoleARN = randomARN.String()
				input.MemberCredentials.RoleARN = randomARN.String()

				return `{"org":` + jsonschema.WithRemovedKeys(t, &input) + `}`
			}(),
		},
		{
			Name: "bad skip_organization_units",
			Err:  true,
			Spec: func() string {
				var input Org
				require.NoError(t, faker.FakeObject(&input))

				ou := []string{"ou-abcdefg123-qwerty789", "r-qwerty789"}
				input.OrganizationUnits = ou

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.AdminAccount.RoleARN = randomARN.String()
				input.MemberCredentials.RoleARN = randomARN.String()

				return `{"org":` + jsonschema.WithRemovedKeys(t, &input) + `}`
			}(),
		},
		{
			Name: "org with null accounts",
			Spec: func() string {
				var org Org
				require.NoError(t, faker.FakeObject(&org))

				ou := []string{"ou-abcdefg123-qwerty789", "r-qwerty789"}
				org.OrganizationUnits = ou
				org.SkipOrganizationalUnits = ou

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				org.AdminAccount.RoleARN = randomARN.String()
				org.MemberCredentials.RoleARN = randomARN.String()

				return `{"accounts":null,"org":` + jsonschema.WithRemovedKeys(t, &org) + `}`
			}(),
		},
		{
			Name: "org with empty accounts",
			Spec: func() string {
				var org Org
				require.NoError(t, faker.FakeObject(&org))

				ou := []string{"ou-abcdefg123-qwerty789", "r-qwerty789"}
				org.OrganizationUnits = ou
				org.SkipOrganizationalUnits = ou

				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				org.AdminAccount.RoleARN = randomARN.String()
				org.MemberCredentials.RoleARN = randomARN.String()

				return `{"accounts":[],"org":` + jsonschema.WithRemovedKeys(t, &org) + `}`
			}(),
		},
	})
}
