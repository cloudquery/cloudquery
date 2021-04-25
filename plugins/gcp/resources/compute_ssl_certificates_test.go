package resources_test

import (
	"context"
	"encoding/json"
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
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func TestComputeSslCertificates(t *testing.T) {
	resource := ResourceTestData{
		Table: resources.ComputeBackendServices(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
			Resources: []client.Resource{
				{Name: "compute.ssl_certificates"},
			},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			computeSvc, err := createSslCertificatesServer()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				Compute: computeSvc,
			})
			return c, nil
		},
	}
	testResource(t, resources.Provider, resource)
}

func createSslCertificatesServer() (*compute.Service, error) {
	ctx := context.Background()
	var inst compute.SslCertificate
	if err := faker.FakeData(&inst); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/projects/testProject/aggregated/sslCertificates", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &compute.SslCertificateAggregatedList{
			Items: map[string]compute.SslCertificatesScopedList{
				"": {
					SslCertificates: []*compute.SslCertificate{&inst},
				},
			},
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
	svc, err := compute.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return svc, nil
}
