// Auto generated code - DO NOT EDIT.

package mysql

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
)

func createConfigurationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {

	mockClient := mocks.NewMockMySQLConfigurationsClient(ctrl)
	s := services.Services{
		MySQL: services.MySQLClient{
			Configurations: mockClient,
		},
	}

	data := mysql.Configuration{}
	require.Nil(t, faker.FakeObject(&data))

	result := mysql.ConfigurationListResult{Value: &[]mysql.Configuration{data}}

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
