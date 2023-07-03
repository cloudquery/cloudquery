package kms

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func createLocations(gsrv *grpc.Server) error {
	fakeServer := &fakeServer{}
	kmspb.RegisterKeyManagementServiceServer(gsrv, fakeServer)
	kmspb.RegisterEkmServiceServer(gsrv, fakeServer)
	locationpb.RegisterLocationsServer(gsrv, fakeServer)

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

type fakeServer struct {
	kmspb.UnimplementedKeyManagementServiceServer
	kmspb.UnimplementedEkmServiceServer
	locationpb.UnimplementedLocationsServer
}

func (*fakeServer) ListKeyRings(context.Context, *kmspb.ListKeyRingsRequest) (*kmspb.ListKeyRingsResponse, error) {
	resp := kmspb.ListKeyRingsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeServer) ListCryptoKeys(context.Context, *kmspb.ListCryptoKeysRequest) (*kmspb.ListCryptoKeysResponse, error) {
	resp := kmspb.ListCryptoKeysResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeServer) ListCryptoKeyVersions(context.Context, *kmspb.ListCryptoKeyVersionsRequest) (*kmspb.ListCryptoKeyVersionsResponse, error) {
	resp := kmspb.ListCryptoKeyVersionsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeServer) ListLocations(context.Context, *locationpb.ListLocationsRequest) (*locationpb.ListLocationsResponse, error) {
	resp := locationpb.ListLocationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeServer) ListImportJobs(context.Context, *kmspb.ListImportJobsRequest) (*kmspb.ListImportJobsResponse, error) {
	resp := kmspb.ListImportJobsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeServer) ListEkmConnections(context.Context, *kmspb.ListEkmConnectionsRequest) (*kmspb.ListEkmConnectionsResponse, error) {
	resp := kmspb.ListEkmConnectionsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestLocations(t *testing.T) {
	client.MockTestGrpcHelper(t, Locations(), createLocations, client.TestOptions{})
}
