package dns

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/api/dns/v1"
)

type MockManagedZonesResult struct {
	ManagedZones []*dns.ManagedZone `json:"managedzones,omitempty"`
}

type MockResourceRecordSetsResult struct {
	Rrsets []*dns.ResourceRecordSet `json:"rrsets,omitempty"`
}

func createManagedZones(mux *httprouter.Router) error {
	var managedZoneItem dns.ManagedZone
	if err := faker.FakeObject(&managedZoneItem); err != nil {
		return err
	}

	mux.GET("/dns/v1/projects/testProject/managedZones", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &MockManagedZonesResult{
			ManagedZones: []*dns.ManagedZone{&managedZoneItem},
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

	var recordSetItem dns.ResourceRecordSet
	if err := faker.FakeObject(&recordSetItem); err != nil {
		return err
	}

	mux.GET("/dns/v1/projects/testProject/managedZones/test string/rrsets", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &MockResourceRecordSetsResult{
			Rrsets: []*dns.ResourceRecordSet{&recordSetItem},
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

	return nil
}

func TestManagedZones(t *testing.T) {
	client.MockTestRestHelper(t, ManagedZones(), createManagedZones, client.TestOptions{})
}
