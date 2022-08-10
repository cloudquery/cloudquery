package mariadb

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildMariaDBServerMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	srv := mocks.NewMockMariaDBServersClient(ctrl)
	var s mariadb.Server
	if err := faker.FakeData(&s); err != nil {
		t.Fatal(err)
	}
	fakeId := client.FakeResourceGroup + "/" + *s.ID
	s.ID = &fakeId
	srv.EXPECT().List(gomock.Any()).Return(
		mariadb.ServerListResult{Value: &[]mariadb.Server{s}},
		nil,
	)

	cfg := mocks.NewMockMariaDBConfigurationsClient(ctrl)
	var v mariadb.Configuration
	if err := faker.FakeData(&v); err != nil {
		t.Fatal(err)
	}
	cfg.EXPECT().ListByServer(gomock.Any(), "test", *s.Name).Return(
		mariadb.ConfigurationListResult{Value: &[]mariadb.Configuration{v}},
		nil,
	)
	return services.Services{
		MariaDB: services.MariaDB{
			Configurations: cfg,
			Servers:        srv,
		},
	}
}

func TestMariaDBServers(t *testing.T) {
	client.AzureMockTestHelper(t, MariadbServers(), buildMariaDBServerMock, client.TestOptions{})
}
