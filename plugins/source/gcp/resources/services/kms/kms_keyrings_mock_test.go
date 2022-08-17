package kms

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	kms "google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/option"
)

func createKmsTestServer() (*client.Services, error) {
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
	keyring.CreateTime = time.Now().Format(time.RFC3339)
	var key kms.CryptoKey
	if err := faker.FakeData(&key); err != nil {
		return nil, err
	}
	key.Name = fmt.Sprintf("%s/cryptokey/%s", keyring.Name, "test")
	key.CreateTime = time.Now().Format(time.RFC3339)
	key.NextRotationTime = time.Now().Format(time.RFC3339)
	key.Primary.CreateTime = time.Now().Format(time.RFC3339)
	key.Primary.DestroyEventTime = time.Now().Format(time.RFC3339)
	key.Primary.DestroyTime = time.Now().Format(time.RFC3339)
	key.Primary.GenerateTime = time.Now().Format(time.RFC3339)
	key.Primary.ImportTime = time.Now().Format(time.RFC3339)
	mux := httprouter.New()
	mux.GET("/v1/projects/testProject/locations", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &kms.ListLocationsResponse{
			Locations: []*kms.Location{{
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
	var policy kms.Policy
	if err := faker.FakeData(&policy); err != nil {
		return nil, err
	}
	mux.GET("/v1/projects/testProject/location/:location/keyring/:key/cryptokey/test:getIamPolicy", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(policy)
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
	return &client.Services{
		Kms: svc,
	}, nil
}

func TestKmsKeyring(t *testing.T) {
	client.GcpMockTestHelper(t, KmsKeyrings(), createKmsTestServer, client.TestOptions{})
}
