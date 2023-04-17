package client

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"github.com/rs/zerolog"
	"github.com/shenzhencenter/google-ads-pb/clients"
	"github.com/shenzhencenter/google-ads-pb/services"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MockedResponses map[string][]*services.GoogleAdsRow

func MockTestHelper(t *testing.T, table *schema.Table, responses MockedResponses) {
	version := "vDev"
	t.Helper()

	table.IgnoreInTests = false
	gsrv := grpc.NewServer()
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer gsrv.Stop()
	m := new(sync.Map)
	for k, v := range responses {
		m.Store(k, v)
	}
	fake := &fakeGoogleAdsServiceServer{responses: m}
	services.RegisterGoogleAdsServiceServer(gsrv, fake)
	eg := &errgroup.Group{}
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, srcSpec specs.Source, opts source.Options) (schema.ClientMeta, error) {
		eg.Go(func() error {
			return gsrv.Serve(listener)
		})
		var spec Spec
		if err := srcSpec.UnmarshalSpec(&spec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal Google Ads spec: %w", err)
		}
		clientOptions := []option.ClientOption{
			option.WithEndpoint(listener.Addr().String()),
			option.WithoutAuthentication(),
			option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
		}

		client, err := clients.NewGoogleAdsClient(ctx, clientOptions...)
		if err != nil {
			return nil, fmt.Errorf("failed to create GoogleAdsClient: %w", err)
		}

		c := &Client{
			GoogleAdsClient: client,
			CustomerID:      "12345",
			ManagerID:       "67890",
			customers: map[string][]string{
				"67890": {"12345"},
			},
			developerToken: "testDeveloperToken",
			logger:         logger,
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
	fake.responses.Range(func(key, _ any) bool {
		t.Fatalf("unused response for %s", key)
		return true
	})
}

type fakeGoogleAdsServiceServer struct {
	*services.UnimplementedGoogleAdsServiceServer
	responses *sync.Map
}

func (s *fakeGoogleAdsServiceServer) SearchStream(req *services.SearchGoogleAdsStreamRequest, srv services.GoogleAdsService_SearchStreamServer) error {
	table := strings.Split(strings.Split(req.Query, "FROM ")[1], "\n")[0]

	rows, ok := s.responses.LoadAndDelete(table)
	if !ok {
		return fmt.Errorf("not expected to query %q", table)
	}

	return srv.Send(&services.SearchGoogleAdsStreamResponse{Results: rows.([]*services.GoogleAdsRow)})
}
