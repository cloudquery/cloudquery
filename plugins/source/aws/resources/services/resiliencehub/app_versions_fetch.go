package resiliencehub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAppVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Resiliencehub
	p := resiliencehub.NewListAppVersionsPaginator(svc, &resiliencehub.ListAppVersionsInput{AppArn: parent.Item.(*types.App).AppArn})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- out.AppVersions
	}
	return nil
}
