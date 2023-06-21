package plugin_test

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/cloudquery/cloudquery/plugins/source/firestore/client"
	"github.com/cloudquery/cloudquery/plugins/source/firestore/resources/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestPlugin(t *testing.T) {
	p := plugin.Plugin()
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
	require.NoError(t, p.Init(ctx, specBytes))
}

func TestPlugin_OrderBy(t *testing.T) {
	p := plugin.Plugin()
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
	require.NoError(t, p.Init(ctx, specBytes))
}

func tableSetup(ctx context.Context, t *testing.T, tableName string) {
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
	for _, d := range testData() {
		_, _, err := col.Add(ctx, d)
		require.NoError(t, err)
	}
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
