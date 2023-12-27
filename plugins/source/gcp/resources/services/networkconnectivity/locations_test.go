package networkconnectivity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/networkconnectivity/v1"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func createLocation(mux *httprouter.Router) error {
	var item networkconnectivity.ListLocationsResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	item.Locations = []*networkconnectivity.Location{
		{
			LocationId: "global",
			Labels:     map[string]string{"cloud.googleapis.com/region": "us-east1"},
			Name:       "projects/testProject/locations/global",
			Metadata:   []byte(`{"@type":"type.googleapis.com/google.cloud.networkconnectivity.v1.LocationMetadata","locationFeatures":["SITE_TO_CLOUD_SPOKES"]}`),
		},
	}
	item.NextPageToken = ""
	mux.GET("/v1/projects/testProject/locations", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&item)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err = w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return createInternalRanges(mux)
}

func TestLocations(t *testing.T) {
	client.MockTestHelper(t, Locations(), client.WithCreateHTTPServer(createLocation))
}
