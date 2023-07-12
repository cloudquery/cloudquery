package kms

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Aliases() *schema.Table {
	tableName := "aws_kms_aliases"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/kms/latest/APIReference/API_AliasListEntry.html`,
		Resolver:    fetchKmsAliases,
		Transform:   transformers.TransformWithStruct(&types.AliasListEntry{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "kms"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("AliasArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchKmsAliases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input kms.ListAliasesInput
	cl := meta.(*client.Client)
	svc := cl.Services().Kms
	paginator := kms.NewListAliasesPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *kms.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Aliases
	}
	return nil
}
