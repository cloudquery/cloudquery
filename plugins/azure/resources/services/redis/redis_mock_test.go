//go:build !integration

package redis

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRedisClientMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockRedisClient(ctrl)
	var rt redis.ResourceType
	if err := faker.FakeData(&rt); err != nil {
		t.Fatal(err)
	}
	ip := "192.168.1.1"
	rt.StaticIP = &ip
	m.EXPECT().ListBySubscription(gomock.Any()).Return(
		redis.NewListResultPage(
			redis.ListResult{Value: &[]redis.ResourceType{rt}},
			func(c context.Context, lr redis.ListResult) (redis.ListResult, error) {
				return redis.ListResult{}, nil
			},
		),
		nil,
	)
	return services.Services{Redis: m}
}

func TestRedisServices(t *testing.T) {
	client.AzureMockTestHelper(t, RedisServices(), buildRedisClientMock, client.TestOptions{})
}
