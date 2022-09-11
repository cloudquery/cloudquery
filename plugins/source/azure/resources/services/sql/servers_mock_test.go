// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func TestSQLServers(t *testing.T) {
	client.MockTestHelper(t, Servers(), createServersMock)
}

func createServersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLServersClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			Servers: mockClient,
		},
	}

	data := sql.Server{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := sql.NewServerListResultPage(sql.ServerListResult{Value: &[]sql.Server{data}}, func(ctx context.Context, result sql.ServerListResult) (sql.ServerListResult, error) {
		return sql.ServerListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
