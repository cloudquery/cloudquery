package websecurityscanner

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"

	pb "cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
)

func createServer(gsrv *grpc.Server) error {
	fakeServer := &fakeServer{}
	pb.RegisterWebSecurityScannerServer(gsrv, fakeServer)
	return nil
}

type fakeServer struct {
	pb.UnimplementedWebSecurityScannerServer
}

func (*fakeServer) ListScanConfigs(context.Context, *pb.ListScanConfigsRequest) (*pb.ListScanConfigsResponse, error) {
	resp := pb.ListScanConfigsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeServer) ListFindings(context.Context, *pb.ListFindingsRequest) (*pb.ListFindingsResponse, error) {
	resp := pb.ListFindingsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeServer) ListScanRuns(context.Context, *pb.ListScanRunsRequest) (*pb.ListScanRunsResponse, error) {
	resp := pb.ListScanRunsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeServer) ListCrawledUrls(context.Context, *pb.ListCrawledUrlsRequest) (*pb.ListCrawledUrlsResponse, error) {
	resp := pb.ListCrawledUrlsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeServer) ListFindingTypeStats(context.Context, *pb.ListFindingTypeStatsRequest) (*pb.ListFindingTypeStatsResponse, error) {
	resp := pb.ListFindingTypeStatsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func TestInstances(t *testing.T) {
	client.MockTestGrpcHelper(t, ScanConfigs(), createServer, client.TestOptions{})
}
