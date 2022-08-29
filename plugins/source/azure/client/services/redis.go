//go:generate mockgen -destination=./mocks/redis.go -package=mocks . RedisResourceTypesClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis"
	"github.com/Azure/go-autorest/autorest"
)

type RedisClient struct {
	ResourceTypes RedisResourceTypesClient
}

type RedisResourceTypesClient interface {
	ListBySubscription(ctx context.Context) (result redis.ListResultPage, err error)
}

func NewRedisClient(subscriptionId string, auth autorest.Authorizer) RedisClient {
	cl := redis.NewClient(subscriptionId)
	cl.Authorizer = auth

	return RedisClient{ResourceTypes: cl}
}
