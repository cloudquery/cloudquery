package client

import (
	"context"
	"net"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestGrpcHelper(t *testing.T, table *schema.Table, createService func(*grpc.Server) error, options TestOptions) {
	t.Helper()

	table.IgnoreInTests = false
	gsrv := grpc.NewServer()
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer gsrv.Stop()
	eg := &errgroup.Group{}
	if err := createService(gsrv); err != nil {
		t.Fatal(err)
	}
	eg.Go(func() error {
		return gsrv.Serve(listener)
	})
	clientOptions := []option.ClientOption{
		option.WithEndpoint(listener.Addr().String()),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	}
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	c := &Client{
		logger:        l,
		ClientOptions: clientOptions,
		projects:      []string{"testProject"},
		orgs:          []*crmv1.Organization{{Name: "organizations/testOrg"}},
		folderIds:     []string{"testFolder"},
		// Backend:       opts.Backend,
	}


	sched := scheduler.NewScheduler(c, scheduler.WithLogger(l))
	messages, err := sched.SyncAll(context.Background(), schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}

	records := filterInserts(messages).GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
	gsrv.Stop()
	if err := eg.Wait(); err != nil {
		t.Fatalf("failed to serve: %v", err)
	}
}

func filterInserts(msgs message.Messages) message.Inserts {
	inserts := []*message.Insert{}
	for _, msg := range msgs {
		if m, ok := msg.(*message.Insert); ok {
			inserts = append(inserts, m)
		}
	}
	return inserts
}

func MockTestRestHelper(t *testing.T, table *schema.Table, createService func(*httprouter.Router) error, options TestOptions) {
	t.Helper()

	table.IgnoreInTests = false
	mux := httprouter.New()
	ts := httptest.NewUnstartedServer(mux)
	defer ts.Close()
	if err := createService(mux); err != nil {
		t.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ts.Start()
	}()
	time.Sleep(1 * time.Second)
	clientOptions := []option.ClientOption{
		option.WithEndpoint(ts.URL),
		option.WithoutAuthentication(),
	}
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	c := &Client{
		logger:        l,
		ClientOptions: clientOptions,
		projects:      []string{"testProject"},
		orgs:          []*crmv1.Organization{{Name: "organizations/testOrg"}},
	}

	sched := scheduler.NewScheduler(c, scheduler.WithLogger(l))
	messages, err := sched.SyncAll(context.Background(), schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}

	records := filterInserts(messages).GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
	ts.Close()
}
