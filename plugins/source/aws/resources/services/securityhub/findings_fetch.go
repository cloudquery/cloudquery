package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Securityhub
	config := securityhub.GetFindingsInput{
		MaxResults: 100,
	}
	p := securityhub.NewGetFindingsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Findings
	}
	return nil
}
