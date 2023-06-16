package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	TableOptions tableoptions.TableOptions
}

func AwsMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, testOpts TestOptions) {
	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	var awsSpec Spec
	awsSpec.SetDefaults()
	awsSpec.UsePaidAPIs = true
	awsSpec.TableOptions = &testOpts.TableOptions
	c := NewAwsClient(l, &awsSpec)
	services := builder(t, ctrl)
	services.Regions = []string{"us-east-1"}
	c.ServicesManager.InitServicesForPartitionAccount("aws", "testAccount", services)
	c.Partition = "aws"

	ch := make(chan message.Message, 1000)
	if err := scheduler.NewScheduler(schema.Tables{table}, &c, scheduler.WithLogger(l)).Sync(context.Background(), ch); err != nil {
		t.Fatal(err)
	}
	// TODO: add tests
}
func extractTables(tables schema.Tables) schema.Tables {
	result := make(schema.Tables, 0)
	for _, table := range tables {
		result = append(result, table)
		result = append(result, extractTables(table.Relations)...)
	}
	return result
}

// func validateTagStructure(t *testing.T, plugin *source.Plugin, resources []*schema.Resource) {
// 	for _, table := range extractTables(plugin.Tables()) {
// 		t.Run(table.Name, func(t *testing.T) {
// 			for _, column := range table.Columns {
// 				if column.Name != "tags" {
// 					continue
// 				}
// 				if !arrow.TypeEqual(column.Type, types.ExtensionTypes.JSON) {
// 					t.Fatalf("tags column in %s should be of type JSON", table.Name)
// 				}
// 				for _, resource := range resources {
// 					if resource.Table.Name != table.Name {
// 						continue
// 					}
// 					value := resource.Get(column.Name)
// 					var tags map[string]any
// 					if err := json.Unmarshal(value.(*scalar.JSON).Value, &tags); err != nil {
// 						t.Fatalf("failed to unmarshal tags column %s: %v", value.(*scalar.JSON).Value, err)
// 					}
// 				}
// 			}
// 		})
// 	}
// }
