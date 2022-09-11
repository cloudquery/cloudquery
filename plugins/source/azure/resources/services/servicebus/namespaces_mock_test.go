// Auto generated code - DO NOT EDIT.

package servicebus

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

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
)

func TestServicebusNamespaces(t *testing.T) {
	client.MockTestHelper(t, Namespaces(), createNamespacesMock)
}

func createNamespacesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockServicebusNamespacesClient(ctrl)
	s := services.Services{
		Servicebus: services.ServicebusClient{
			Namespaces: mockClient,
		},
	}

	data := servicebus.SBNamespace{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := servicebus.NewSBNamespaceListResultPage(servicebus.SBNamespaceListResult{Value: &[]servicebus.SBNamespace{data}}, func(ctx context.Context, result servicebus.SBNamespaceListResult) (servicebus.SBNamespaceListResult, error) {
		return servicebus.SBNamespaceListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}
