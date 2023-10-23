package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestRetryOptionsJSONSchema(t *testing.T) {
	schema, err := jsonschema.Generate(RetryOptions{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "zero max_retries",
			Spec: `{"max_retries":0}`,
		},
		{
			Name: "null max_retries",
			Spec: `{"max_retries":null}`,
		},
		{
			Name: "bad max_retries",
			Err:  true,
			Spec: `{"max_retries":true}`,
		},
		{
			Name: "positive max_retries",
			Spec: `{"max_retries":100}`,
		},
		{
			Name: "negative max_retries",
			Spec: `{"max_retries":100}`,
		},
		{
			Name: "null try_timeout_seconds",
			Spec: `{"try_timeout_seconds":null}`,
		},
		{
			Name: "bad try_timeout_seconds",
			Err:  true,
			Spec: `{"try_timeout_seconds":"abc"}`,
		},
		{
			Name: "zero try_timeout_seconds",
			Spec: `{"try_timeout_seconds":0}`,
		},
		{
			Name: "positive try_timeout_seconds",
			Spec: `{"try_timeout_seconds":100}`,
		},
		{
			Name: "negative try_timeout_seconds",
			Err:  true,
			Spec: `{"try_timeout_seconds":-100}`,
		},
		{
			Name: "null retry_delay_seconds",
			Spec: `{"retry_delay_seconds":null}`,
		},
		{
			Name: "bad retry_delay_seconds",
			Err:  true,
			Spec: `{"retry_delay_seconds":"abc"}`,
		},
		{
			Name: "zero retry_delay_seconds",
			Spec: `{"retry_delay_seconds":0}`,
		},
		{
			Name: "positive retry_delay_seconds",
			Spec: `{"retry_delay_seconds":100}`,
		},
		{
			Name: "negative retry_delay_seconds",
			Err:  true,
			Spec: `{"retry_delay_seconds":-100}`,
		},
		{
			Name: "null max_retry_delay_seconds",
			Spec: `{"max_retry_delay_seconds":null}`,
		},
		{
			Name: "bad max_retry_delay_seconds",
			Err:  true,
			Spec: `{"max_retry_delay_seconds":"abc"}`,
		},
		{
			Name: "zero max_retry_delay_seconds",
			Spec: `{"max_retry_delay_seconds":0}`,
		},
		{
			Name: "positive max_retry_delay_seconds",
			Spec: `{"max_retry_delay_seconds":100}`,
		},
		{
			Name: "negative max_retry_delay_seconds",
			Err:  true,
			Spec: `{"max_retry_delay_seconds":-100}`,
		},
		{
			Name: "empty status_codes",
			Spec: `{"status_codes":[]}`,
		},
		{
			Name: "null status_codes",
			Spec: `{"status_codes":null}`,
		},
		{
			Name: "bad status_codes",
			Err:  true,
			Spec: `{"status_codes":123}`,
		},
		{
			Name: "zero status_codes entry",
			Spec: `{"status_codes":[0]}`,
		},
		{
			Name: "null status_codes entry",
			Err:  true,
			Spec: `{"status_codes":[null]}`,
		},
		{
			Name: "bad status_codes entry",
			Err:  true,
			Spec: `{"status_codes":[true]}`,
		},
		{
			Name: "proper status_codes",
			Spec: `{"status_codes":[0,1,2,3,4,5,6,7,8]}`,
		},
		{
			Name: "duplicate status_codes entries",
			Err:  true,
			Spec: `{"status_codes":[400,400]}`,
		},
	})
}
