// Auto generated code - DO NOT EDIT.

package redis

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

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis"
)

func TestRedisResourceTypes(t *testing.T) {
	client.AzureMockTestHelper(t, ResourceTypes(), createResourceTypesMock, client.TestOptions{})
}

func createResourceTypesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockRedisResourceTypesClient(ctrl)
	s := services.Services{
		Redis: services.RedisClient{
			ResourceTypes: mockClient,
		},
	}

	data := redis.ResourceType{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := redis.NewListResultPage(redis.ListResult{Value: &[]redis.ResourceType{data}}, func(ctx context.Context, result redis.ListResult) (redis.ListResult, error) {
		return redis.ListResult{}, nil
	})

	mockClient.EXPECT().ListBySubscription(gomock.Any()).Return(result, nil)
	return s
}
