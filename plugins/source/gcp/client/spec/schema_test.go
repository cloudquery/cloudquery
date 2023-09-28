package spec

import (
	"encoding/json"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/stretchr/testify/require"
)

func TestSpec(t *testing.T) {
	validator, err := plugin.JSONSchemaValidator(JSONSchema())
	require.NoError(t, err)

	type testCase struct {
		name string
		spec string
		err  bool
	}

	for _, tc := range []testCase{
		{
			name: "empty",
			spec: `{}`,
		},
		{
			name: "extra properties",
			spec: `{"extra_field_is_not_welcome":true}`,
			err:  true,
		},
		{
			name: "project_ids",
			spec: `{"project_ids": ["my-project-id"]}`,
		},
		{
			name: "empty project_ids",
			spec: `{"project_ids": []}`,
		},
		{
			name: "bad project_ids",
			spec: `{"project_ids": 3}`,
			err:  true,
		},
		{
			name: "empty project ID",
			spec: `{"project_ids": [""]}`,
			err:  true,
		},
		{
			name: "bad project ID",
			spec: `{"project_ids": [3]}`,
			err:  true,
		},
		{
			name: "folder_ids",
			spec: `{"folder_ids": ["folders/123", "organizations/567"]}`,
		},
		{
			name: "empty folder_ids",
			spec: `{"folder_ids": []}`,
		},
		{
			name: "bad folder_ids",
			spec: `{"folder_ids": 3}`,
			err:  true,
		},
		{
			name: "empty folder ID",
			spec: `{"folder_ids": [""]}`,
			err:  true,
		},
		{
			name: "bad folder ID",
			spec: `{"folder_ids": ["not-a-folder-id"]}`,
			err:  true,
		},
		{
			name: "bad folder ID",
			spec: `{"folder_ids": [3]}`,
			err:  true,
		},
		{
			name: "organization_ids",
			spec: `{"organization_ids": ["my-org-id"]}`,
		},
		{
			name: "empty organization_ids",
			spec: `{"organization_ids": []}`,
		},
		{
			name: "bad organization_ids",
			spec: `{"organization_ids": 3}`,
			err:  true,
		},
		{
			name: "empty organization ID",
			spec: `{"organization_ids": [""]}`,
			err:  true,
		},
		{
			name: "bad organization ID",
			spec: `{"organization_ids": [3]}`,
			err:  true,
		},
		{
			name: "folder_recursion_depth",
			spec: `{"folder_recursion_depth":0}`,
		},
		{
			name: "bad folder_recursion_depth",
			spec: `{"folder_recursion_depth":-1}`,
			err:  true,
		},
		{
			name: "bad folder_recursion_depth",
			spec: `{"folder_recursion_depth":true}`,
			err:  true,
		},
		{
			name: "project_filter",
			spec: `{"project_filter":"some"}`,
		},
		{
			name: "empty project_filter",
			spec: `{"project_filter":""}`,
		},
		{
			name: "bad project_filter",
			spec: `{"project_filter":false}`,
			err:  true,
		},
		{
			name: "folder_ids & project_filter are mutually exclusive",
			spec: `{"folder_ids": ["folders/123", "organizations/567"], "project_filter":"some"}`,
			err:  true,
		},
		{
			name: "organization_filter",
			spec: `{"organization_filter":"some"}`,
		},
		{
			name: "empty organization_filter",
			spec: `{"organization_filter":""}`,
		},
		{
			name: "bad organization_filter",
			spec: `{"organization_filter":false}`,
			err:  true,
		},
		{
			name: "service_account_key_json",
			spec: `{"service_account_key_json":"some"}`,
		},
		{
			name: "empty service_account_key_json",
			spec: `{"service_account_key_json":""}`,
		},
		{
			name: "bad service_account_key_json",
			spec: `{"service_account_key_json":-1}`,
			err:  true,
		},
		{
			name: "backoff_delay",
			spec: `{"backoff_delay":1}`,
		},
		{
			name: "zero backoff_delay",
			spec: `{"backoff_delay":0}`,
		},
		{
			name: "bad backoff_delay",
			spec: `{"backoff_delay":-1}`,
			err:  true,
		},
		{
			name: "bad backoff_delay",
			spec: `{"backoff_delay":true}`,
			err:  true,
		},
		{
			name: "backoff_retries",
			spec: `{"backoff_retries":1}`,
		},
		{
			name: "zero backoff_retries",
			spec: `{"backoff_retries":0}`,
		},
		{
			name: "bad backoff_retries",
			spec: `{"backoff_retries":-1}`,
			err:  true,
		},
		{
			name: "bad backoff_retries",
			spec: `{"backoff_retries":true}`,
			err:  true,
		},
		{
			name: "discovery_concurrency",
			spec: `{"discovery_concurrency":1}`,
		},
		{
			name: "bad discovery_concurrency",
			spec: `{"discovery_concurrency":-1}`,
			err:  true,
		},
		{
			name: "bad discovery_concurrency",
			spec: `{"discovery_concurrency":true}`,
			err:  true,
		},
		{
			name: "enabled_services_only",
			spec: `{"enabled_services_only":true}`,
		},
		{
			name: "bad enabled_services_only",
			spec: `{"enabled_services_only":1}`,
			err:  true,
		},
		{
			name: "concurrency",
			spec: `{"concurrency":1}`,
		},
		{
			name: "bad concurrency",
			spec: `{"concurrency":-1}`,
			err:  true,
		},
		{
			name: "bad concurrency",
			spec: `{"concurrency":false}`,
			err:  true,
		},
		{
			name: "scheduler:dfs",
			spec: `{"scheduler":"dfs"}`,
		},
		{
			name: "scheduler:round-robin",
			spec: `{"scheduler":"round-robin"}`,
		},
		{
			name: "scheduler:shuffle",
			spec: `{"scheduler":"shuffle"}`,
		},
		{
			name: "empty scheduler",
			spec: `{"scheduler":""}`,
			err:  true,
		},
		{
			name: "bad scheduler",
			spec: `{"scheduler":"no-such-thing"}`,
			err:  true,
		},
		{
			name: "bad scheduler",
			spec: `{"scheduler":-1}`,
			err:  true,
		},
		{
			name: "empty service_account_impersonation",
			spec: `{"service_account_impersonation":{}}`,
			err:  true,
		},
		{
			name: "extra field in service_account_impersonation",
			spec: `{"service_account_impersonation":{"target_principal":"a@some"},"extra_field_is_not_welcome":true}`,
			err:  true,
		},
		{
			name: "bad service_account_impersonation",
			spec: `{"service_account_impersonation":false}`,
			err:  true,
		},
		{
			name: "service_account_impersonation.target_principal",
			spec: `{"service_account_impersonation":{"target_principal":"a@some"}}`,
		},
		{
			name: "empty service_account_impersonation.target_principal",
			spec: `{"service_account_impersonation":{"target_principal":""}}`,
			err:  true,
		},
		{
			name: "bad service_account_impersonation.target_principal",
			spec: `{"service_account_impersonation":{"target_principal":"some"}}`,
			err:  true,
		},
		{
			name: "bad service_account_impersonation.target_principal",
			spec: `{"service_account_impersonation":{"target_principal":1}}`,
			err:  true,
		},
		{
			name: "service_account_impersonation.scopes",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "scopes":["https://www.googleapis.com/auth/cloud-platform.read-only"]}}`,
		},
		{
			name: "empty service_account_impersonation.scopes",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "scopes":[]}}`,
		},
		{
			name: "bad service_account_impersonation.scopes",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "scopes":false}}`,
			err:  true,
		},
		{
			name: "empty service_account_impersonation scope",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "scopes":[""]}}`,
			err:  true,
		},
		{
			name: "bad service_account_impersonation scope",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "scopes":["https://www.g00gleapis.com/auth/cloud-platform"]}}`,
			err:  true,
		},
		{
			name: "bad service_account_impersonation scope",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "scopes":[1]}}`,
			err:  true,
		},
		{
			name: "service_account_impersonation.delegates",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "delegates":["a@some"]}}`,
		},
		{
			name: "empty service_account_impersonation.delegates",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "delegates":[]}}`,
		},
		{
			name: "bad service_account_impersonation.delegates",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "delegates":-1}}`,
			err:  true,
		},
		{
			name: "empty service_account_impersonation delegate",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "delegates":[""]}}`,
			err:  true,
		},
		{
			name: "bad service_account_impersonation delegate",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "delegates":["some"]}}`,
			err:  true,
		},
		{
			name: "bad service_account_impersonation delegate",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "delegates":[1]}}`,
			err:  true,
		},
		{
			name: "service_account_impersonation.subject",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "subject":"some"}}`,
		},
		{
			name: "empty service_account_impersonation.subject",
			spec: `{"service_account_impersonation":{"target_principal":"a@some", "subject":""}}`,
			err:  true,
		},
		{
			name: "bad service_account_impersonation.subject",
			spec: `{"service_account_impersonation":{"target_principal":1}}`,
			err:  true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var v any
			require.NoError(t, json.Unmarshal([]byte(tc.spec), &v))
			err := validator.Validate(v)
			if tc.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestJSONSchema(t *testing.T) {
	snap := cupaloy.New(
		cupaloy.SnapshotFileExtension(".json"),
		cupaloy.SnapshotSubdirectory(""),
	)
	require.NoError(t, snap.SnapshotWithName("schema", jsonSchema))
}
