package osconfig

import (
	"net/http"
	"testing"

	pb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/protobuf/encoding/protojson"
)

func createInventories(mux *httprouter.Router) error {
	var item pb.ListInventoriesResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	item.NextPageToken = ""
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := protojson.Marshal(&item)
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

func TestInventories(t *testing.T) {
	client.MockTestRestHelper(t, Inventories(), createInventories, client.TestOptions{})
}
