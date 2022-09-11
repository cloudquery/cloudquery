// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func createServerAdminsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLServerAdminsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ServerAdmins: mockClient,
		},
	}

	data := sql.ServerAzureADAdministrator{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewServerAzureADAdministratorListResultPage(sql.ServerAzureADAdministratorListResult{Value: &[]sql.ServerAzureADAdministrator{data}}, func(ctx context.Context, result sql.ServerAzureADAdministratorListResult) (sql.ServerAzureADAdministratorListResult, error) {
		return sql.ServerAzureADAdministratorListResult{}, nil
	})

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
