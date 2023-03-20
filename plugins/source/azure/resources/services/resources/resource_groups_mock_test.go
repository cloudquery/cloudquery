package resources

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/gorilla/mux"
)

func createResourceGroups(router *mux.Router) error {
	// We already fetched the resource groups for this subscription, no need to fetch again
	return nil
}

func TestResourceGroups(t *testing.T) {
	client.MockTestHelper(t, ResourceGroups(), createResourceGroups)
}
