package kms

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func createKeyrings(gsrv *grpc.Server) error {
	fakeServer := &fakeKeyringsServer{}
	kmspb.RegisterKeyManagementServiceServer(gsrv, fakeServer)
	locationpb.RegisterLocationsServer(gsrv, new(fakeLocationsServer))

	location := &locationpb.Location{
		DisplayName: "testLocation",
		Name:        "projects/testProject/location/testLocation",
	}

	var keyring kmspb.KeyRing
	if err := faker.FakeObject(&keyring); err != nil {
		return err
	}
	keyring.Name = fmt.Sprintf("projects/testProject/location/%s/keyring/%s", location.Name, keyring.Name)
	keyring.CreateTime = timestamppb.Now()
	var key kmspb.CryptoKey
	if err := faker.FakeObject(&key); err != nil {
		return err
	}
	key.Name = fmt.Sprintf("%s/cryptokey/%s", keyring.Name, "test")
	key.CreateTime = timestamppb.Now()
	key.NextRotationTime = timestamppb.Now()
	key.Primary.CreateTime = timestamppb.Now()
	key.Primary.DestroyEventTime = timestamppb.Now()
	key.Primary.DestroyTime = timestamppb.Now()
	key.Primary.GenerateTime = timestamppb.Now()
	key.Primary.ImportTime = timestamppb.Now()

	return nil
}

type fakeKeyringsServer struct {
	kmspb.UnimplementedKeyManagementServiceServer
}

func (*fakeKeyringsServer) ListKeyRings(context.Context, *kmspb.ListKeyRingsRequest) (*kmspb.ListKeyRingsResponse, error) {
	resp := kmspb.ListKeyRingsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeKeyringsServer) ListCryptoKeys(context.Context, *kmspb.ListCryptoKeysRequest) (*kmspb.ListCryptoKeysResponse, error) {
	resp := kmspb.ListCryptoKeysResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

type fakeLocationsServer struct {
	location *locationpb.Location
	locationpb.UnimplementedLocationsServer
}

func (srv *fakeLocationsServer) ListLocations(context.Context, *locationpb.ListLocationsRequest) (*locationpb.ListLocationsResponse, error) {
	return &locationpb.ListLocationsResponse{Locations: []*locationpb.Location{srv.location}}, nil
}

func TestKeyrings(t *testing.T) {
	client.MockTestGrpcHelper(t, Keyrings(), createKeyrings, client.TestOptions{})
}
