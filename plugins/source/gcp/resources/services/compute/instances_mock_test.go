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

func createInstances(mux *httprouter.Router) error {
	var item pb.InstanceAggregatedList
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	emptyStr := ""
	item.NextPageToken = &emptyStr
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&item)
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

func TestInstances(t *testing.T) {
	client.MockTestRestHelper(t, Instances(), createInstances, client.TestOptions{})
}
