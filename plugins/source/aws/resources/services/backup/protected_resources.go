package backup

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func ProtectedResources() *schema.Table {
	tableName := "aws_backup_protected_resources"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeProtectedResource.html`,
		Resolver:            listResources,
		PreResourceResolver: getResource,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:           transformers.TransformWithStruct(&backup.DescribeProtectedResourceOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ResourceArn"),
				PrimaryKey: true,
			},
		},
	}
}

func listResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListProtectedResourcesInput{MaxResults: aws.Int32(1000)} // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListProtectedResources.html
	paginator := backup.NewListProtectedResourcesPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *backup.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Results
	}
	return nil
}

func getResource(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	pr := resource.Item.(types.ProtectedResource)
	resp, err := svc.DescribeProtectedResource(ctx, &backup.DescribeProtectedResourceInput{
		ResourceArn: pr.ResourceArn,
	}, func(options *backup.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = resp
	return nil
}
