// Auto generated code - DO NOT EDIT.

package postgresql

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
)

func createConfigurationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockPostgreSQLConfigurationsClient(ctrl)
	s := services.Services{
		PostgreSQL: services.PostgreSQLClient{
			Configurations: mockClient,
		},
	}

	data := postgresql.Configuration{}
	require.Nil(t, faker.FakeObject(&data))

	result := postgresql.ConfigurationListResult{Value: &[]postgresql.Configuration{data}}

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
