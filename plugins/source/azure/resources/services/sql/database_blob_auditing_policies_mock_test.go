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

func createDatabaseBlobAuditingPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLDatabaseBlobAuditingPoliciesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			DatabaseBlobAuditingPolicies: mockClient,
		},
	}

	data := sql.DatabaseBlobAuditingPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewDatabaseBlobAuditingPolicyListResultPage(sql.DatabaseBlobAuditingPolicyListResult{Value: &[]sql.DatabaseBlobAuditingPolicy{data}}, func(ctx context.Context, result sql.DatabaseBlobAuditingPolicyListResult) (sql.DatabaseBlobAuditingPolicyListResult, error) {
		return sql.DatabaseBlobAuditingPolicyListResult{}, nil
	})

	mockClient.EXPECT().ListByDatabase(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
