package client

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
	"github.com/stripe/stripe-go/v74"
	sclient "github.com/stripe/stripe-go/v74/client"
	"github.com/stripe/stripe-mock/server"
)

type TestOptions struct {
	Backend state.Client
}

func MockTestHelper(t *testing.T, table *schema.Table, opts TestOptions) {
	t.Helper()
	table.IgnoreInTests = false

	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	addr, teardown, err := startMockServer()
	if err != nil {
		t.Fatalf("startMockServer: %v", err)
	}
	defer func() {
		if err := teardown(); err != nil {
			t.Logf("Teardown error: %v", err)
		}
	}()

	sched := scheduler.NewScheduler(scheduler.WithLogger(logger))
	spec := &Spec{
		APIKey: "sk_test_myTestKey",
	}
	if err := spec.Validate(); err != nil {
		t.Fatalf("failed to validate spec: %v", err)
	}
	spec.SetDefaults()

	cl := sclient.New(spec.APIKey, getBackends(logger, *addr))

	c := New(logger, *spec, cl, nil)

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

func startMockServer() (*string, func() error, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, nil, fmt.Errorf("failed to get current file")
	}

	stripeSpec, err := server.LoadSpec(nil, filepath.Join(filepath.Dir(filename), "..", "resources", "testdata", "spec3.json"))
	if err != nil {
		return nil, nil, err
	}
	fixtures, err := server.LoadFixtures(nil, filepath.Join(filepath.Dir(filename), "..", "resources", "testdata", "fixtures_gen.json"))
	if err != nil {
		return nil, nil, err
	}

	// This is mostly copied from the stripe-mock's main function at https://github.com/stripe/stripe-mock/blob/master/main.go
	stub, err := server.NewStubServer(fixtures, stripeSpec, false, false)
	if err != nil {
		return nil, nil, fmt.Errorf("Error initializing router: %w", err)
	}

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", stub.HandleRequest)

	// Deduplicates doubled slashes in paths. e.g. `//v1/charges` becomes `/v1/charges`.
	handler := &server.DoubleSlashFixHandler{Mux: httpMux}

	listener, err := net.Listen("tcp", "localhost:0") // auto-choose port
	if err != nil {
		return nil, nil, fmt.Errorf("net.Listen: %w", err)
	}

	fmt.Printf("Listening for TCP at address: %v\n", listener.Addr())

	srv := &http.Server{
		Handler: handler,
	}
	go func() {
		err := srv.Serve(listener)
		if err != nil && err != http.ErrServerClosed {
			panic(err.Error())
		}
	}()

	addr := listener.Addr().String()
	return &addr, func() error {
		return srv.Shutdown(context.Background())
	}, nil
}

func getBackends(logger zerolog.Logger, addr string) *stripe.Backends {
	// This is mostly copied from stripe-go's testing package at https://github.com/stripe/stripe-go/blob/master/testing/testing.go
	// Since the code is inside init() it can't be used directly without multiple headaches, and even if we do that we can't force it to run HTTP only.
	stripeMockBackend := stripe.GetBackendWithConfig(
		stripe.APIBackend,
		&stripe.BackendConfig{
			URL:           stripe.String("http://" + addr),
			LeveledLogger: &LeveledLogger{Logger: logger},
		},
	)
	backends := stripe.NewBackends(nil)
	backends.API = stripeMockBackend
	backends.Uploads = stripeMockBackend
	return backends
}
