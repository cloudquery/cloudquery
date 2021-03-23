package resource

import "context"

type ClientInterface interface {
	CollectResource(ctx context.Context, resource string, config interface{}) error
}
