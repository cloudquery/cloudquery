package emr

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Studios() *schema.Table {
	tableName := "aws_emr_studios"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/emr/latest/APIReference/API_Studio.html`,
		Resolver:            fetchEmrStudios,
		PreResourceResolver: getStudio,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "elasticmapreduce"),
		Transform:           transformers.TransformWithStruct(&types.Studio{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Description:         `The Amazon Resource Name (ARN) of the EMR Studio.`,
				Resolver:            schema.PathResolver("StudioArn"),
				PrimaryKeyComponent: true,
			},
		},
		Relations: []*schema.Table{
			studioSessionMapping(),
		},
	}
}

func fetchEmrStudios(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	config := emr.ListStudiosInput{}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	paginator := emr.NewListStudiosPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Studios
	}
	return nil
}

func getStudio(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	response, err := svc.DescribeStudio(ctx, &emr.DescribeStudioInput{StudioId: resource.Item.(types.StudioSummary).StudioId}, func(options *emr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = response.Studio
	return nil
}
