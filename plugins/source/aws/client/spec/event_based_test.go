package spec

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/require"
)

func TestEventBasedSyncJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true,
			Spec: `{"event_based_sync":{}}`,
		},
		{
			Name: "null",
			Spec: `{"event_based_sync":null}`,
		},
		{
			Name: "bad",
			Err:  true,
			Spec: `{"event_based_sync":123}`,
		},
		{
			Name: "proper",
			Spec: func() string {
				var input EventBasedSync
				require.NoError(t, faker.FakeObject(&input))
				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.KinesisStreamARN = randomARN.String()
				input.Account.RoleARN = randomARN.String()
				return `{"event_based_sync":` + jsonschema.WithRemovedKeys(t, &input) + `}`
			}(),
		},
		{
			Name: "bad kinesis_stream_arn",
			Err:  true,
			Spec: func() string {
				var input EventBasedSync
				require.NoError(t, faker.FakeObject(&input))
				var randomARN arn.ARN
				require.NoError(t, faker.FakeObject(&randomARN))
				input.Account.RoleARN = randomARN.String()
				return `{"event_based_sync":` + jsonschema.WithRemovedKeys(t, &input) + `}`
			}(),
		},
		{
			Name: "missing kinesis_stream_arn",
			Err:  true,
			Spec: func() string {
				var input EventBasedSync
				require.NoError(t, faker.FakeObject(&input))
				return `{"event_based_sync":` + jsonschema.WithRemovedKeys(t, &input, "kinesis_stream_arn") + `}`
			}(),
		},
	})
}
