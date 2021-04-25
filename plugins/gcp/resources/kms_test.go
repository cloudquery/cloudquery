package resources_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-gcp/resources"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	kms "google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/option"
)

func createKmsTestServer() (*kms.Service, error) {
	ctx := context.Background()
	var location kms.Location
	if err := faker.FakeData(&location); err != nil {
		return nil, err
	}
	var keyring kms.KeyRing
	if err := faker.FakeData(&keyring); err != nil {
		return nil, err
	}
	keyring.Name = fmt.Sprintf("projects/testProject/location/%s/keyring/%s", location.Name, keyring.Name)
	var key kms.CryptoKey
	if err := faker.FakeData(&key); err != nil {
		return nil, err
	}
	key.Name = fmt.Sprintf("%s/cryptokey/%s", keyring.Name, key.Name)

	mux := httprouter.New()
	mux.GET("/v1/projects/testProject/locations", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &kms.ListLocationsResponse{
			Locations: []*kms.Location{{
				DisplayName: faker.Name(),
				Name:        fmt.Sprintf("projects/testProject/location/%s", location.Name),
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
	mux.GET("/v1/projects/testProject/location/:location/keyRings", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &kms.ListKeyRingsResponse{
			KeyRings: []*kms.KeyRing{&keyring},
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
	mux.GET("/v1/projects/testProject/location/:location/keyRing/:keyring/cryptoKeys", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &kms.ListCryptoKeysResponse{
			CryptoKeys: []*kms.CryptoKey{&key},
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
	svc, err := kms.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func TestKmsKeyring(t *testing.T) {
	resource := ResourceTestData{
		Table: resources.KmsKeyring(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
			Resources:  []client.Resource{{Name: "kms.keys"}},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			kmsSvc, err := createKmsTestServer()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				Kms: kmsSvc,
			})
			return c, nil
		},
	}
	testResource(t, resources.Provider, resource)
}
