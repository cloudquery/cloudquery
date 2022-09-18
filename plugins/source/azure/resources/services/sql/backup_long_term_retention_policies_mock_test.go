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

func createBackupLongTermRetentionPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLBackupLongTermRetentionPoliciesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			BackupLongTermRetentionPolicies: mockClient,
		},
	}

	data := sql.BackupLongTermRetentionPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewBackupLongTermRetentionPolicyListResultPage(sql.BackupLongTermRetentionPolicyListResult{Value: &[]sql.BackupLongTermRetentionPolicy{data}}, func(ctx context.Context, result sql.BackupLongTermRetentionPolicyListResult) (sql.BackupLongTermRetentionPolicyListResult, error) {
		return sql.BackupLongTermRetentionPolicyListResult{}, nil
	})

	mockClient.EXPECT().ListByDatabase(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
