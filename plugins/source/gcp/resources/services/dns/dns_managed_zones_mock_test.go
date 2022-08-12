package dns

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	dns "google.golang.org/api/dns/v1"
	"google.golang.org/api/option"
)

func createDnsManagedZones() (*client.Services, error) {
	ctx := context.Background()
	var zone dns.ManagedZone
	if err := faker.FakeData(&zone); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/dns/v1/projects/testProject/managedZones", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &dns.ManagedZonesListResponse{
			ManagedZones: []*dns.ManagedZone{&zone},
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
	svc, err := dns.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Dns: svc,
	}, nil
}

func TestDnsManagedZones(t *testing.T) {
	client.GcpMockTestHelper(t, DNSManagedZones(), createDnsManagedZones, client.TestOptions{})
}
