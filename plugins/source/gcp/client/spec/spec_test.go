package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "extra properties",
			Spec: `{"extra_field_is_not_welcome":true}`,
			Err:  true,
		},
		{
			Name: "empty project_ids",
			Spec: `{"project_ids": []}`,
		},
		{
			Name: "null project_ids",
			Spec: `{"project_ids": null}`,
		},
		{
			Name: "bad project_ids",
			Err:  true,
			Spec: `{"project_ids": 123}`,
		},
		{
			Name: "empty project_ids entry",
			Err:  true,
			Spec: `{"project_ids": [""]}`,
		},
		{
			Name: "null project_ids entry",
			Err:  true,
			Spec: `{"project_ids": [null]}`,
		},
		{
			Name: "bad project_ids entry",
			Err:  true,
			Spec: `{"project_ids": [123]}`,
		},
		{
			Name: "proper project_ids",
			Spec: `{"project_ids": ["a-123"]}`,
		},
		{
			Name: "empty folder_ids",
			Spec: `{"folder_ids": []}`,
		},
		{
			Name: "null folder_ids",
			Spec: `{"folder_ids": null}`,
		},
		{
			Name: "bad folder_ids",
			Err:  true,
			Spec: `{"folder_ids": 123}`,
		},
		{
			Name: "empty folder_ids entry",
			Err:  true,
			Spec: `{"folder_ids": [""]}`,
		},
		{
			Name: "null folder_ids entry",
			Err:  true,
			Spec: `{"folder_ids": [null]}`,
		},
		{
			Name: "bad folder_ids entry",
			Err:  true,
			Spec: `{"folder_ids": [123]}`,
		},
		{
			Name: "malformed folder_ids entry",
			Err:  true,
			Spec: `{"folder_ids": ["not-a-folder-id"]}`,
		},
		{
			Name: "proper folder_ids",
			Spec: `{"folder_ids": ["folders/123", "organizations/567"]}`,
		},
		{
			Name: "empty organization_ids",
			Spec: `{"organization_ids": []}`,
		},
		{
			Name: "null organization_ids",
			Spec: `{"organization_ids": null}`,
		},
		{
			Name: "bad organization_ids",
			Err:  true,
			Spec: `{"organization_ids": 123}`,
		},
		{
			Name: "empty organization_ids entry",
			Err:  true,
			Spec: `{"organization_ids": [""]}`,
		},
		{
			Name: "null organization_ids entry",
			Err:  true,
			Spec: `{"organization_ids": [null]}`,
		},
		{
			Name: "bad organization_ids entry",
			Err:  true,
			Spec: `{"organization_ids": [123]}`,
		},
		{
			Name: "proper organization_ids",
			Spec: `{"organization_ids": ["my-org-id"]}`,
		},
		{
			Name: "null folder_recursion_depth",
			Spec: `{"folder_recursion_depth":null}`,
		},
		{
			Name: "bad folder_recursion_depth",
			Err:  true,
			Spec: `{"folder_recursion_depth":-1}`,
		},
		{
			Name: "bad folder_recursion_depth",
			Err:  true,
			Spec: `{"folder_recursion_depth":true}`,
		},
		{
			Name: "zero folder_recursion_depth",
			Spec: `{"folder_recursion_depth":0}`,
		},
		{
			Name: "proper folder_recursion_depth",
			Spec: `{"folder_recursion_depth":123}`,
		},
		{
			Name: "empty project_filter",
			Spec: `{"project_filter":""}`,
		},
		{
			Name: "null project_filter",
			Err:  true,
			Spec: `{"project_filter":null}`,
		},
		{
			Name: "bad project_filter",
			Err:  true,
			Spec: `{"project_filter":123}`,
		},
		{
			Name: "proper project_filter",
			Spec: `{"project_filter":"some"}`,
		},
		{
			Name: "folder_ids & project_filter are mutually exclusive",
			Err:  true,
			Spec: `{"folder_ids": ["folders/123", "organizations/567"], "project_filter":"some"}`,
		},
		{
			Name: "folder_ids & empty project_filter",
			Spec: `{"folder_ids": ["folders/123", "organizations/567"], "project_filter":""}`,
		},
		{
			Name: "empty folder_ids & project_filter",
			Spec: `{"folder_ids": [], "project_filter":"some"}`,
		},
		{
			Name: "null folder_ids & project_filter",
			Spec: `{"folder_ids": null, "project_filter":"some"}`,
		},
		{
			Name: "empty organization_filter",
			Spec: `{"organization_filter":""}`,
		},
		{
			Name: "null organization_filter",
			Err:  true,
			Spec: `{"organization_filter":null}`,
		},
		{
			Name: "bad organization_filter",
			Err:  true,
			Spec: `{"organization_filter":123}`,
		},
		{
			Name: "proper organization_filter",
			Spec: `{"organization_filter":"some"}`,
		},
		{
			Name: "empty service_account_key_json",
			Spec: `{"service_account_key_json":""}`,
		},
		{
			Name: "null service_account_key_json",
			Err:  true,
			Spec: `{"service_account_key_json":null}`,
		},
		{
			Name: "bad service_account_key_json",
			Err:  true,
			Spec: `{"service_account_key_json":123}`,
		},
		{
			Name: "proper service_account_key_json",
			Spec: `{"service_account_key_json":"some"}`,
		},
		{
			Name: "zero backoff_delay",
			Spec: `{"backoff_delay":0}`,
		},
		{
			Name: "null backoff_delay",
			Err:  true,
			Spec: `{"backoff_delay":null}`,
		},
		{
			Name: "bad backoff_delay",
			Err:  true,
			Spec: `{"backoff_delay":-1}`,
		},
		{
			Name: "bad backoff_delay type",
			Err:  true,
			Spec: `{"backoff_delay":false}`,
		},
		{
			Name: "proper backoff_delay",
			Spec: `{"backoff_delay":1}`,
		},
		{
			Name: "zero backoff_retries",
			Spec: `{"backoff_retries":0}`,
		},
		{
			Name: "null backoff_retries",
			Err:  true,
			Spec: `{"backoff_retries":null}`,
		},
		{
			Name: "bad backoff_retries",
			Err:  true,
			Spec: `{"backoff_retries":-1}`,
		},
		{
			Name: "bad backoff_retries type",
			Err:  true,
			Spec: `{"backoff_retries":false}`,
		},
		{
			Name: "proper backoff_retries",
			Spec: `{"backoff_retries":1}`,
		},
		{
			Name: "zero discovery_concurrency",
			Err:  true,
			Spec: `{"discovery_concurrency":0}`,
		},
		{
			Name: "null discovery_concurrency",
			Err:  true,
			Spec: `{"discovery_concurrency":null}`,
		},
		{
			Name: "bad discovery_concurrency",
			Err:  true,
			Spec: `{"discovery_concurrency":-1}`,
		},
		{
			Name: "bad discovery_concurrency type",
			Err:  true,
			Spec: `{"discovery_concurrency":false}`,
		},
		{
			Name: "proper discovery_concurrency",
			Spec: `{"discovery_concurrency":1}`,
		},
		{
			Name: "enabled_services_only:false",
			Spec: `{"enabled_services_only":false}`,
		},
		{
			Name: "enabled_services_only:true",
			Spec: `{"enabled_services_only":true}`,
		},
		{
			Name: "null enabled_services_only",
			Err:  true,
			Spec: `{"enabled_services_only":null}`,
		},
		{
			Name: "bad enabled_services_only",
			Err:  true,
			Spec: `{"enabled_services_only":123}`,
		},
		{
			Name: "zero concurrency",
			Err:  true,
			Spec: `{"concurrency":0}`,
		},
		{
			Name: "null concurrency",
			Err:  true,
			Spec: `{"concurrency":null}`,
		},
		{
			Name: "bad concurrency",
			Err:  true,
			Spec: `{"concurrency":-1}`,
		},
		{
			Name: "bad concurrency type",
			Err:  true,
			Spec: `{"concurrency":false}`,
		},
		{
			Name: "proper concurrency",
			Spec: `{"concurrency":1}`,
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
			Err:  true,
			Spec: `{"scheduler":null}`,
		},
		// ServiceAccountImpersonation(CredentialsConfig) is tested separately
		{
			Name: "null service_account_impersonation",
			Spec: `{"service_account_impersonation":null}`,
		},
		{
			Name: "bad service_account_impersonation",
			Err:  true,
			Spec: `{"service_account_impersonation":123}`,
		},
	})
}
