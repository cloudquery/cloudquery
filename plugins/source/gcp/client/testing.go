package client

import (
	"context"
	"fmt"
	"net"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestGrpcHelper(t *testing.T, table *schema.Table, createService func(*grpc.Server) error, options TestOptions) {
	version := "vDev"
	t.Helper()

	table.IgnoreInTests = false
	gsrv := grpc.NewServer()
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer gsrv.Stop()
	eg := &errgroup.Group{}
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		err := createService(gsrv)
		if err != nil {
			return nil, fmt.Errorf("failed to createService: %w", err)
		}
		eg.Go(func() error {
			return gsrv.Serve(listener)
		})
		var gcpSpec Spec
		if err := spec.UnmarshalSpec(&gcpSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
		}
		clientOptions := []option.ClientOption{
			option.WithEndpoint(listener.Addr().String()),
			option.WithoutAuthentication(),
			option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
		}
		c := &Client{
			logger:        logger,
			ClientOptions: clientOptions,
			projects:      []string{"testProject"},
			orgs:          []string{"testOrg"},
		}

		return c, nil
	}
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	p.SetLogger(l)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
	gsrv.Stop()
	if err := eg.Wait(); err != nil {
		t.Fatalf("failed to serve: %v", err)
	}
}

func MockTestRestHelper(t *testing.T, table *schema.Table, createService func(*httprouter.Router) error, options TestOptions) {
	version := "vDev"
	t.Helper()

	table.IgnoreInTests = false
	mux := httprouter.New()
	ts := httptest.NewUnstartedServer(mux)
	defer ts.Close()
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		err := createService(mux)
		if err != nil {
			return nil, fmt.Errorf("failed to createService: %w", err)
		}
		ts.Start()
		var gcpSpec Spec
		if err := spec.UnmarshalSpec(&gcpSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
		}
		clientOptions := []option.ClientOption{
			option.WithEndpoint(ts.URL),
			option.WithoutAuthentication(),
		}
		c := &Client{
			logger:        logger,
			ClientOptions: clientOptions,
			projects:      []string{"testProject"},
			orgs:          []string{"testOrg"},
		}

		return c, nil
	}
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	p.SetLogger(l)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
