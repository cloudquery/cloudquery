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

	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/stripe/stripe-go/v74"
	sclient "github.com/stripe/stripe-go/v74/client"
	"github.com/stripe/stripe-mock/server"
)

type TestOptions struct {
	Backend backend.Backend
}

func MockTestHelper(t *testing.T, table *schema.Table, opts TestOptions) {
	version := "vDev"

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

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ ...source.Option) (schema.ClientMeta, error) {
		var stSpec Spec
		if err := spec.UnmarshalSpec(&stSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal stripe spec: %w", err)
		}

		cl := sclient.New("sk_test_myTestKey", getBackends(logger, *addr))
		return New(logger, spec, stSpec, cl, opts.Backend), nil
	}

	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient,
	)
	p.SetLogger(logger)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
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
