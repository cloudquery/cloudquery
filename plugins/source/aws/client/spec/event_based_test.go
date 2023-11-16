package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestEventBasedSyncJSONSchema(t *testing.T) {
	data, err := jsonschema.Generate(EventBasedSync{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true, // missing kinesis_stream_arn
			Spec: `{}`,
		},
		{
			Name: "empty kinesis_stream_arn",
			Err:  true,
			Spec: `{"kinesis_stream_arn":""}`,
		},
		{
			Name: "null kinesis_stream_arn",
			Err:  true,
			Spec: `{"kinesis_stream_arn":null}`,
		},
		{
			Name: "bad kinesis_stream_arn",
			Err:  true,
			Spec: `{"kinesis_stream_arn":123}`,
		},
		{
			Name: "bad kinesis_stream_arn format",
			Err:  true,
			Spec: `{"kinesis_stream_arn":"abc"}`,
		},
		{
			Name: "proper kinesis_stream_arn",
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6"}`,
		},
		// account is tested separately, here we test basic cases
		{
			Name: "empty account",
			Err:  true, // missing account.id
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "account":{}}`,
		},
		{
			Name: "null account",
			Err:  true,
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "account":null}`,
		},
		{
			Name: "bad account",
			Err:  true,
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "account":123}`,
		},
		{
			Name: "null full_sync",
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "full_sync":null}`,
		},
		{
			Name: "bad full_sync",
			Err:  true,
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "full_sync":123}`,
		},
		{
			Name: "full_sync:true",
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "full_sync":true}`,
		},
		{
			Name: "full_sync:false",
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "full_sync":false}`,
		},
		{
			Name: "empty start_time",
			Err:  true,
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "start_time":""}`,
		},
		{
			Name: "null start_time",
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "start_time":null}`,
		},
		{
			Name: "bad start_time",
			Err:  true,
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "start_time":123}`,
		},
		{
			Name: "bad start_time format",
			Err:  true,
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "start_time":"abc"}`,
		},
		{
			Name: "proper start_time format",
			Spec: `{"kinesis_stream_arn":"arn:1:2:3:4:5/6", "start_time":"2006-01-02T15:04:05+07:00"}`,
		},
	})
}
