package client

import (
	"context"
	"net"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TestOptions struct {
	SkipEmptyJsonB    bool
	CreateGrpcService func(*grpc.Server) error
	CreateHTTPServer  func(*httprouter.Router) error
}

type Option func(*TestOptions)

func WithCreateGrpcService(createGrpcService func(*grpc.Server) error) Option {
	return func(to *TestOptions) {
		to.CreateGrpcService = createGrpcService
	}
}

func WithCreateHTTPServer(createHTTPServer func(*httprouter.Router) error) Option {
	return func(to *TestOptions) {
		to.CreateHTTPServer = createHTTPServer
	}
}

func MockTestHelper(t *testing.T, table *schema.Table, opts ...Option) {
	t.Helper()

	options := TestOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	table.IgnoreInTests = false
	clientOptions := []option.ClientOption{
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	}
	var gsrv *grpc.Server
	var eg *errgroup.Group
	var grpcEndpoint string
	if options.CreateGrpcService != nil {
		gsrv = grpc.NewServer()
		listener, err := net.Listen("tcp", "localhost:0")
		if err != nil {
			t.Fatalf("failed to listen: %v", err)
		}
		defer gsrv.Stop()
		eg = &errgroup.Group{}
		if err := options.CreateGrpcService(gsrv); err != nil {
			t.Fatal(err)
		}
		eg.Go(func() error {
			return gsrv.Serve(listener)
		})
		grpcEndpoint = listener.Addr().String()
		clientOptions = append(clientOptions, option.WithEndpoint(grpcEndpoint))
	}

	var mux *httprouter.Router
	var ts *httptest.Server
	var wg *sync.WaitGroup
	var httpTestEndpoint string
	if options.CreateHTTPServer != nil {
		mux = httprouter.New()
		ts = httptest.NewUnstartedServer(mux)
		httpTestEndpoint = "http://" + ts.Listener.Addr().String()
		defer ts.Close()
		if err := options.CreateHTTPServer(mux); err != nil {
			t.Fatal(err)
		}
		wg = &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			ts.Start()
		}()
		clientOptions = append(clientOptions,
			option.WithEndpoint(httpTestEndpoint),
			option.WithGRPCDialOption(grpc.WithBlock()),
		)
	}
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	c := &Client{
		Backend:             &state.NoOpClient{},
		ClientOptions:       clientOptions,
		folderIds:           []string{"testFolder"},
		logger:              l,
		orgs:                []*crmv1.Organization{{Name: "organizations/testOrg"}},
		projects:            []string{"testProject"},
		TestingGRPCEndpoint: grpcEndpoint,
		TestingHTTPEndpoint: httpTestEndpoint,
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
	if gsrv != nil {
		gsrv.Stop()
		if err := eg.Wait(); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}
	if mux != nil {
		ts.Close()
		wg.Wait()
	}
}
