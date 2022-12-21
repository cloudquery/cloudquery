package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchElastictranscoderPipelines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Elastictranscoder

	p := elastictranscoder.NewListPipelinesPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- response.Pipelines
	}

	return nil
}
