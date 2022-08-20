package security

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
	secretmanager "google.golang.org/api/secretmanager/v1"
)

func createSecretsServer() (*client.Services, error) {
	var secret secretmanager.Secret

	if err := faker.FakeData(&secret); err != nil {
		return nil, err
	}

	mux := httprouter.New()
	mux.GET("/v1/projects/testProject/secrets", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &secretmanager.ListSecretsResponse{
			Secrets:   []*secretmanager.Secret{&secret},
			TotalSize: 1,
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
	svc, err := secretmanager.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		SecretManager: svc,
	}, nil
}

func TestSecrets(t *testing.T) {
	client.GcpMockTestHelper(t, Secrets(), createSecretsServer, client.TestOptions{})
}
