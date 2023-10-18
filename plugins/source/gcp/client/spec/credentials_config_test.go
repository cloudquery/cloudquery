package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestCredentialsConfigJSONSchema(t *testing.T) {
	schema, err := jsonschema.Generate(CredentialsConfig{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true, // missing target_principal
			Spec: `{}`,
		},
		{
			Name: "empty target_principal",
			Err:  true,
			Spec: `{"target_principal":""}`,
		},
		{
			Name: "null target_principal",
			Err:  true,
			Spec: `{"target_principal":null}`,
		},
		{
			Name: "bad target_principal",
			Err:  true,
			Spec: `{"target_principal":123}`,
		},
		{
			Name: "bad target_principal format",
			Err:  true,
			Spec: `{"target_principal":"some"}`,
		},
		{
			Name: "proper target_principal",
			Spec: `{"target_principal":"a@some"}`,
		},
		{
			Name: "extra field",
			Err:  true,
			Spec: `{"target_principal":"a@some","extra":true}`,
		},
		{
			Name: "empty scopes",
			Spec: `{"target_principal":"a@some", "scopes":[]}`,
		},
		{
			Name: "null scopes",
			Spec: `{"target_principal":"a@some", "scopes":null}`,
		},
		{
			Name: "bad scopes",
			Err:  true,
			Spec: `{"target_principal":"a@some", "scopes":123}`,
		},
		{
			Name: "empty scopes entry",
			Err:  true,
			Spec: `{"target_principal":"a@some", "scopes":[""]}`,
		},
		{
			Name: "null scopes entry",
			Err:  true,
			Spec: `{"target_principal":"a@some", "scopes":[null]}`,
		},
		{
			Name: "bad scopes entry",
			Err:  true,
			Spec: `{"target_principal":"a@some", "scopes":[123]}`,
		},
		{
			Name: "bad scopes entry format",
			Err:  true,
			Spec: `{"target_principal":"a@some", "scopes":["https://www.g00gleapis.com/auth/cloud-platform"]}`,
		},
		{
			Name: "proper scopes entry",
			Spec: `{"target_principal":"a@some", "scopes":["https://www.googleapis.com/auth/cloud-platform.read-only"]}`,
		},

		{
			Name: "empty delegates",
			Spec: `{"target_principal":"a@some", "delegates":[]}`,
		},
		{
			Name: "null delegates",
			Spec: `{"target_principal":"a@some", "delegates":null}`,
		},
		{
			Name: "bad delegates",
			Err:  true,
			Spec: `{"target_principal":"a@some", "delegates":123}`,
		},
		{
			Name: "empty delegates entry",
			Err:  true,
			Spec: `{"target_principal":"a@some", "delegates":[""]}`,
		},
		{
			Name: "null delegates entry",
			Err:  true,
			Spec: `{"target_principal":"a@some", "delegates":[null]}`,
		},
		{
			Name: "bad delegates entry",
			Err:  true,
			Spec: `{"target_principal":"a@some", "delegates":[123]}`,
		},
		{
			Name: "bad delegates entry format",
			Err:  true,
			Spec: `{"target_principal":"a@some", "delegates":["abc"]}`,
		},
		{
			Name: "proper delegates entry",
			Spec: `{"target_principal":"a@some", "delegates":["a@some"]}`,
		},
		{
			Name: "empty subject",
			Err:  true,
			Spec: `{"target_principal":"a@some", "subject":""}`,
		},
		{
			Name: "null subject",
			Err:  true,
			Spec: `{"target_principal":"a@some", "subject":null}`,
		},
		{
			Name: "bad subject",
			Err:  true,
			Spec: `{"target_principal":"a@some", "subject":123}`,
		},
		{
			Name: "proper subject",
			Spec: `{"target_principal":"a@some", "subject":"some"}`,
		},
	})
}
