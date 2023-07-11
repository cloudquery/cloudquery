package cloudtrail

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Imports() *schema.Table {
	tableName := "aws_cloudtrail_imports"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_GetImport.html`,
		Resolver:            fetchImports,
		PreResourceResolver: getImport,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cloudtrail"),
		Transform: transformers.TransformWithStruct(
			&cloudtrail.GetImportOutput{},
			transformers.WithSkipFields("ResultMetadata"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ImportId"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchImports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudtrail

	paginator := cloudtrail.NewListImportsPaginator(svc, nil)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudtrail.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Imports
	}
	return nil
}

func getImport(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudtrail
	item := resource.Item.(types.ImportsListItem)
	importOutput, err := svc.GetImport(ctx, &cloudtrail.GetImportInput{ImportId: item.ImportId}, func(options *cloudtrail.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = importOutput

	return nil
}
