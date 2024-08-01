package client

import (
	"context"
	"os"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOption func(*Client)

func WithTestNamespaces(namespaces ...v1.Namespace) TestOption {
	return func(c *Client) {
		c.namespaces[c.Context] = namespaces
	}
}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, opts ...TestOption) {
	t.Helper()
	table.IgnoreInTests = false

	mockController := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	c := &Client{
		logger:     l,
		Context:    "testContext",
		contexts:   []string{"testContext"},
		namespaces: map[string][]v1.Namespace{},
	}
	c.apis = map[contextName]Services{"testContext": builder(t, mockController)}
	for _, opt := range opts {
		opt(c)
	}
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))
	tables := schema.Tables{table}
	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}
	messages, err := sched.SyncAll(context.Background(), c, tables)
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	plugin.ValidateNoEmptyColumns(t, tables, messages)
}
