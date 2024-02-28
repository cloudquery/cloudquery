package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestSpec_JSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "missing stream_arn",
			Spec: `{}`,
			Err:  true,
		},
		{
			Name: "empty stream_arn",
			Spec: `{"stream_arn": ""}`,
			Err:  true,
		},
		{
			Name: "null stream_arn",
			Spec: `{"stream_arn": null}`,
			Err:  true,
		},
		{
			Name: "int stream_arn",
			Spec: `{"stream_arn": 123}`,
			Err:  true,
		},
		{
			Name: "proper stream_arn",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name"}`,
		},
		{
			Name: "zero max_retries",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_retries": 0}`,
			Err:  true,
		},
		{
			Name: "negative max_retries",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_retries": -1}`,
			Err:  true,
		},
		{
			Name: "float max_retries",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_retries": 1.5}`,
			Err:  true,
		},
		{
			Name: "null max_retries",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_retries": null}`,
			Err:  true,
		},
		{
			Name: "string max_retries",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_retries": "123"}`,
			Err:  true,
		},
		{
			Name: "proper max_retries",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_retries": 123}`,
		},
		{
			Name: "zero max_record_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_record_size_bytes": 0}`,
			Err:  true,
		},
		{
			Name: "negative max_record_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_record_size_bytes": -1}`,
			Err:  true,
		},
		{
			Name: "float max_record_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_record_size_bytes": 1.5}`,
			Err:  true,
		},
		{
			Name: "null max_record_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_record_size_bytes": null}`,
			Err:  true,
		},
		{
			Name: "string max_record_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_record_size_bytes": "123"}`,
			Err:  true,
		},
		{
			Name: "proper max_record_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_record_size_bytes": 123}`,
		},
		{
			Name: "zero max_batch_records",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_records": 0}`,
			Err:  true,
		},
		{
			Name: "negative max_batch_records",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_records": -1}`,
			Err:  true,
		},
		{
			Name: "float max_batch_records",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_records": 1.5}`,
			Err:  true,
		},
		{
			Name: "null max_batch_records",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_records": null}`,
			Err:  true,
		},
		{
			Name: "string max_batch_records",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_records": "123"}`,
			Err:  true,
		},
		{
			Name: "proper max_batch_records",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_records": 123}`,
		},
		{
			Name: "zero max_batch_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_size_bytes": 0}`,
			Err:  true,
		},
		{
			Name: "negative max_batch_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_size_bytes": -1}`,
			Err:  true,
		},
		{
			Name: "float max_batch_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_size_bytes": 1.5}`,
			Err:  true,
		},
		{
			Name: "null max_batch_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_size_bytes": null}`,
			Err:  true,
		},
		{
			Name: "string max_batch_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_size_bytes": "123"}`,
			Err:  true,
		},
		{
			Name: "proper max_batch_size_bytes",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "max_batch_size_bytes": 123}`,
		},
		{
			Name: "extra key",
			Spec: `{"stream_arn": "arn:aws:firehose:us-east-1:01234:deliverystream/name", "extra": true}`,
			Err:  true,
		},
	})
}
