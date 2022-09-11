// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func TestSQLManagedInstances(t *testing.T) {
	client.MockTestHelper(t, ManagedInstances(), createManagedInstancesMock)
}

func createManagedInstancesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLManagedInstancesClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ManagedInstances: mockClient,
		},
	}

	data := sql.ManagedInstance{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := sql.NewManagedInstanceListResultPage(sql.ManagedInstanceListResult{Value: &[]sql.ManagedInstance{data}}, func(ctx context.Context, result sql.ManagedInstanceListResult) (sql.ManagedInstanceListResult, error) {
		return sql.ManagedInstanceListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
