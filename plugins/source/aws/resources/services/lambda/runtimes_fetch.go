package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/plugin-sdk/schema"
)

type RuntimeWrapper struct {
	Name string
}

func fetchLambdaRuntimes(_ context.Context, _ schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	runtimes := make([]RuntimeWrapper, len(types.RuntimeProvidedal2.Values()))
	for i, runtime := range types.RuntimeProvidedal2.Values() {
		runtimes[i] = RuntimeWrapper{
			Name: string(runtime),
		}
	}
	res <- runtimes
	return nil
}
