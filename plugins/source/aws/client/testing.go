package client

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct{}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, _ TestOptions) {
	t.Helper()
	table.IgnoreInTests = false

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newTestExecutionClient := func(ctx context.Context, _ zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
		var awsSpec Spec
		if err := spec.UnmarshalSpec(&awsSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal aws spec: %w", err)
		}
		logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
			zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
		).Level(zerolog.DebugLevel).With().Timestamp().Logger()
		c := NewAwsClient(logger)
		c.AccountID = "testAccount"
		c.Region = "us-east-1"
		c.Partition = "aws"
		c.maxRetries = 3
		c.maxBackoff = 60

		c.ServicesManager.InitServicesForPartitionAccountAndRegion(c.Partition, c.AccountID, c.Region, builder(t, ctrl))

		return &c, nil
	}

	p := plugins.NewSourcePlugin(
		table.Name,
		"dev",
		[]*schema.Table{
			table,
		},
		newTestExecutionClient,
	)
	plugins.TestSourcePluginSync(t, p, specs.Source{
		Name:   "dev",
		Tables: []string{table.Name},
	})
}
