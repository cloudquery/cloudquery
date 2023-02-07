package acl

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/julienschmidt/httprouter"
)

func createAcls(mux *httprouter.Router) error {
	// this is called by init
	return nil
}

func TestAcls(t *testing.T) {
	client.MockTestHelper(t, Acls(), createAcls)
}
