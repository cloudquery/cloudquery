package compute

import (
	"encoding/json"
	"net/http"
	"testing"

	"cloud.google.com/go/compute/apiv1/computepb"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/julienschmidt/httprouter"
)

func createRoutes(mux *httprouter.Router) error {
	var item computepb.RouteList
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	emptyStr := ""
	item.NextPageToken = &emptyStr
	mux.GET("/compute/v1/projects/testProject/global/routes", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func TestRoutes(t *testing.T) {
	client.MockTestRestHelper(t, Routes(), createRoutes, client.TestOptions{})
}
