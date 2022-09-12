package databases

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createDatabases(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDatabasesService(ctrl)

	var data []godo.Database
	if err := faker.FakeData(&data); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	var backups []godo.DatabaseBackup
	if err := faker.FakeData(&backups); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListBackups(gomock.Any(), gomock.Any(), gomock.Any()).Return(backups, &godo.Response{}, nil)

	var replicas []godo.DatabaseReplica
	if err := faker.FakeData(&replicas); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListReplicas(gomock.Any(), gomock.Any(), gomock.Any()).Return(replicas, &godo.Response{}, nil)

	var firewallRules []godo.DatabaseFirewallRule
	if err := faker.FakeData(&replicas); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetFirewallRules(gomock.Any(), gomock.Any()).Return(firewallRules, &godo.Response{}, nil)

	return client.Services{
		Databases: m,
	}
}

func TestDatabases(t *testing.T) {
	client.MockTestHelper(t, Databases(), createDatabases, client.TestOptions{})
}
