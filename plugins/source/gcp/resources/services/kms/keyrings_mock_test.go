package kms

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	kms "cloud.google.com/go/kms/apiv1"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	kmsold "google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/kms/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createKeyrings() (*client.Services, error) {
	fakeServer := &fakeKeyringsServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterKeyManagementServiceServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	mux := httprouter.New()
	mux.GET("/v1/projects/testProject/locations", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &kmsold.ListLocationsResponse{
			Locations: []*kmsold.Location{{
				DisplayName: "testLocation",
				Name:        "projects/testProject/location/testLocation",
			}},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	ts := httptest.NewServer(mux)

	kmsOld, err := kmsold.NewService(
		context.Background(),
		option.WithoutAuthentication(),
		option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, fmt.Errorf("failed to create kms client: %w", err)
	}

	// Create a client.
	svc, err := kms.NewKeyManagementClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &client.Services{
		KmsKeyManagementClient: svc,
		KmsoldService:          kmsOld,
	}, nil
}

type fakeKeyringsServer struct {
	pb.UnimplementedKeyManagementServiceServer
}

func (*fakeKeyringsServer) ListKeyRings(context.Context, *pb.ListKeyRingsRequest) (*pb.ListKeyRingsResponse, error) {
	resp := pb.ListKeyRingsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func (*fakeKeyringsServer) ListCryptoKeys(context.Context, *pb.ListCryptoKeysRequest) (*pb.ListCryptoKeysResponse, error) {
	resp := pb.ListCryptoKeysResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestKeyrings(t *testing.T) {
	client.MockTestHelper(t, Keyrings(), createKeyrings, client.TestOptions{})
}
