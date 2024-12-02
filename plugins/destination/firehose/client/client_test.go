package client

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/firehose/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

func TestPluginJSON(t *testing.T) {
	t.Skip("skip per invalid ARN error")
	ctx := context.Background()
	p := plugin.NewPlugin("firehose", "development", New)
	const streamARN = "cq-playground-test"
	s := &spec.Spec{StreamARN: streamARN}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	require.NoError(t, p.Init(ctx, b, plugin.NewClientOptions{}))

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipUpsert:       true,
			SkipMigrate:      true,
			SkipDeleteStale:  true,
			SkipDeleteRecord: true,
		},
	)
}
