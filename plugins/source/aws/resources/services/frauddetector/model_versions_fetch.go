package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchFrauddetectorModelVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := frauddetector.NewDescribeModelVersionsPaginator(meta.(*client.Client).Services().Frauddetector,
		&frauddetector.DescribeModelVersionsInput{ModelId: parent.Item.(types.Model).ModelId})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.ModelVersionDetails
	}
	return nil
}
