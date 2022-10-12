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

func AwsMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, _ TestOptions) {
	version := "vDev"

	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
		var awsSpec Spec
		if err := spec.UnmarshalSpec(&awsSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal aws spec: %w", err)
		}
		c := NewAwsClient(l)
		c.ServicesManager.InitServicesForPartitionAccountAndRegion("aws", "testAccount", "us-east-1", builder(t, ctrl))
		c.Partition = "aws"
		return &c, nil
	}

	p := plugins.NewSourcePlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	plugins.TestSourcePluginSync(t, p, l, specs.Source{
		Name:         "dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
