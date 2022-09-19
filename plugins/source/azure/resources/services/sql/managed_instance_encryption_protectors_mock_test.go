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

func createManagedInstanceEncryptionProtectorsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLManagedInstanceEncryptionProtectorsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ManagedInstanceEncryptionProtectors: mockClient,
		},
	}

	data := sql.ManagedInstanceEncryptionProtector{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewManagedInstanceEncryptionProtectorListResultPage(sql.ManagedInstanceEncryptionProtectorListResult{Value: &[]sql.ManagedInstanceEncryptionProtector{data}}, func(ctx context.Context, result sql.ManagedInstanceEncryptionProtectorListResult) (sql.ManagedInstanceEncryptionProtectorListResult, error) {
		return sql.ManagedInstanceEncryptionProtectorListResult{}, nil
	})

	mockClient.EXPECT().ListByInstance(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
