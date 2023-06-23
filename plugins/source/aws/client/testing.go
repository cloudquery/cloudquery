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
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
	tables := schema.Tables{table}

	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}

	sc := scheduler.NewScheduler(&c, scheduler.WithLogger(l))
	messages, err := sc.SyncAll(context.Background(), tables)
	if err != nil {
		t.Fatal(err)
	}

	records := filterInserts(messages).GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}

}

func filterInserts(msgs message.Messages) message.Inserts {
	inserts := []*message.Insert{}
	for _, msg := range msgs {
		if m, ok := msg.(*message.Insert); ok {
			inserts = append(inserts, m)
		}
	}
	return inserts
}
