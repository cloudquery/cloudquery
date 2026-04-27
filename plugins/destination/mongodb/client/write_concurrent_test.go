package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/sync/errgroup"
)

// TestConcurrentWrites stresses the MongoDB destination write path from many
// goroutines in parallel against a healthy local mongo. It catches concurrency
// regressions in the write/transform code; it does *not* exercise the retry
// path -- see TestRetryAbsorbsConnectionDrop in retry_integration_test.go for
// a deterministic regression test for ENG-3281.
func TestConcurrentWrites(t *testing.T) {
	const (
		workers    = 32
		iterations = 5
		batchSize  = 200
	)

	ctx := context.Background()
	s := &spec.Spec{
		ConnectionString: getTestConnection(),
		Database:         "destination_mongodb_concurrent_test",
	}
	specBytes, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	logger := zerolog.New(zerolog.NewTestWriter(t))
	pc, err := New(ctx, logger, specBytes, plugin.NewClientOptions{})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	c := pc.(*Client)
	t.Cleanup(func() {
		if err := c.client.Database(s.Database).Drop(ctx); err != nil {
			t.Logf("failed to drop test database: %v", err)
		}
		if err := pc.Close(ctx); err != nil {
			t.Logf("failed to close plugin client: %v", err)
		}
	})

	appendTable := &schema.Table{
		Name: "concurrent_append",
		Columns: schema.ColumnList{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "val", Type: arrow.BinaryTypes.String},
		},
	}
	overwriteTable := &schema.Table{
		Name: "concurrent_overwrite",
		Columns: schema.ColumnList{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
			{Name: "val", Type: arrow.BinaryTypes.String},
		},
	}

	t.Logf("running %d workers x %d iterations x %d docs/batch against %s", workers, iterations, batchSize, s.ConnectionString)

	g, gctx := errgroup.WithContext(ctx)
	for w := 0; w < workers; w++ {
		g.Go(func() error {
			for i := 0; i < iterations; i++ {
				docs := make([]any, batchSize)
				for k := 0; k < batchSize; k++ {
					id := int64(w*iterations*batchSize + i*batchSize + k)
					docs[k] = bson.M{
						"id":  id,
						"val": fmt.Sprintf("w%d-i%d-k%d", w, i, k),
					}
				}
				if err := c.appendTableBatch(gctx, appendTable, docs); err != nil {
					return fmt.Errorf("worker %d iter %d append: %w", w, i, err)
				}
				if err := c.overwriteTableBatch(gctx, overwriteTable, docs); err != nil {
					return fmt.Errorf("worker %d iter %d overwrite: %w", w, i, err)
				}
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}
}
