package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestOrganizationJSONSchema(t *testing.T) {
	data, err := jsonschema.Generate(Organization{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true, // missing member_role_name
			Spec: `{}`,
		},
		{
			Name: "empty member_role_name",
			Err:  true,
			Spec: `{"member_role_name": ""}`,
		},
		{
			Name: "null member_role_name",
			Err:  true,
			Spec: `{"member_role_name": null}`,
		},
		{
			Name: "bad member_role_name",
			Err:  true,
			Spec: `{"member_role_name": 123}`,
		},
		{
			Name: "proper member_role_name",
			Spec: `{"member_role_name": "abc"}`,
		},
		{
			Name: "empty admin_account",
			Err:  true, // ~ empty account
			Spec: `{"member_role_name": "abc", "admin_account":{}}`,
		},
		{
			Name: "null admin_account",
			Spec: `{"member_role_name": "abc", "admin_account":null}`,
		},
		{
			Name: "bad admin_account",
			Err:  true,
			Spec: `{"member_role_name": "abc", "admin_account":123}`,
		},
		{
			Name: "empty member_trusted_principal",
			Err:  true, // ~ empty account
			Spec: `{"member_role_name": "abc", "member_trusted_principal":{}}`,
		},
		{
			Name: "null member_trusted_principal",
			Spec: `{"member_role_name": "abc", "member_trusted_principal":null}`,
		},
		{
			Name: "bad member_trusted_principal",
			Err:  true,
			Spec: `{"member_role_name": "abc", "member_trusted_principal":123}`,
		},
		{
			Name: "empty member_role_session_name",
			Spec: `{"member_role_name": "abc", "member_role_session_name":""}`,
		},
		{
			Name: "null member_role_session_name",
			Err:  true,
			Spec: `{"member_role_name": "abc", "member_role_session_name":null}`,
		},
		{
			Name: "bad member_role_session_name",
			Err:  true,
			Spec: `{"member_role_name": "abc", "member_role_session_name":123}`,
		},
		{
			Name: "proper member_role_session_name",
			Spec: `{"member_role_name": "abc", "member_role_session_name":"abc"}`,
		},
		{
			Name: "empty member_external_id",
			Spec: `{"member_role_name": "abc", "member_external_id":""}`,
		},
		{
			Name: "null member_external_id",
			Err:  true,
			Spec: `{"member_role_name": "abc", "member_external_id":null}`,
		},
		{
			Name: "bad member_external_id",
			Err:  true,
			Spec: `{"member_role_name": "abc", "member_external_id":123}`,
		},
		{
			Name: "proper member_external_id",
			Spec: `{"member_role_name": "abc", "member_external_id":"abc"}`,
		},
		{
			Name: "empty member_regions",
			Spec: `{"member_role_name": "abc", "member_regions":[]}`,
		},
		{
			Name: "null member_regions",
			Spec: `{"member_role_name": "abc", "member_regions":null}`,
		},
		{
			Name: "bad member_regions",
			Err:  true,
			Spec: `{"member_role_name": "abc", "member_regions":123}`,
		},
		{
			Name: "empty member_regions entry",
			Err:  true,
			Spec: `{"member_role_name": "abc", "member_regions":[""]}`,
		},
		{
			Name: "null member_regions entry",
			Err:  true,
			Spec: `{"member_role_name": "abc", "member_regions":[null]}`,
		},
		{
			Name: "bad member_regions entry",
			Err:  true,
			Spec: `{"member_role_name": "abc", "member_regions":[123]}`,
		},
		{
			Name: "proper member_regions entry",
			Spec: `{"member_role_name": "abc", "member_regions":["abc"]}`,
		},
		{
			Name: "empty organization_units",
			Spec: `{"member_role_name": "abc", "organization_units":[]}`,
		},
		{
			Name: "null organization_units",
			Spec: `{"member_role_name": "abc", "organization_units":null}`,
		},
		{
			Name: "bad organization_units",
			Err:  true,
			Spec: `{"member_role_name": "abc", "organization_units":123}`,
		},
		{
			Name: "empty organization_units entry",
			Err:  true,
			Spec: `{"member_role_name": "abc", "organization_units":[""]}`,
		},
		{
			Name: "null organization_units entry",
			Err:  true,
			Spec: `{"member_role_name": "abc", "organization_units":[null]}`,
		},
		{
			Name: "bad organization_units entry",
			Err:  true,
			Spec: `{"member_role_name": "abc", "organization_units":[123]}`,
		},
		{
			Name: "bad organization_units entry format",
			Err:  true,
			Spec: `{"member_role_name": "abc", "organization_units":["abc"]}`,
		},
		{
			Name: "proper organization_units entry format",
			Spec: `{"member_role_name": "abc", "organization_units":["ou-abcdefg123-qwerty789","r-qwerty789"]}`,
		},
		{
			Name: "empty skip_organization_units",
			Spec: `{"member_role_name": "abc", "skip_organization_units":[]}`,
		},
		{
			Name: "null skip_organization_units",
			Spec: `{"member_role_name": "abc", "skip_organization_units":null}`,
		},
		{
			Name: "bad skip_organization_units",
			Err:  true,
			Spec: `{"member_role_name": "abc", "skip_organization_units":123}`,
		},
		{
			Name: "empty skip_organization_units entry",
			Err:  true,
			Spec: `{"member_role_name": "abc", "skip_organization_units":[""]}`,
		},
		{
			Name: "null skip_organization_units entry",
			Err:  true,
			Spec: `{"member_role_name": "abc", "skip_organization_units":[null]}`,
		},
		{
			Name: "bad skip_organization_units entry",
			Err:  true,
			Spec: `{"member_role_name": "abc", "skip_organization_units":[123]}`,
		},
		{
			Name: "bad skip_organization_units entry format",
			Err:  true,
			Spec: `{"member_role_name": "abc", "skip_organization_units":["abc"]}`,
		},
		{
			Name: "proper skip_organization_units entry format",
			Spec: `{"member_role_name": "abc", "skip_organization_units":["ou-abcdefg123-qwerty789","r-qwerty789"]}`,
		},
	})
}
