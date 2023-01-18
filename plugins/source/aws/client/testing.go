package client

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
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

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		var awsSpec Spec
		if err := spec.UnmarshalSpec(&awsSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal aws spec: %w", err)
		}
		c := NewAwsClient(l)
		c.ServicesManager.InitServicesForPartitionAccountAndRegion("aws", "testAccount", "us-east-1", builder(t, ctrl))
		c.Partition = "aws"
		return &c, nil
	}

	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	p.SetLogger(l)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
