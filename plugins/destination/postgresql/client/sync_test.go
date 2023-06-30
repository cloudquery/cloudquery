package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"golang.org/x/sync/errgroup"
)

func TestSync(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("postgresql", "development", New)
	s := &Spec{
		ConnectionString: getTestConnection(),
		PgxLogLevel:      LogLevelWarn,
	}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	err = p.Init(ctx, b)
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	tableName := fmt.Sprintf("cq_test_sync_%d", time.Now().Unix())
	tbl := schema.TestTable(tableName, schema.TestSourceOptions{
		SkipMaps: true,
	})
	records := schema.GenTestData(tbl, schema.GenTestDataOptions{
		SourceName:    "test-source",
		SyncTime:      now,
		MaxRows:       2,
		StableTime:    now,
		TimePrecision: 0,
		Seed:          0,
	})
	err = p.WriteAll(ctx, []message.WriteMessage{
		&message.WriteMigrateTable{
			Table:        tbl,
			MigrateForce: false,
		},
		&message.WriteInsert{
			Record: records[0],
		},
		&message.WriteInsert{
			Record: records[1],
		},
	})
	if err != nil {
		t.Fatal("failed to write test messages:", err)
	}
	opts := plugin.SyncOptions{
		Tables:              []string{tableName},
		SkipTables:          []string{},
		SkipDependentTables: false,
		DeterministicCQID:   false,
	}
	res := make(chan message.SyncMessage)
	g := errgroup.Group{}
	g.Go(func() error {
		defer close(res)
		return p.Sync(ctx, opts, res)
	})
	gotInserts := make([]*message.SyncInsert, 0, 2)
	for msg := range res {
		switch v := msg.(type) {
		case *message.SyncInsert:
			gotInserts = append(gotInserts, v)
		}
	}
	err = g.Wait()
	if err != nil {
		t.Fatal(err)
	}
	if len(gotInserts) != 2 {
		t.Fatalf("expected 2 inserts, got %d", len(gotInserts))
	}
	// we don't expect the records to match exactly because some columns would have been
	// converted to strings.
}
