// Auto generated code - DO NOT EDIT.

package postgresql

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
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
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := postgresql.ServerListResult{Value: &[]postgresql.Server{data}}

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
