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

func createDatabaseThreatDetectionPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLDatabaseThreatDetectionPoliciesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			DatabaseThreatDetectionPolicies: mockClient,
		},
	}

	data := sql.DatabaseSecurityAlertPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewDatabaseSecurityAlertPolicyListResultPage(sql.DatabaseSecurityAlertPolicyListResult{Value: &[]sql.DatabaseSecurityAlertPolicy{data}}, func(ctx context.Context, result sql.DatabaseSecurityAlertPolicyListResult) (sql.DatabaseSecurityAlertPolicyListResult, error) {
		return sql.DatabaseSecurityAlertPolicyListResult{}, nil
	})

	mockClient.EXPECT().Get(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
