// Auto generated code - DO NOT EDIT.

package redis

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis"
)

func TestRedisCaches(t *testing.T) {
	client.MockTestHelper(t, Caches(), createCachesMock)
}

func createCachesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockRedisCachesClient(ctrl)
	s := services.Services{
		Redis: services.RedisClient{
			Caches: mockClient,
		},
	}

	data := redis.ResourceType{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := redis.NewListResultPage(redis.ListResult{Value: &[]redis.ResourceType{data}}, func(ctx context.Context, result redis.ListResult) (redis.ListResult, error) {
		return redis.ListResult{}, nil
	})

	mockClient.EXPECT().ListBySubscription(gomock.Any()).Return(result, nil)
	return s
}
