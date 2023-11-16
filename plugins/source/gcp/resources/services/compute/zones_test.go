package compute

import (
	"encoding/json"
	"net/http"
	"testing"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/julienschmidt/httprouter"
)

func createZones(mux *httprouter.Router) error {
	var zones pb.ZoneList
	if err := faker.FakeObject(&zones); err != nil {
		return err
	}
	emptyStr := ""
	zones.NextPageToken = &emptyStr
	zones.Items = zones.Items[:1] // leave only 1

	mux.GET("/compute/v1/projects/testProject/zones", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&zones)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	if err := createMachineTypes(mux, zones.Items[0]); err != nil {
		return err
	}

	return createInventories(mux, zones.Items[0])
}

func TestZones(t *testing.T) {
	client.MockTestRestHelper(t, Zones(), createZones, client.TestOptions{})
}
