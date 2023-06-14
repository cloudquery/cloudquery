package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/scalar"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	TableOptions tableoptions.TableOptions
}

func AwsMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, testOpts TestOptions) {
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
		awsSpec.SetDefaults()
		awsSpec.UsePaidAPIs = true
		awsSpec.TableOptions = &testOpts.TableOptions
		c := NewAwsClient(l, nil, &awsSpec)
		services := builder(t, ctrl)
		services.Regions = []string{"us-east-1"}
		c.ServicesManager.InitServicesForPartitionAccount("aws", "testAccount", services)
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
	}, source.WithTestPluginAdditionalValidators(validateTagStructure))
}
func extractTables(tables schema.Tables) schema.Tables {
	result := make(schema.Tables, 0)
	for _, table := range tables {
		result = append(result, table)
		result = append(result, extractTables(table.Relations)...)
	}
	return result
}

func validateTagStructure(t *testing.T, plugin *source.Plugin, resources []*schema.Resource) {
	for _, table := range extractTables(plugin.Tables()) {
		t.Run(table.Name, func(t *testing.T) {
			for _, column := range table.Columns {
				if column.Name != "tags" {
					continue
				}
				if !arrow.TypeEqual(column.Type, types.ExtensionTypes.JSON) {
					t.Fatalf("tags column in %s should be of type JSON", table.Name)
				}
				for _, resource := range resources {
					if resource.Table.Name != table.Name {
						continue
					}
					value := resource.Get(column.Name)
					var tags map[string]any
					if err := json.Unmarshal(value.(*scalar.JSON).Value, &tags); err != nil {
						t.Fatalf("failed to unmarshal tags column %s: %v", value.(*scalar.JSON).Value, err)
					}
				}
			}
		})
	}
}
