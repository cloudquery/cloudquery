package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	TableOptions tableoptions.TableOptions
	Region       string
}

func AwsMockTestHelper(t *testing.T, parentTable *schema.Table, builder func(*testing.T, *gomock.Controller) Services, testOpts TestOptions) {
	parentTable.IgnoreInTests = false
	if testOpts.Region == "" {
		testOpts.Region = "us-east-1"
	}

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
	services.Regions = []string{testOpts.Region}
	c.ServicesManager.InitServicesForPartitionAccount("aws", "testAccount", services)
	c.Partition = "aws"
	tables := schema.Tables{parentTable}

	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}

	sc := scheduler.NewScheduler(scheduler.WithLogger(l))
	messages, err := sc.SyncAll(context.Background(), &c, tables)
	if err != nil {
		t.Fatal(err)
	}
	for _, table := range tables.FlattenTables() {
		records := messages.GetInserts().GetRecordsForTable(table)
		emptyColumns := schema.FindEmptyColumns(table, records)
		if len(emptyColumns) > 0 {
			t.Fatalf("found empty column(s): %v in %s", emptyColumns, table.Name)
		}
	}
}
