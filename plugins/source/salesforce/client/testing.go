package client

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
	c := &Client{
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
