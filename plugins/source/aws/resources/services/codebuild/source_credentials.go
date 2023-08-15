package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SourceCredentials() *schema.Table {
	tableName := "aws_codebuild_source_credentials"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/codebuild/latest/APIReference/API_SourceCredentialsInfo.html`,
		Resolver:    fetchSourceCredentials,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "codebuild"),
		Transform:   transformers.TransformWithStruct(&types.SourceCredentialsInfo{}, transformers.WithPrimaryKeys("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchSourceCredentials(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Codebuild
	credentialsOutput, err := svc.ListSourceCredentials(ctx, nil, func(options *codebuild.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- credentialsOutput.SourceCredentialsInfos
	return nil
}
