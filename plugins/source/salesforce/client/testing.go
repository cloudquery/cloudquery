package client

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func MockTestHelper(t *testing.T, table *schema.Table, createServices func(*mux.Router) error) {
	t.Helper()

	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	sched := scheduler.NewScheduler(scheduler.WithLogger(logger))

	table.IgnoreInTests = false
	router := mux.NewRouter()
	h := httptest.NewServer(router)
	defer h.Close()

	if err := createServices(router); err != nil {
		t.Fatal(err)
	}
	spec := Spec{
		Username:     "test",
		Password:     "test",
		ClientId:     "test",
		ClientSecret: "test",
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		t.Fatal(err)
	}
	cqClient := &Client{
		logger: logger,
		Client: h.Client(),
		LoginResponse: LoginResponse{
			AccessToken: "SomeTestToken",
			InstanceUrl: h.URL,
		},
		HTTPDataEndpoint: h.URL + "/services/data/" + ApiVersion,
		ListObjectsResponse: ListObjectsResponse{
			Sobject: []Sobject{
				{
					Name: "TestObject",
				},
			},
		},
		spec: spec,
	}

	messages, err := sched.SyncAll(context.Background(), cqClient, schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	records := messages.GetInserts().GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}
