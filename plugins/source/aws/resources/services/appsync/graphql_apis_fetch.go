package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAppsyncGraphqlApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config appsync.ListGraphqlApisInput
	c := meta.(*client.Client)
	svc := c.Services().Appsync
	for {
		output, err := svc.ListGraphqlApis(ctx, &config, func(options *appsync.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.GraphqlApis
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
