package acl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func createAcls(mux *httprouter.Router) error {
	var acl tailscale.ACL
	if err := faker.FakeObject(&acl); err != nil {
		return fmt.Errorf("failed to fake ACL: %w", err)
	}

	mux.GET("/api/v2/tailnet/:tailnet/acl", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(acl)
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

func TestAcls(t *testing.T) {
	client.MockTestHelper(t, Acls(), createAcls)
}
