package networkconnectivity

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/networkconnectivity/v1"
	"net/http"

	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func createInternalRanges(mux *httprouter.Router) error {
	var item networkconnectivity.ListInternalRangesResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	item.NextPageToken = ""
	mux.GET("/v1/projects/testProject/locations/global/internalRanges", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err = w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}
