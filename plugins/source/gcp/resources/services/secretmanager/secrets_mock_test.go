package secretmanager

import (
	"context"
	"fmt"
	"testing"

	pb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
)

func createSecrets(gsrv *grpc.Server) error {
	fakeServer := &fakeSecretsServer{}
	pb.RegisterSecretManagerServiceServer(gsrv, fakeServer)
	return nil
}

type fakeSecretsServer struct {
	pb.UnimplementedSecretManagerServiceServer
}

func (*fakeSecretsServer) ListSecrets(context.Context, *pb.ListSecretsRequest) (*pb.ListSecretsResponse, error) {
	resp := pb.ListSecretsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.Secrets[0].Expiration = &pb.Secret_Ttl{Ttl: &durationpb.Duration{Seconds: 1}}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestSecrets(t *testing.T) {
	client.MockTestHelper(t, Secrets(), client.WithCreateGrpcService(createSecrets))
}
