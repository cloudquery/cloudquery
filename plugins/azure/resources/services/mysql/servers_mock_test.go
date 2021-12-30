//go:build !integration

package mysql

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildMySQLServerMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	serverSvc := mocks.NewMockMySQLServerClient(ctrl)
	configSvc := mocks.NewMockMySQLConfigurationClient(ctrl)
	s := services.Services{
		MySQL: services.MySQL{
			Servers:       serverSvc,
			Configuration: configSvc,
		},
	}
	server := mysql.Server{}
	err := faker.FakeData(&server)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	name := "testServer"
	server.Name = &name
	fakeId := client.FakeResourceGroup + "/" + *server.ID
	server.ID = &fakeId
	page := mysql.ServerListResult{Value: &[]mysql.Server{server}}
	serverSvc.EXPECT().List(gomock.Any()).Return(page, nil)

	config := mysql.Configuration{}
	if err := faker.FakeData(&config); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	configSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(mysql.ConfigurationListResult{Value: &[]mysql.Configuration{config}}, nil)
	return s
}

func TestMySQLServers(t *testing.T) {
	client.AzureMockTestHelper(t, MySQLServers(), buildMySQLServerMock, client.TestOptions{})
}
