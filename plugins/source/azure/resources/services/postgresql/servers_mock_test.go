// Auto generated code - DO NOT EDIT.

package postgresql

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
)

func TestPostgreSQLServers(t *testing.T) {
	client.MockTestHelper(t, Servers(), createServersMock)
}

func createServersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockPostgreSQLServersClient(ctrl)
	s := services.Services{
		PostgreSQL: services.PostgreSQLClient{
			Servers: mockClient,
		},
	}

	data := postgresql.Server{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := postgresql.ServerListResult{Value: &[]postgresql.Server{data}}

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
