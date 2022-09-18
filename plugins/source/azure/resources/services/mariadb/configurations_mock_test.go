// Auto generated code - DO NOT EDIT.

package mariadb

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"
)

func createConfigurationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMariaDBConfigurationsClient(ctrl)
	s := services.Services{
		MariaDB: services.MariaDBClient{
			Configurations: mockClient,
		},
	}

	data := mariadb.Configuration{}
	require.Nil(t, faker.FakeObject(&data))

	result := mariadb.ConfigurationListResult{Value: &[]mariadb.Configuration{data}}

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
