// Auto generated code - DO NOT EDIT.

package mariadb

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"
)

func TestMariaDBServers(t *testing.T) {
	client.MockTestHelper(t, Servers(), createServersMock)
}

func createServersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMariaDBServersClient(ctrl)
	s := services.Services{
		MariaDB: services.MariaDBClient{
			Servers: mockClient,
		},
	}

	data := mariadb.Server{}
	require.Nil(t, faker.FakeObject(&data))

	result := mariadb.ServerListResult{Value: &[]mariadb.Server{data}}

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
