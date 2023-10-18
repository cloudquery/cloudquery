package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestAccountJSONSchema(t *testing.T) {
	data, err := jsonschema.Generate(Account{})
	require.NoError(t, err)
	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true, // missing id
			Spec: `{}`,
		},
		{
			Name: "empty id",
			Err:  true,
			Spec: `{"id":""}`,
		},
		{
			Name: "null id",
			Err:  true,
			Spec: `{"id":null}`,
		},
		{
			Name: "bad id",
			Err:  true,
			Spec: `{"id":123}`,
		},
		{
			Name: "proper id",
			Spec: `{"id":"abc"}`,
		},
		{
			Name: "empty account_name",
			Spec: `{"id":"abc","account_name":""}`,
		},
		{
			Name: "null account_name",
			Err:  true,
			Spec: `{"id":"abc","account_name":null}`,
		},
		{
			Name: "bad account_name",
			Err:  true,
			Spec: `{"id":"abc","account_name":123}`,
		},
		{
			Name: "proper account_name",
			Spec: `{"id":"abc","account_name":"abc"}`,
		},
		{
			Name: "empty local_profile",
			Spec: `{"id":"abc","local_profile":""}`,
		},
		{
			Name: "null local_profile",
			Err:  true,
			Spec: `{"id":"abc","local_profile":null}`,
		},
		{
			Name: "bad local_profile",
			Err:  true,
			Spec: `{"id":"abc","local_profile":123}`,
		},
		{
			Name: "proper local_profile",
			Spec: `{"id":"abc","local_profile":"abc"}`,
		},
		{
			Name: "empty role_arn",
			Spec: `{"id":"abc","role_arn":""}`,
		},
		{
			Name: "null role_arn",
			Err:  true,
			Spec: `{"id":"abc","role_arn":null}`,
		},
		{
			Name: "bad role_arn type",
			Err:  true,
			Spec: `{"id":"abc","role_arn":123}`,
		},
		{
			Name: "bad role_arn",
			Err:  true,
			Spec: `{"id":"abc","role_arn":"abc"}`,
		},
		{
			Name: "proper role_arn",
			Spec: `{"id":"abc","role_arn":"arn:1:2:3:4:5"}`,
		},
		{
			Name: "empty role_session_name",
			Spec: `{"id":"abc","role_session_name":""}`,
		},
		{
			Name: "null role_session_name",
			Err:  true,
			Spec: `{"id":"abc","role_session_name":null}`,
		},
		{
			Name: "bad role_session_name",
			Err:  true,
			Spec: `{"id":"abc","role_session_name":123}`,
		},
		{
			Name: "proper role_session_name",
			Spec: `{"id":"abc","role_session_name":"abc"}`,
		},
		{
			Name: "empty external_id",
			Spec: `{"id":"abc","external_id":""}`,
		},
		{
			Name: "null external_id",
			Err:  true,
			Spec: `{"id":"abc","external_id":null}`,
		},
		{
			Name: "bad external_id",
			Err:  true,
			Spec: `{"id":"abc","external_id":123}`,
		},
		{
			Name: "proper external_id",
			Spec: `{"id":"abc","external_id":"abc"}`,
		},
		{
			Name: "empty default_region",
			Err:  true,
			Spec: `{"id":"abc","default_region":""}`,
		},
		{
			Name: "null default_region",
			Err:  true,
			Spec: `{"id":"abc","default_region":null}`,
		},
		{
			Name: "bad default_region",
			Err:  true,
			Spec: `{"id":"abc","default_region":123}`,
		},
		{
			Name: "proper default_region",
			Spec: `{"id":"abc","default_region":"abc"}`,
		},
		{
			Name: "empty regions",
			Spec: `{"id":"abc","regions":[]}`,
		},
		{
			Name: "null regions",
			Spec: `{"id":"abc","regions":null}`,
		},
		{
			Name: "bad regions",
			Err:  true,
			Spec: `{"id":"abc","regions":123}`,
		},
		{
			Name: "empty regions entry",
			Err:  true,
			Spec: `{"id":"abc","regions":[""]}`,
		},
		{
			Name: "null regions entry",
			Err:  true,
			Spec: `{"id":"abc","regions":[null]}`,
		},
		{
			Name: "bad regions entry",
			Err:  true,
			Spec: `{"id":"abc","regions":[123]}`,
		},
		{
			Name: "proper regions entry",
			Spec: `{"id":"abc","regions":["abc"]}`,
		},
	})
}
