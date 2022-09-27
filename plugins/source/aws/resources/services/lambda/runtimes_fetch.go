package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lambda/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchLambdaRuntimes(_ context.Context, _ schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	runtimes := make([]models.RuntimeWrapper, len(types.RuntimeProvidedal2.Values()))
	for i, runtime := range types.RuntimeProvidedal2.Values() {
		runtimes[i] = models.RuntimeWrapper{
			Name: string(runtime),
		}
	}
	res <- runtimes
	return nil
}
