package compute

import (
	"encoding/json"
	"net/http"
	"testing"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
)

func createZones(mux *httprouter.Router) error {
	var zones pb.ZoneList
	if err := faker.FakeObject(&zones); err != nil {
		return err
	}
	emptyStr := ""
	zones.NextPageToken = &emptyStr

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

	return createMachineTypes(mux, &zones)
}

func TestZones(t *testing.T) {
	client.MockTestRestHelper(t, Zones(), createZones, client.TestOptions{})
}
