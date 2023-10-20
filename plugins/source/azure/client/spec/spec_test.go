package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestSpec_JSONSchemaExtend(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "extra keyword",
			Err:  true,
			Spec: `{"extra":true}`,
		},
		{
			Name: "empty subscriptions",
			Spec: `{"subscriptions":[]}`,
		},
		{
			Name: "null subscriptions",
			Spec: `{"subscriptions":null}`,
		},
		{
			Name: "bad subscriptions",
			Err:  true,
			Spec: `{"subscriptions":123}`,
		},
		{
			Name: "empty subscriptions entry",
			Err:  true,
			Spec: `{"subscriptions":[""]}`,
		},
		{
			Name: "null subscriptions entry",
			Err:  true,
			Spec: `{"subscriptions":[null]}`,
		},
		{
			Name: "bad subscriptions entry",
			Err:  true,
			Spec: `{"subscriptions":[123]}`,
		},
		{
			Name: "duplicate subscriptions entry",
			Err:  true,
			Spec: `{"subscriptions":["a", "a"]}`,
		},
		{
			Name: "proper subscriptions entries",
			Spec: `{"subscriptions":["a", "b"]}`,
		},
		{
			Name: "empty skip_subscriptions",
			Spec: `{"skip_subscriptions":[]}`,
		},
		{
			Name: "null skip_subscriptions",
			Spec: `{"skip_subscriptions":null}`,
		},
		{
			Name: "bad skip_subscriptions",
			Err:  true,
			Spec: `{"skip_subscriptions":123}`,
		},
		{
			Name: "empty skip_subscriptions entry",
			Err:  true,
			Spec: `{"skip_subscriptions":[""]}`,
		},
		{
			Name: "null skip_subscriptions entry",
			Err:  true,
			Spec: `{"skip_subscriptions":[null]}`,
		},
		{
			Name: "bad skip_subscriptions entry",
			Err:  true,
			Spec: `{"skip_subscriptions":[123]}`,
		},
		{
			Name: "duplicate skip_subscriptions entry",
			Err:  true,
			Spec: `{"skip_subscriptions":["a", "a"]}`,
		},
		{
			Name: "proper skip_subscriptions entries",
			Spec: `{"skip_subscriptions":["a", "b"]}`,
		},
		{
			Name: "empty cloud_name",
			Err:  true,
			Spec: `{"cloud_name":""}`,
		},
		{
			Name: "null cloud_name",
			Err:  true,
			Spec: `{"cloud_name":null}`,
		},
		{
			Name: "bad cloud_name",
			Err:  true,
			Spec: `{"cloud_name":123}`,
		},
		{
			Name: "bad cloud_name value",
			Err:  true,
			Spec: `{"cloud_name":"abc"}`,
		},
		{
			Name: "cloud_name: AzurePublic",
			Spec: `{"cloud_name":"AzurePublic"}`,
		},
		{
			Name: "cloud_name: AzureGovernment",
			Spec: `{"cloud_name":"AzureGovernment"}`,
		},
		{
			Name: "cloud_name: AzureChina",
			Spec: `{"cloud_name":"AzureChina"}`,
		},
		{
			Name: "normalize_ids:false",
			Spec: `{"normalize_ids":false}`,
		},
		{
			Name: "normalize_ids:true",
			Spec: `{"normalize_ids":true}`,
		},
		{
			Name: "null normalize_ids",
			Err:  true,
			Spec: `{"normalize_ids":null}`,
		},
		{
			Name: "bad normalize_ids",
			Err:  true,
			Spec: `{"normalize_ids":123}`,
		},
		{
			Name: "empty oidc_token",
			Err:  true,
			Spec: `{"oidc_token":""}`,
		},
		{
			Name: "null oidc_token",
			Err:  true,
			Spec: `{"oidc_token":null}`,
		},
		{
			Name: "bad oidc_token",
			Err:  true,
			Spec: `{"oidc_token":123}`,
		},
		{
			Name: "proper oidc_token",
			Spec: `{"oidc_token":"a"}`,
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
			Spec: `{"concurrency":"123"}`,
		},
		{
			Name: "proper concurrency",
			Spec: `{"concurrency":123}`,
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
			Spec: `{"discovery_concurrency":"123"}`,
		},
		{
			Name: "proper discovery_concurrency",
			Spec: `{"discovery_concurrency":123}`,
		},
		{
			Name: "proper retry_options.max_retries",
			Spec: `{"retry_options":{"max_retries":1}}`,
		},
		{
			Name: "bad retry_options.max_retries",
			Spec: `{"retry_options":{"max_retries":"1"}}`,
			Err:  true,
		},
		{
			Name: "proper retry_options.try_timeout_seconds",
			Spec: `{"retry_options":{"try_timeout_seconds":1}}`,
		},
		{
			Name: "bad retry_options.try_timeout_seconds",
			Spec: `{"retry_options":{"try_timeout_seconds":-1}}`,
			Err:  true,
		},
		{
			Name: "proper retry_options.retry_delay_seconds",
			Spec: `{"retry_options":{"try_timeout_seconds":1}}`,
		},
		{
			Name: "bad retry_options.retry_delay_seconds",
			Spec: `{"retry_options":{"try_timeout_seconds":-1}}`,
			Err:  true,
		},
		{
			Name: "proper retry_options.max_retry_delay_seconds",
			Spec: `{"retry_options":{"try_timeout_seconds":1}}`,
		},
		{
			Name: "bad retry_options.max_retry_delay_seconds",
			Spec: `{"retry_options":{"try_timeout_seconds":-1}}`,
			Err:  true,
		},
		{
			Name: "empty retry_options.status_codes",
			Spec: `{"retry_options":{"status_codes":[]}}`,
		},
		{
			Name: "proper retry_options.status_codes",
			Spec: `{"retry_options":{"status_codes":[200]}}`,
		},
		{
			Name: "bad retry_options.status_codes",
			Spec: `{"retry_options":{"status_codes":[200, 200]}}`,
			Err:  true,
		},
	})
}
