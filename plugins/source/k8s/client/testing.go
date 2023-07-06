package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	v1 "k8s.io/api/core/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
)

type TestOption func(*Client)

func WithTestNamespaces(namespaces ...v1.Namespace) TestOption {
	return func(c *Client) {
		c.namespaces[c.Context] = namespaces
	}
}

func K8sMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) kubernetes.Interface, opts ...TestOption) {
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
	c.clients = map[string]kubernetes.Interface{"testContext": builder(t, mockController)}
	for _, opt := range opts {
		opt(c)
	}
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))
	messages, err := sched.SyncAll(context.Background(), c, schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	records := messages.GetInserts().GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}

func APIExtensionsMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) apiextensionsclientset.Interface, opts ...TestOption) {
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
	c.apiExtensions = map[string]apiextensionsclientset.Interface{"testContext": builder(t, mockController)}
	for _, opt := range opts {
		opt(c)
	}
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))
	messages, err := sched.SyncAll(context.Background(), c, schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	records := messages.GetInserts().GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}
