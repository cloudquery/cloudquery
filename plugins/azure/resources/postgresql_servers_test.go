package resources_test

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildPostgresServerMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	serverSvc := mocks.NewMockPostgresqlServerClient(ctrl)
	configSvc := mocks.NewMockPostgresqlConfigurationClient(ctrl)
	firewallRuleSvc := mocks.NewMockPostgresqlFirewallRuleClient(ctrl)
	s := services.Services{
		PostgreSQL: services.PostgreSQL{
			Servers:       serverSvc,
			Configuration: configSvc,
			FirewallRule:  firewallRuleSvc,
		},
	}
	server := postgresql.Server{}
	err := faker.FakeData(&server)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	name := "testServer"
	server.Name = &name
	id := fakeResourceGroup + "/" + *server.ID
	server.ID = &id
	page := postgresql.ServerListResult{Value: &[]postgresql.Server{server}}
	serverSvc.EXPECT().List(gomock.Any()).Return(page, nil)

	config := postgresql.Configuration{}
	if err := faker.FakeData(&config); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	configSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(postgresql.ConfigurationListResult{Value: &[]postgresql.Configuration{config}}, nil)

	firewallRule := postgresql.FirewallRule{}
	if err := faker.FakeData(&firewallRule); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	firewallRuleSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(postgresql.FirewallRuleListResult{Value: &[]postgresql.FirewallRule{firewallRule}}, nil)
	return s
}

func TestPostgresqlServers(t *testing.T) {
	azureTestHelper(t, resources.PostgresqlServers(), buildPostgresServerMock)
}
