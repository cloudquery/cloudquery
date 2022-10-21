package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchFrauddetectorVariables(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	paginator := frauddetector.NewGetVariablesPaginator(meta.(*client.Client).Services().FraudDetector, nil)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Variables
	}
	return nil
}
