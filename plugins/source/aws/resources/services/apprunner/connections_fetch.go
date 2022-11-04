package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchApprunnerConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apprunner.ListConnectionsInput
	svc := meta.(*client.Client).Services().Apprunner
	paginator := apprunner.NewListConnectionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.ConnectionSummaryList
	}
	return nil
}
