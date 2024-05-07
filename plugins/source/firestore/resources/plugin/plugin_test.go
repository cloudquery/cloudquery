package plugin

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/cloudquery/cloudquery/plugins/source/firestore/client"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestPlugin(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)

	tableSetup(ctx, t, "test_firestore_source")

	spec := client.Spec{
		ProjectID: "cqtest-project",
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, p.Init(ctx, specBytes, plugin.NewClientOptions{}))
}

func TestPlugin_OrderBy(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)

	tableSetup(ctx, t, "test_firestore_source")

	spec := client.Spec{
		ProjectID: "cqtest-project",
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, p.Init(ctx, specBytes, plugin.NewClientOptions{}))
}

func assertTimestamp(t *testing.T, expected time.Time, ts *array.Timestamp) {
	timeUnit := ts.DataType().(*arrow.TimestampType).Unit
	require.Equal(t, expected, ts.Value(0).ToTime(timeUnit))
}

func TestPluginSync(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)

	expectedData := tableSetup(ctx, t, "test_firestore_source")

	spec := client.Spec{
		ProjectID: "cqtest-project",
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, p.Init(ctx, specBytes, plugin.NewClientOptions{}))

	res := make(chan message.SyncMessage, 1)
	g := errgroup.Group{}
	g.Go(func() error {
		defer close(res)
		opts := plugin.SyncOptions{
			Tables: []string{"test_firestore_source"}}
		return p.Sync(ctx, opts, res)
	})
	actualRecords := make([]arrow.Record, 0)
	for r := range res {
		m, ok := r.(*message.SyncInsert)
		if ok {
			actualRecords = append(actualRecords, m.Record)
		}
	}
	err = g.Wait()
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	if len(actualRecords) != len(expectedData) {
		t.Fatalf("expected 1 resource, got %d", len(actualRecords))
	}

	sort.Slice(actualRecords, func(i, j int) bool {
		return actualRecords[i].Column(0).ValueStr(0) < actualRecords[j].Column(0).ValueStr(0)
	})
	sort.Slice(expectedData, func(i, j int) bool {
		return expectedData[i].ID < expectedData[j].ID
	})

	for i, actual := range actualRecords {
		snap, err := expectedData[i].Get(ctx)
		if err != nil {
			t.Fatal(err)
		}
		require.Equal(t, snap.Ref.ID, actual.Column(0).ValueStr(0))

		assertTimestamp(t, snap.CreateTime, actual.Column(1).(*array.Timestamp))
		assertTimestamp(t, snap.UpdateTime, actual.Column(2).(*array.Timestamp))

		expectedData, err := json.Marshal(snap.Data())
		require.NoError(t, err)
		require.Equal(t, string(expectedData), actual.Column(3).(*types.JSONArray).ValueStr(0))
	}
}

func tableSetup(ctx context.Context, t *testing.T, tableName string) []*firestore.DocumentRef {
	// insert test data
	store, err := firestore.NewClient(ctx, "cqtest-project")
	require.NoError(t, err)
	col := store.Collection(tableName)

	// delete all collection data
	docs, err := col.Documents(ctx).GetAll()
	require.NoError(t, err)
	for _, d := range docs {
		_, err := d.Ref.Delete(ctx)
		require.NoError(t, err)
	}
	// insert new data
	data := testData()
	refs := make([]*firestore.DocumentRef, 0)
	for _, d := range data {
		ref, _, err := col.Add(ctx, d)
		require.NoError(t, err)
		refs = append(refs, ref)
	}
	return refs
}

func testData() []map[string]any {
	randData := make([]map[string]any, 0)
	for i := 0; i < 100; i++ {
		randData = append(randData, map[string]any{
			"key1": i,
			"key":  rand.Int(),
		})
	}
	return randData
}
