package kms

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func keyGrants() *schema.Table {
	tableName := "aws_kms_key_grants"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/kms/latest/APIReference/API_GrantListEntry.html`,
		Resolver:    fetchKmsKeyGrants,
		Transform:   transformers.TransformWithStruct(&types.GrantListEntry{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "kms"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "key_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:       "grant_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("GrantId"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchKmsKeyGrants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	k := parent.Item.(*types.KeyMetadata)
	config := kms.ListGrantsInput{
		KeyId: k.Arn,
		Limit: aws.Int32(100),
	}

	cl := meta.(*client.Client)
	svc := cl.Services().Kms
	p := kms.NewListGrantsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *kms.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Grants
	}
	return nil
}
