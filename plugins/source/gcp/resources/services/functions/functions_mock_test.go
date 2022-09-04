// Code generated by codegen; DO NOT EDIT.

package functions

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"

	"cloud.google.com/go/functions/apiv1"

	pb "google.golang.org/genproto/googleapis/cloud/functions/v1"

	"google.golang.org/api/option"
)

func createFunctions() (*client.Services, error) {
	fakeServer := &fakeFunctionsServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterCloudFunctionsServiceServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	// Create a client.
	svc, err := functions.NewCloudFunctionsClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &client.Services{
		FunctionsCloudFunctionsClient: svc,
	}, nil
}

type fakeFunctionsServer struct {
	pb.UnimplementedCloudFunctionsServiceServer
}

func (f *fakeFunctionsServer) ListFunctions(context.Context, *pb.ListFunctionsRequest) (*pb.ListFunctionsResponse, error) {
	resp := pb.ListFunctionsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestFunctions(t *testing.T) {
	client.MockTestHelper(t, Functions(), createFunctions, client.TestOptions{})
}
