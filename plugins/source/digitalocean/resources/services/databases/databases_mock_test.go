package databases

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createDatabases(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDatabasesService(ctrl)

	var data []godo.Database
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	createFirewallRules(t, m)
	createReplicas(t, m)
	createBackups(t, m)

	return client.Services{
		Databases: m,
	}
}

func TestDatabases(t *testing.T) {
	client.MockTestHelper(t, Databases(), createDatabases, client.TestOptions{})
}
