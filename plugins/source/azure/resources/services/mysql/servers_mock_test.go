// Auto generated code - DO NOT EDIT.

package mysql

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
)

func TestMySQLServers(t *testing.T) {
	client.MockTestHelper(t, Servers(), createServersMock)
}

func createServersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMySQLServersClient(ctrl)
	s := services.Services{
		MySQL: services.MySQLClient{
			Servers:        mockClient,
			Configurations: createConfigurationsMock(t, ctrl).MySQL.Configurations,
		},
	}

	data := mysql.Server{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := mysql.ServerListResult{Value: &[]mysql.Server{data}}

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
