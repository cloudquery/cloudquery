package device

import (
	"encoding/json"
	"net/http"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func createRoutes(mux *httprouter.Router) error {
	var route tailscale.DeviceRoutes
	if err := faker.FakeObject(&route); err != nil {
		return err
	}

	mux.GET("/api/v2/device/:device/routes", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(route)
		if err != nil {
			http.Error(w, "unable to marshal response: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}
