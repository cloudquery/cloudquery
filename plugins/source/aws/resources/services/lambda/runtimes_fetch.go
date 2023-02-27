package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lambda/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchLambdaRuntimes(_ context.Context, _ schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	for _, runtime := range types.Runtime("").Values() {
		res <- &models.RuntimeWrapper{Name: string(runtime)}
	}
	return nil
}
