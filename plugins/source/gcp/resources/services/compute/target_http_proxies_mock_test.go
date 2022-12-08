// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func createTargetHttpProxies(mux *httprouter.Router) error {
	var item pb.TargetHttpProxyAggregatedList
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

func TestTargetHttpProxies(t *testing.T) {
	client.MockTestRestHelper(t, TargetHttpProxies(), createTargetHttpProxies, client.TestOptions{})
}
